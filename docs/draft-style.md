---
type: Style Specification
title: Draft style
description: Normative specification for the IETF-Draft Vale style, which lints Internet-Draft and RFC-bound prose.
resource: https://github.com/OR13/ietf-vale
tags: [vale, internet-drafts, rfc-editor, style]
status: draft
timestamp: 2026-09-19T00:00:00Z
generated: { by: claude/opus-5, at: 2026-09-19T00:00:00Z }
sources:
  - id: rfc7322
    resource: https://www.rfc-editor.org/rfc/rfc7322.html
    title: "RFC 7322: RFC Style Guide"
    last_modified: 2014-09-01T00:00:00Z
  - id: styleguide
    resource: https://www.rfc-editor.org/styleguide/
    title: Updates to the RFC Style Guide (living document)
  - id: authors-language
    resource: https://authors.ietf.org/language-and-style
    title: "Internet-Draft Author Resources: Language and style"
  - id: authors-examples
    resource: https://authors.ietf.org/example-addresses
    title: "Internet-Draft Author Resources: Example addresses"
  - id: bcp14
    resource: https://www.rfc-editor.org/info/bcp14
    title: "BCP 14: RFC 2119 and RFC 8174"
  - id: rpc-terms
    resource: https://rpc-wiki.rfc-editor.org/doku.php?id=terms
    title: RFC-Specific Terms list
  - id: rpc-abbrev
    resource: https://rpc-wiki.rfc-editor.org/doku.php?id=abbrev_list
    title: Abbreviations list
  - id: iesg-bcp14
    resource: https://datatracker.ietf.org/doc/statement-iesg-statement-on-clarifying-the-use-of-bcp-14-key-words/
    title: IESG Statement on clarifying the use of BCP 14 key words (2025-03-17)
  - id: idnits
    resource: https://github.com/ietf-tools/idnits
    title: idnits, the Internet-Draft nits checker
---

This specifies the `IETF-Draft` Vale style: the rules it contains, the guidance each
rule implements, and the boundaries of what it claims. Severity levels follow the
[evidence policy](/evidence-policy.md). The companion
[email style](/email-style.md) covers mailing list prose and shares no rules with
this one.

# Scope

**In scope**: prose in Internet-Drafts and in documents heading for RFC
publication, in Markdown (including kramdown-rfc), RFCXML, and plain text.

**Out of scope**, because tools already own these and duplicating them badly is
worse than leaving them alone:

| Concern | Owner |
|---|---|
| IPR and copyright boilerplate, expiry statement, required-section presence, outdated references, version and date nits | idnits[^idnits] and datatracker submission checks |
| RFCXML validity, reference resolution, generated boilerplate | xml2rfc, kramdown-rfc |
| Reference entry formatting, title-page header, "Status of This Memo" text | the RFC Production Center at publication |

# Conformance

A document conforms to this style when Vale reports no `error`-level alerts from
it. `warning` and `suggestion` alerts are review prompts, not defects.

Passing this style does **not** mean a document is ready for submission, has
correct boilerplate, or will survive RFC Editor review. It means the prose
contains none of the specific, cited defects enumerated below.

# Normative basis

Every rule traces to one of:

- **RFC 7322**[^rfc7322], the RFC Style Guide, still current and updated by
  RFC 7997.
- The RFC Editor's living **Updates to the RFC Style Guide**[^styleguide], which
  tiers its guidance as MUSTs, RECOMMENDED, and Author Choice.
- **authors.ietf.org** content guidelines[^authors-language][^authors-examples],
  which the site states are "the direct replacement for the previous 'I-D
  Guidelines'".
- **BCP 14**[^bcp14] and the IESG statement on its key words[^iesg-bcp14].
- The RFC Production Center's **Terms**[^rpc-terms] and
  **Abbreviations**[^rpc-abbrev] lists, both of which expose a raw export that
  rules are generated from.

# Rule inventory

Each rule is one PR. `Level` is the target level under the evidence policy. A
rule whose alert rate makes it unsupportable ships one level lower, or disabled
by default; it is not dropped for being rarely triggered, since a clean corpus
says nothing about whether the guidance exists. See the
[validation methodology](/validation.md).

## Group 1: mechanical

These have precise triggers and, by inspection of their patterns, near-zero
scope for matching text their sources do not describe.

| Rule | Requirement | Basis | Extends | Level |
|---|---|---|---|---|
| `IETF-Draft.Terms` | Use the RFC Production Center's settled spelling of RFC-specific terms, except where the list defers to internal consistency | Terms list: "IPsec — Not 'IPSEC' or 'IPSec'"; "email … (not 'e-mail' or 'Email')"; "online … (not 'on-line')"; "timestamp … (not 'time-stamp')"; "ASCII — Not 'US-ASCII'"[^rpc-terms] | `substitution` | warning |
| `IETF-Draft.CitationSpacing` | A citation tag contains no spaces | RFC 7322 §3.5: "A citation/reference tag must not contain spaces", e.g. `[RFC2119]` not `[RFC 2119]`[^rfc7322] | `substitution` | error |
| `IETF-Draft.RFCCompoundHyphen` | Do not form compounds by hyphenating an RFC citation | Style Guide RECOMMENDED: "Avoid forming compounds by hyphenating RFC numbers", e.g. `[RFC5011]-style rollover`[^styleguide] | `existence` | warning |
| `IETF-Draft.AbbreviationVerbs` | Affix suffixes to abbreviations without punctuation | Style Guide RECOMMENDED: "'XORed' (not XOR'ed) and 'NATed' (not NAT-ed)"[^styleguide] | `existence` | warning |
| `IETF-Draft.DidacticCapitalization` | Do not capitalize mid-word to explain an abbreviation | Style Guide Author Choice: "Use of didactic capitalization is not needed", e.g. `Extensible Markup Language (XML)` not `eXtensible`[^styleguide] | `existence` | suggestion |
| `IETF-Draft.NonAsciiPunctuation` | Use ASCII punctuation | Style Guide XML Formatting: "ASCII equivalents are to be used for punctuation (e.g., smart quotes and em dashes)"[^styleguide] | `existence` (`nonword`) | warning |

`IETF-Draft.Terms` is generated from the Terms list, and the generator must drop
the entries that defer to the document rather than settle the spelling. The
clearest case is Acknowledgments, where the list says "If consistent (with or
without the 'e'), leave it. If inconsistent, delete the 'e'."[^rpc-terms] — a
`consistency` check, not a substitution.

`IETF-Draft.NonAsciiPunctuation` excludes the `code` scope, and must not flag non-ASCII
characters used as examples, which RFC 7997 §3.1 permits.

## Group 2: example values

Reserved ranges exist so that documentation cannot collide with real deployments.
Each rule tests range membership, which regular expressions cannot express, so
these use Vale's Tengo `script` extension point.

| Rule | Requirement | Basis | Level |
|---|---|---|---|
| `IETF-Draft.ExampleIPv4` | IPv4 addresses in examples come from a documentation range | RFC 5737: "192.0.2.0/24 (TEST-NET-1), 198.51.100.0/24 (TEST-NET-2), and 203.0.113.0/24 (TEST-NET-3) are provided for use in documentation"; authors.ietf.org cites RFC 6890[^authors-examples] | warning |
| `IETF-Draft.ExampleIPv6` | IPv6 addresses in examples come from a documentation prefix | RFC 3849 (`2001:db8::/32`) together with RFC 9637 (`3fff::/20`) | warning |
| `IETF-Draft.ExampleASN` | AS numbers in examples come from a documentation range | RFC 5398: "64496 - 64511" and "65536 - 65551" | warning |
| `IETF-Draft.ExamplePhone` | Telephone numbers in examples are reserved for fictitious use | authors.ietf.org: "USA: +1-<area code>-555-<0100-0199>", "UK: +44-<geographic-area-code>-496-<0000-0999>"[^authors-examples] | warning |
| `IETF-Draft.ExampleMAC` | MAC and EUI values in examples come from the documentation range | RFC 9542 (BCP 141), which reserves `00-00-5E-00-53-xx` for documentation | warning |
| `IETF-Draft.ExampleDomains` | Example domain names use the reserved names | RFC 7322 §3.3: "DNS names … used as generic examples in RFCs should use the particular examples defined in 'Reserved Top Level DNS Names' [BCP32]"[^rfc7322] | warning |

`IETF-Draft.ExampleDomains` flags a **curated list of placeholder domains**
(`foo.com`, `mydomain.com`, `test.com`, `company.com` and similar), never "any
domain outside the reserved set". Drafts legitimately cite real sites, and a rule
that flags them would be withdrawn within a week.

`IETF-Draft.ExampleIPv4` must not flag version strings, netmasks, or addresses such as
`0.0.0.0` and `127.0.0.1` where they carry protocol meaning.

## Group 3: abbreviations and structure

| Rule | Requirement | Basis | Extends | Level |
|---|---|---|---|---|
| `IETF-Draft.Abbreviations` | Expand an abbreviation on first use, unless it is on the RFC Production Center's well-known list | authors.ietf.org: "Abbreviations should generally be expanded in parentheses. The RFC Production Center maintains a list of approved abbreviations that do not need to be expanded"[^authors-language]; RFC 7322 §3.6[^rfc7322] | `conditional` | warning |
| `IETF-Draft.AbstractCitations` | The Abstract contains no citations | RFC 7322 §4.3: "the Abstract must not contain citations"[^rfc7322] | `existence`, section-scoped | error |

The exception list for `IETF-Draft.Abbreviations` is **generated** from the entries
marked `*` in the Abbreviations list[^rpc-abbrev], per the evidence policy.

## Group 4: draft-stage only

These rules are correct for an Internet-Draft and wrong for a published RFC.
They carry the `Draft` prefix, and the style README instructs authors to disable
them once a document is in the RFC Editor queue.

| Rule | Requirement | Basis | Level |
|---|---|---|---|
| `IETF-Draft.DraftSelfReference` | An I-D does not refer to itself as an RFC | authors.ietf.org: "Your I-D must not refer to itself as an RFC or a draft RFC"[^authors-language] | error |
| `IETF-Draft.DraftTitleStatusWords` | The I-D title does not imply standards status | authors.ietf.org: "Avoid the use of the terms 'Standard', 'Proposed', 'Draft', 'Experimental', 'Historic', 'Required', 'Recommended', 'Elective', or 'Restricted' in the I-D title"[^authors-language] | warning |

`IETF-Draft.DraftSelfReference` needs an exception for the legitimate "this document
updates RFC NNNN" construction.

## Group 5: judgment-level

These implement guidance that is real but soft, or checks whose implementation
cannot be precise. All are `suggestion`, and each is a prompt to look, not a
claim of error.

| Rule | Requirement | Basis | Extends |
|---|---|---|---|
| `IETF-Draft.SpellingConsistency` | Spelling is internally consistent | RFC 7322 §3.1 (see the exclusion below)[^rfc7322] | `consistency` (level: warning) |
| `IETF-Draft.SectionTitleCase` | Section titles use title case | RFC 7322 §3.4, which also permits sentence-form titles[^rfc7322] | `capitalization` |
| `IETF-Draft.HttpsUris` | Prefer HTTPS URIs | Style Guide RECOMMENDED: "HTTPS URIs should be used when possible"[^styleguide] | `existence` |
| `IETF-Draft.DoubleNegatives` | Avoid double negatives | Style Guide RECOMMENDED: "Double negatives are discouraged"[^styleguide] | `existence` |
| `IETF-Draft.AngleBracketsURI` | Enclose bare URIs in angle brackets | RFC 7322 §3.3: "Angle brackets are strongly recommended around URIs"[^rfc7322] | `existence` |
| `IETF-Draft.Bcp14Boilerplate` | Uppercase key words are accompanied by the BCP 14 boilerplate and citation | RFC 8174 §2; RFC 7322 §4.8: "RFC 2119 must be cited … and included as a normative reference"[^bcp14][^rfc7322] | `conditional` (level: warning) |
| `IETF-Draft.Bcp14Lowercase` | Review lowercase key words for intent | RFC 8174: "The words have the meanings specified herein only when they are in all capitals"[^bcp14] | `existence`, **off by default** |
| `IETF-Draft.Bcp14Sparingly` | Do not use SHOULD merely to express a preference | IESG statement: key words "shouldn't be used merely to express a preference"[^iesg-bcp14] | `occurrence` |
| `IETF-Draft.InclusiveLanguage` | Review potentially biased terminology | RFC Editor RECOMMENDED, via the IESG statement, via NISTIR 8366 — see the evidence policy on referenced guidance | `substitution` |
| `IETF-Draft.StaleText` | Avoid text that dates the document | authors.ietf.org "Stale text": non-permanent URLs, "send comments to" a named list, assigning future work to a named WG[^authors-language] | `existence` |

`IETF-Draft.Bcp14Lowercase` ships disabled because lowercase key words are explicitly
legitimate under RFC 8174; the rule exists for authors who want to audit them,
and enabling it is the author's choice, not this style's assertion.

# Excluded checks

Recorded so they are not re-proposed.

**Two spaces after a period.** RFC 7322 §3.2 says "there must be two blank spaces
after the period", but this describes the plaintext rendering the RFC Production
Center produces, not Markdown or RFCXML source. Linting it in source flags
correct input.

**"Use American English."** RFC 7322 §3.1 says: "Spelling may be either American
or British, as long as an individual document is internally consistent."[^rfc7322]
The defensible check is `IETF-Draft.SpellingConsistency`, not a substitution rule that
rewrites British spellings. Both the Google and Microsoft Vale styles apply a US
preference that this style must not copy.

**Section ordering and required-section presence.** Owned by idnits[^idnits].

# Implementation constraints

Verified against Vale 3.21.0 rather than taken from documentation:

- `conditional` correctly flags an abbreviation used with no prior
  `Expansion (ABBR)`, which is the mechanism `IETF-Draft.Abbreviations` depends on.
- `script` (Tengo) works with `scope: raw` and `text.re_find`, appending
  `{begin, end}` matches. Its regular expressions are Go's RE2: no backreferences
  and no lookaround. The example-range rules therefore use `script` rather than a
  clever pattern.
- `capitalization` with `match: $title` and `style: Chicago` works on
  `scope: heading`.
- Vale lints `.txt` directly. **`.xml` requires** a `Transform` key naming an
  XSLT 1.0 stylesheet plus `xsltproc` on PATH; without it Vale halts with "no
  XSLT transform provided". RFCXML support therefore needs a stylesheet shipped
  from this repository, and Markdown is the primary supported input until then.

[^rfc7322]: [RFC 7322](https://www.rfc-editor.org/rfc/rfc7322.html), RFC Style Guide.
[^styleguide]: [Updates to the RFC Style Guide](https://www.rfc-editor.org/styleguide/).
[^authors-language]: [Language and style](https://authors.ietf.org/language-and-style), Internet-Draft Author Resources.
[^authors-examples]: [Example addresses](https://authors.ietf.org/example-addresses), Internet-Draft Author Resources.
[^bcp14]: [BCP 14](https://www.rfc-editor.org/info/bcp14) — RFC 2119 and RFC 8174.
[^rpc-terms]: [RFC-Specific Terms list](https://rpc-wiki.rfc-editor.org/doku.php?id=terms).
[^rpc-abbrev]: [Abbreviations list](https://rpc-wiki.rfc-editor.org/doku.php?id=abbrev_list).
[^iesg-bcp14]: [IESG Statement on clarifying the use of BCP 14 key words](https://datatracker.ietf.org/doc/statement-iesg-statement-on-clarifying-the-use-of-bcp-14-key-words/), 17 March 2025.
[^idnits]: [idnits](https://github.com/ietf-tools/idnits).
