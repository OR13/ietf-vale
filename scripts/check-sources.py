#!/usr/bin/env python3
"""Fail if any rule cites an RFC that has since been obsoleted.

Citing a superseded document is the fastest way to lose an argument about a
rule, so this runs in CI rather than relying on anyone remembering. It reads the
RFC index, which is the authoritative record of what replaced what.
"""
import pathlib, re, sys, urllib.request

INDEX = 'https://www.rfc-editor.org/rfc-index.xml'


def cited_rfcs():
    """Every RFC number a rule's `link:` points at, by rule file."""
    out = {}
    for style in pathlib.Path('.').glob('IETF-*'):
        for rule in style.glob('*.yml'):
            text = rule.read_text()
            m = re.search(r'^link:\s*(\S+)', text, re.M)
            if not m:
                continue
            for num in re.findall(r'rfc(\d{3,5})', m.group(1), re.I):
                out.setdefault(f'{style.name}/{rule.name}', set()).add(num)
            # info/bcpNN links resolve to whatever RFCs the BCP holds today, so
            # they cannot go stale and are not checked here.
    return out


def obsoletions(numbers):
    index = urllib.request.urlopen(INDEX, timeout=120).read().decode('utf-8', 'replace')
    entries = re.findall(r'<rfc-entry>.*?</rfc-entry>', index, re.S)
    by_id = {}
    for e in entries:
        m = re.search(r'<doc-id>RFC(\d+)</doc-id>', e)
        if m:
            by_id[str(int(m.group(1)))] = e
    result = {}
    for n in numbers:
        e = by_id.get(str(int(n)))
        if e is None:
            result[n] = ['NOT IN INDEX']
            continue
        blocks = re.findall(r'<obsoleted-by>(.*?)</obsoleted-by>', e, re.S)
        ids = [f'RFC {int(x)}' for b in blocks for x in re.findall(r'<doc-id>RFC(\d+)</doc-id>', b)]
        if ids:
            result[n] = ids
    return result


def main():
    cited = cited_rfcs()
    every = {n for nums in cited.values() for n in nums}
    if not every:
        print('no RFC-numbered links to check')
        return 0
    stale = obsoletions(every)
    if not stale:
        print(f'checked {len(every)} cited RFCs, none obsoleted')
        return 0
    for rule, nums in sorted(cited.items()):
        for n in sorted(nums):
            if n in stale:
                print(f'{rule}: cites RFC {int(n)}, obsoleted by {", ".join(stale[n])}', file=sys.stderr)
    return 1


if __name__ == '__main__':
    sys.exit(main())
