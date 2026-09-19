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
  - id: bcp22
    resource: https://www.rfc-editor.org/rfc/rfc2360.html
    title: "RFC 2360 (BCP 22): Guide for Internet Standards Writers"
  - id: idnits
    resource: https://github.com/ietf-tools/idnits
    title: idnits, the Internet-Draft nits checker
  - id: draftforge
    resource: https://github.com/ietf-tools/draftforge
    title: "IETF DraftForge: VS Code extension to write, review, refine and submit Internet-Drafts"
  - id: rpc-abbrev-json
    resource: https://github.com/rfc-editor-drafts/common/blob/main/abbreviations.json
    title: abbreviations.json, the RPC's machine-readable abbreviation data
  - id: rpc-names-json
    resource: https://github.com/rfc-editor-drafts/common/blob/main/names.json
    title: names.json, author names with preferred forms
---

This specifies the `IETF-Draft` Vale style: the rules it contains, the guidance each
rule implements, and the boundaries of what it claims. Severity levels follow the
[evidence policy](/evidence-policy.md). The companion
[email style](/email-style.md) covers mailing list prose and shares no rules with
this one.

# Scope

**In scope**: prose in Internet-Drafts and in documents heading for RFC
publication, in Markdown (including kramdown-rfc), RFCXML, and plain text.

**Out of scope**, because another tool owns them and duplicating them badly is
worse than leaving them alone:

| Concern | Owner |
|---|---|
| IPR and copyright boilerplate, expiry statement, required-section presence, outdated references, downrefs, line and page length | idnits[^idnits] and datatracker submission checks |
| RFCXML validity, reference resolution, generated boilerplate | xml2rfc, kramdown-rfc |
| Reference entry formatting, title-page header, "Status of This Memo" text | the RFC Production Center at publication |

Out of scope does not mean unchecked. The author's goal is to hear everything a
reviewer familiar with IETF rules would raise, which no single tool provides:
idnits owns the structural half, and this style owns the prose half. The
[tool landscape](/tool-landscape.md) maps who checks what, names the gaps that
belong to nobody, and specifies `make check-draft`, which runs idnits and Vale
over one document and prints one report.

# Conformance

A document conforms to this style when Vale reports no `error`-level alerts from
it. `warning` and `suggestion` alerts are review prompts, not defects.

Passing this style does **not** mean a document is ready for submission, has
correct boilerplate, or will survive RFC Editor review. It means the prose
contains none of the specific, cited defects enumerated below.

# Normative basis

**A note on standing.** Most of this style rests on RFC Editor guidance, and
RFC 7322 and RFC 7997 are **IAB stream** documents: editorial policy for the RFC
Series, not IETF consensus. That is the right authority for how an RFC is
written, and it is not the same claim as "the IETF agreed this". Where an
IETF-stream document covers the same ground, it is cited in preference, as
`Concision` cites BCP 22. The stream of every cited document is recorded in the
[evidence policy](/evidence-policy.md).

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
  rules are generated from, and the machine-readable
  `abbreviations.json`[^rpc-abbrev-json] and `names.json`[^rpc-names-json] that
  the RPC's own tooling runs on.
Official IETF tooling, meaning idnits[^idnits] and DraftForge[^draftforge], is
**not** in that list. Those tools corroborate that a requirement is enforced in practice
and set a floor for what an author will hear from a reviewer, and their source is
a useful implementation reference, but a rule here traces to published guidance
or it does not ship. See the [tool landscape](/tool-landscape.md).

# Rule inventory

**Shipped**: `Terms`, `CitationSpacing`, `RFCCompoundHyphen`, `AngleBracketsURI`,
`AbbreviationVerbs`, `Abbreviations`, `DidacticCapitalization`,
`NonAsciiPunctuation`, `DraftSelfReference`, `HttpsUris`, `InclusiveLanguage`,
`Concision`, `AbstractCitations`, `DraftTitleStatusWords`, `ExampleIPv4`,
`ExampleIPv6`, `ExampleDomains`, `StaleText`, `DoubleNegatives`,
`SectionTitleCase`, `SpellingConsistency`, `Bcp14Lowercase`, and
`Bcp14Sparingly`, and `ProtocolNames`. The
`Bcp14Boilerplate`, `CitationReferences`, `ExampleASN`, `ExamplePhone`, and
`ExampleMAC` rules are also shipped. The
inventory below records both shipped and planned rules.

Each rule is one PR. `Level` is the target level under the evidence policy. A
rule whose alert rate makes it unsupportable ships one level lower, or disabled
by default; it is not dropped for being rarely triggered, since a clean corpus
says nothing about whether the guidance exists. See the
[validation methodology](/validation.md).

## Group 0: requirements that official tooling already enforces

Every rule below implements a **documented requirement**. The tooling column
records whether idnits or DraftForge also checks it, because a requirement an
IETF-aware reviewer's tools already flag is one an author will certainly hear
about. The tool is corroboration, not the source. Where a tool covers less
than the requirement, these rules cover the requirement.

| Rule | Requirement and source | Also checked by | Extends | Level |
|---|---|---|---|---|
| `IETF-Draft.InclusiveLanguage` | The full NISTIR 8366 Table 1, roughly two dozen terms, which the IESG statement points to and the RFC Editor marks RECOMMENDED[^styleguide] | DraftForge, 7 of those terms | `substitution` | suggestion |
| `IETF-Draft.Abbreviations` | authors.ietf.org: "Abbreviations should generally be expanded in parentheses"[^authors-language]; RFC 7322 §3.6[^rfc7322] | DraftForge `abbreviations` | `conditional` | warning |
| `IETF-Draft.RFCTerms` | RFC 7322 §3.4: "Capitalization must be consistent within the document and ideally should be consistent with related RFCs"[^rfc7322] | DraftForge `rfc-terms`, `inconsistent-capitalization` | `existence` | suggestion |
| `IETF-Draft.ProtocolNames` | A narrow list of unambiguous protocol-name spellings, consistent with RFC 7322 §3.4[^rfc7322] | DraftForge `rfc-terms` | `substitution` | suggestion |
| `IETF-Draft.Names` | RFC 7322 §4.12 on author names, and the RPC's recorded preferred forms[^rpc-names-json] | DraftForge `names` | `existence` | suggestion |

The authoritative exception data for `IETF-Draft.Abbreviations` is
`abbreviations.json`[^rpc-abbrev-json], the RPC's machine-readable data file.
Run `make update-abbreviations` to regenerate the checked-in exceptions from
that source.

`IETF-Draft.ProtocolNames` deliberately covers only unambiguous capitalization
forms such as `HTTP`, `TLS`, `IPv6`, and `WebSocket`. It does not attempt to
infer whether an arbitrary word is a protocol name, or whether a document's
chosen capitalization agrees with a related RFC. Those broader checks remain
covered by issue [#9](https://github.com/OR13/ietf-vale/issues/9).

`IETF-Draft.InclusiveLanguage` is generated from NISTIR 8366 Table 1. Its
message names the chain of authority (RFC Editor RECOMMENDED, via the IESG
statement, via NIST) so a reader can see at a glance that it is encouragement,
not IETF consensus.

## Group 0b: tooling conventions with no documented requirement

These implement checks that official tooling performs but no published guidance
states. They are useful, and an author may well hear them from a reviewer using
DraftForge, but they are **not requirements**. The shipped optional checks live
in `IETF-Draft-Optional`, a separate package that users must add explicitly;
the remaining entries are not shipped. See the
[evidence policy](/evidence-policy.md).

| Rule | What it checks | Upstream | Why it is not a requirement |
|---|---|---|---|
| `IETF-Draft.Typos` | 62 dictionary entries covering 79 observed misspellings | DraftForge `typos.js` | A list of mistakes people made, not guidance anyone published |
| `IETF-Draft-Optional.Articles` | A conservative set of `a`/`an` article errors | DraftForge `articles.js` | General English grammar; RFC 7322 says only that the publication language is English; opt-in package |
| `IETF-Draft-Optional.RepeatedWords` | A word repeated across a line break | DraftForge `repeated-words.js` | Same; opt-in package |

**`Placeholders` is deliberately absent.** DraftForge flags `RFCXXXX`, `RFCTBD`
and similar because it is used during final review, where such a placeholder must
be resolved before publication. In an Internet-Draft the same string is correct:
the number is not assigned yet, and authors leave a note for the RFC Editor. The
check belongs to a workflow stage this style does not lint, and adopting it would
turn correct text into an error. If it is ever added, it belongs with the
draft-stage rules in Group 4, inverted: flag a placeholder that has **no**
accompanying RFC Editor note.

## Group 0c: writing quality, from consensus guidance

Most of this style rests on RFC Editor guidance, which is **IAB stream**
editorial policy rather than IETF consensus. One rule rests on an IETF-stream
BCP, which makes it the easiest rule here to defend.

| Rule | Requirement and source | Extends | Level |
|---|---|---|---|
| `IETF-Draft.Concision` | RFC 2360 (**BCP 22**, IETF stream) §2.4: "Concise text has several advantages. It makes the document easier to read. Such text reduces the chance for conflict between different portions of the specification."[^bcp22] Reinforced by the RFC Editor: "authors should focus on using clear, concise language", pointing at Richard Lanham's Paramedic Method "for information on revising verbose text"[^styleguide] | `substitution` | suggestion |

The rule lists only fixed phrases with an unambiguous shorter form, and it never
claims a passage is too long. BCP 22 §2.4 is explicit that the trade-off runs
both ways: "Longer descriptions may be necessary to explain purpose, background",
and a rule that flagged length would be taking a side the source does not take.

Measured over 60 Internet-Drafts at `-00`, wordy constructions appear roughly
2.2 times per 1000 lines.

## Group 1: mechanical

These have precise triggers and, by inspection of their patterns, near-zero
scope for matching text their sources do not describe.

| Rule | Requirement | Basis | Extends | Level |
|---|---|---|---|---|
| `IETF-Draft.Terms` | Use the RFC Production Center's settled spelling of RFC-specific terms, except where the list defers to internal consistency | Terms list: `IPsec` not `IPSEC`; `email` not `e-mail`; `online` not `on-line`; `timestamp` not `time-stamp`; `ASCII` not `US-ASCII`[^rpc-terms] | `substitution` | warning |
| `IETF-Draft.CitationSpacing` | A citation tag contains no spaces | RFC 7322 §3.5: "A citation/reference tag must not contain spaces", e.g. `[RFC2119]` not `[RFC 2119]`[^rfc7322] | `substitution` | error |
| `IETF-Draft.CitationReferences` | Every bracketed citation has a matching reference definition | RFC 7322 §4.8, which requires cited documents to appear in the References section[^rfc7322] | `script` | error |
| `IETF-Draft.AngleBracketsURI` | Enclose bare URIs in angle brackets | RFC 7322 §3.3: "Angle brackets are strongly recommended around URIs"[^rfc7322] | `script` | warning |
| `IETF-Draft.RFCCompoundHyphen` | Do not form compounds by hyphenating an RFC citation | Style Guide RECOMMENDED: "Avoid forming compounds by hyphenating RFC numbers", e.g. `[RFC5011]-style rollover`[^styleguide] | `existence` | warning |
| `IETF-Draft.AbbreviationVerbs` | Affix suffixes to abbreviations without punctuation | Style Guide RECOMMENDED: `XORed` not `XOR'ed`, `NATed` not `NAT-ed`[^styleguide] | `existence` | warning |
| `IETF-Draft.DidacticCapitalization` | Do not capitalize mid-word to explain an abbreviation | Style Guide Author Choice: "Use of didactic capitalization is not needed", e.g. `Extensible Markup Language (XML)` not `eXtensible`[^styleguide] | `existence` | suggestion |
| `IETF-Draft.NonAsciiPunctuation` | Use ASCII punctuation | Style Guide XML Formatting: "ASCII equivalents are to be used for punctuation (e.g., smart quotes and em dashes)"[^styleguide] | `existence` (`nonword`) | warning |

`IETF-Draft.Terms` is generated from the Terms list, and the generator must drop
the entries that defer to the document rather than settle the spelling. The
clearest case is Acknowledgments, where the list says "If consistent (with or
without the 'e'), leave it. If inconsistent, delete the 'e'."[^rpc-terms] That is
a `consistency` check, not a substitution.

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
| `IETF-Draft.ExampleDomains` | Example domain names use the reserved names | RFC 7322 §3.3: "DNS names ... used as generic examples in RFCs should use the particular examples defined in 'Reserved Top Level DNS Names' [BCP32]"[^rfc7322] | warning |

`IETF-Draft.ExampleDomains` duplicates idnits, which already enforces reserved
example domains through `FQDN_EXAMPLE_RE` and raises `INVALID_DOMAIN_TLD`. It
earns its place only by running at authoring time, in Markdown, before a draft is
submitted. It flags a **curated list of placeholder domains** (`foo.com`,
`mydomain.com`, `test.com`, `company.com` and similar), never "any domain outside
the reserved set". Drafts legitimately cite real sites, and a rule that flags
them would be withdrawn within a week.

The IP rules fill a genuine gap: **idnits checks whether an address is
syntactically valid, not whether it comes from a documentation range.** Nothing
in the official tooling asks that question today.

`IETF-Draft.ExampleIPv4` must not flag version strings, section numbers, netmasks,
or addresses such as `0.0.0.0` and `127.0.0.1` where they carry protocol meaning.
idnits solves the section-number problem with the lookbehind `(?<![0-9a-zA-Z]+\.)`,
which Vale's RE2 engine cannot express; see the
[tool landscape](/tool-landscape.md) for what to do instead.

## Group 3: structure

| Rule | Requirement | Basis | Extends | Level |
|---|---|---|---|---|
| `IETF-Draft.AbstractCitations` | The Abstract contains no citations | RFC 7322 §4.3: "the Abstract must not contain citations"[^rfc7322] | `script`, section-scoped | error |

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
| `IETF-Draft.SpellingConsistency` | Spelling is internally consistent for a curated set of common variants | RFC 7322 §3.1 (see the exclusion below)[^rfc7322] | `script` (level: suggestion) |
| `IETF-Draft.SectionTitleCase` | Section titles use title case | RFC 7322 §3.4, which also permits sentence-form titles[^rfc7322] | `capitalization` |
| `IETF-Draft.HttpsUris` | Prefer HTTPS URIs | Style Guide RECOMMENDED: "HTTPS URIs should be used when possible"[^styleguide] | `existence` |
| `IETF-Draft.DoubleNegatives` | Avoid double negatives | Style Guide RECOMMENDED: "Double negatives are discouraged"[^styleguide] | `existence` |
| `IETF-Draft.Bcp14Boilerplate` | Uppercase key words are accompanied by the BCP 14 boilerplate and citation | RFC 8174 §2; RFC 7322 §4.8: "RFC 2119 must be cited ... and included as a normative reference"[^bcp14][^rfc7322] | `script` (level: warning) |
| `IETF-Draft.Bcp14Lowercase` | Review lowercase key words for intent | RFC 8174: "The words have the meanings specified herein only when they are in all capitals"[^bcp14] | `existence` |
| `IETF-Draft.Bcp14Sparingly` | Review SHOULD when nearby wording indicates a preference | IESG statement: key words "shouldn't be used merely to express a preference"[^iesg-bcp14] | `existence` |
| `IETF-Draft.StaleText` | Avoid text that dates the document | authors.ietf.org "Stale text": non-permanent URLs, "send comments to" a named list, assigning future work to a named WG[^authors-language] | `existence` |

`IETF-Draft.Abbreviations` uses Vale's `conditional` extension. It checks for an
expansion in the same lint block, which keeps the rule predictable across Vale's
supported input formats. It does not claim full document-wide first-use analysis.

`IETF-Draft.Bcp14Lowercase` is a suggestion because lowercase key words are
explicitly legitimate under RFC 8174. The alert asks authors to review intent;
it does not claim that lowercase usage is an error.

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

- **Vale does not match a phrase across a line break.** `due to the fact that`
  is found when it sits on one line and missed when the text wraps between
  `due to` and `the fact that`, in Markdown and plain text alike. Internet-Drafts are
  hard-wrapped, so every multi-word rule loses some recall: measured over 60
  drafts at `-00`, 110 of 119 multi-word phrases sat on one line, a recall of
  92%. Authors who write one sentence per line, or who let their editor soft-wrap,
  lose nothing.
- **Vale wraps `existence` and `substitution` tokens in `\b`.** A pattern that
  begins or ends with a non-word character, such as a `[RFC1234]` citation tag,
  can therefore never match without `nonword: true`. This cost two rules a
  silent no-op before it was caught by reading their recorded expectations.
- **Markdown hides bracketed citations and bare URIs from the text scope.**
  `[RFC 2119]` is parsed as a link reference and `http://example.com` is
  autolinked, so neither reaches a rule scoped to prose. `CitationSpacing`,
  `RFCCompoundHyphen` and `HttpsUris` use `scope: raw` for this reason, which
  also means they see fenced code, a trade-off recorded in each rule.

- `conditional` correctly flags an abbreviation when its lint block has no
  matching `Expansion (ABBR)`, which is the mechanism
  `IETF-Draft.Abbreviations` depends on. The rule deliberately documents this
  block-level boundary rather than claiming document-wide first-use analysis.
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
[^bcp14]: [BCP 14](https://www.rfc-editor.org/info/bcp14), which is RFC 2119 and RFC 8174.
[^rpc-terms]: [RFC-Specific Terms list](https://rpc-wiki.rfc-editor.org/doku.php?id=terms).
[^rpc-abbrev]: [Abbreviations list](https://rpc-wiki.rfc-editor.org/doku.php?id=abbrev_list).
[^iesg-bcp14]: [IESG Statement on clarifying the use of BCP 14 key words](https://datatracker.ietf.org/doc/statement-iesg-statement-on-clarifying-the-use-of-bcp-14-key-words/), 17 March 2025.
[^bcp22]: [RFC 2360](https://www.rfc-editor.org/rfc/rfc2360.html), BCP 22, Guide for Internet Standards Writers. IETF stream, never obsoleted or updated.
[^idnits]: [idnits](https://github.com/ietf-tools/idnits), v3 branch.
[^draftforge]: [ietf-tools/draftforge](https://github.com/ietf-tools/draftforge), whose checks are documented at [draftforge.ietf.org/checks](https://draftforge.ietf.org/checks/).
[^rpc-abbrev-json]: [abbreviations.json](https://github.com/rfc-editor-drafts/common/blob/main/abbreviations.json).
[^rpc-names-json]: [names.json](https://github.com/rfc-editor-drafts/common/blob/main/names.json).
