#!/usr/bin/env python3
"""Fetch a sample of real IETF list mail and report what IETF-Email does to it.

The IETF archive is readable over anonymous IMAP, so the corpus is reproducible
by anyone: no credentials, no scraping. Messages are never committed.

Per docs/validation.md, this measures two things and decides nothing: whether a
rule matches text its source does not describe, and whether its alert rate is
supportable. It cannot tell us which rules should exist, and a rule that finds
nothing here has not been disproven -- it may simply mean the sample is clean.
"""
import argparse, email, imaplib, pathlib, re, subprocess, sys, collections, json
from email import policy

imaplib._MAXLINE = 10_000_000

DEFAULT_LISTS = ['ietf', 'last-call', 'tls', 'dnsop', 'httpapi', 'oauth',
                 'gendispatch', 'rfced-future']


def body_text(msg):
    if msg.is_multipart():
        for part in msg.walk():
            if part.get_content_type() == 'text/plain':
                try:
                    return part.get_content()
                except Exception:
                    pass
        return ''
    try:
        return msg.get_content() if msg.get_content_type() == 'text/plain' else ''
    except Exception:
        return ''


def clean(text):
    """Keep what the author wrote: drop quoted text, signatures and bare links."""
    out = []
    for line in text.split('\n'):
        t = line.rstrip()
        if t.startswith('>'):
            continue
        if re.match(r'^\s*(On .{0,80}wrote:|-- $|--$|___{3,})', t):
            break
        if re.match(r'^\s*https?://\S+\s*$', t):
            continue
        out.append(t)
    return '\n'.join(out).strip()


def fetch(lists, count, email_address, dest):
    M = imaplib.IMAP4_SSL('imap.ietf.org', 993, timeout=30)
    M.login('anonymous', email_address)
    saved = 0
    for lst in lists:
        typ, data = M.select(f'"Shared Folders/{lst}"', readonly=True)
        if typ != 'OK':
            print(f'  {lst}: cannot select', file=sys.stderr)
            continue
        n = int(data[0])
        ids = range(max(1, n - count + 1), n + 1)
        typ, fetched = M.fetch(','.join(map(str, ids)), '(RFC822)')
        got = 0
        for item in fetched:
            if not isinstance(item, tuple):
                continue
            try:
                text = clean(body_text(email.message_from_bytes(item[1], policy=policy.default)))
            except Exception:
                continue
            if len(text) < 120:
                continue
            got += 1
            saved += 1
            (dest / f'{lst}_{got:03d}.txt').write_text(text + '\n', encoding='utf-8')
        print(f'  {lst}: saved {got} of {n} archived')
    M.logout()
    return saved


def report(dest, styles_path):
    cfg = dest / '.vale.ini'
    cfg.write_text(f'StylesPath = {styles_path}\nMinAlertLevel = suggestion\n'
                   f'[*.txt]\nBasedOnStyles = IETF-Email\n')
    proc = subprocess.run(['vale', '--output=JSON', '--no-global', '--no-exit', str(dest)],
                          capture_output=True, text=True, cwd=dest)
    alerts = json.loads(proc.stdout or '{}')
    files = list(dest.glob('*.txt'))
    lines = sum(1 for f in files for _ in f.open(encoding='utf-8', errors='replace'))

    per_rule = collections.Counter()
    hit_docs = collections.defaultdict(set)
    samples = collections.defaultdict(list)
    for fname, items in alerts.items():
        for a in items:
            per_rule[a['Check']] += 1
            hit_docs[a['Check']].add(fname)
            if len(samples[a['Check']]) < 8:
                samples[a['Check']].append((pathlib.Path(fname).name, a['Line'], a['Match']))

    print(f'\ncorpus: {len(files)} messages, {lines:,} lines of unquoted body text\n')
    print(f"{'rule':38} {'alerts':>7} {'msgs':>6} {'per 1k lines':>13}")
    for rule, n in per_rule.most_common():
        print(f'{rule:38} {n:>7} {len(hit_docs[rule]):>6} {1000 * n / lines:>13.2f}')
    total = sum(per_rule.values())
    print(f"{'TOTAL':38} {total:>7} {'':>6} {1000 * total / lines:>13.2f}")

    print('\nsamples to adjudicate (accept = a finding you would raise):')
    for rule, s in samples.items():
        print(f'\n  {rule}')
        for f, ln, m in s:
            print(f'    {f}:{ln}  {m!r}')


def main():
    ap = argparse.ArgumentParser(description=__doc__)
    ap.add_argument('--email', required=True, help='your address, used as the anonymous IMAP password')
    ap.add_argument('--lists', nargs='*', default=DEFAULT_LISTS)
    ap.add_argument('--count', type=int, default=80, help='most recent messages per list')
    ap.add_argument('--dest', default='.corpus/mail')
    ap.add_argument('--styles-path', default=str(pathlib.Path(__file__).resolve().parent.parent))
    args = ap.parse_args()

    dest = pathlib.Path(args.dest)
    dest.mkdir(parents=True, exist_ok=True)
    print(f'fetching from imap.ietf.org as anonymous:')
    fetch(args.lists, args.count, args.email, dest)
    report(dest, args.styles_path)


if __name__ == '__main__':
    main()
