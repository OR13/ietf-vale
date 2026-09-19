---
type: Reference
title: Tool landscape
description: What idnits, DraftForge, xml2rfc and this style each own, and how an author runs them together so nothing falls between them.
tags: [tooling, idnits, draftforge, coverage]
status: draft
timestamp: 2026-09-19T00:00:00Z
generated: { by: claude/opus-5, at: 2026-09-19T00:00:00Z }
sources:
  - id: idnits
    resource: https://github.com/ietf-tools/idnits
    title: "idnits v3 (ietf-tools), the Internet-Draft nits checker"
  - id: draftforge
    resource: https://draftforge.ietf.org/
    title: "IETF DraftForge: VS Code extension to write, review, refine and submit Internet-Drafts"
  - id: df-checks
    resource: https://draftforge.ietf.org/checks/
    title: DraftForge validation checks
  - id: rpc-common
    resource: https://github.com/rfc-editor-drafts/common
    title: "rfc-editor-drafts/common: reference configs and data for editor tools and automations"
  - id: rpc-abbrev-json
    resource: https://github.com/rfc-editor-drafts/common/blob/main/abbreviations.json
    title: abbreviations.json (3277 entries, 279 marked well-known)
  - id: rpc-names-json
    resource: https://github.com/rfc-editor-drafts/common/blob/main/names.json
    title: names.json (author names with preferred forms)
---

The goal of the [draft style](/draft-style.md) is that running it over a draft
warns the author about what a reviewer familiar with IETF rules would raise.
That reviewer is increasingly backed by tooling, so this concept records what the
official tools check, what they do not, and where this style fits.

# The official tools

## idnits v3

`ietf-tools/idnits`[^idnits] is the checker the submission process uses. Its v3
modules are `downref`, `filename`, `fqdn`, `ip`, `keywords`, `metadata`, `raw`,
`sections`, `txt`, and `xml`. It owns everything structural, and its rule
identifiers are specific enough to cite:

| Area | Examples of what idnits raises |
|---|---|
| Sections | `MISSING_ABSTRACT_SECTION`, `MISSING_SECURITY_CONSIDERATIONS_SECTION`, `MISSING_REFERENCES_SUBSECTIONS`, `REFERENCE_NOT_USED`, `TOO_MANY_AUTHORS` |
| Boilerplate and metadata | `COPYRIGHT_LINE_MISSING`, `OBSELETE_COPYRIGHT_LINE`, `EXPIRES_LINE_MISSING`, `SUBMISSION_COMPLIANCE_LINE_MISSING` |
| BCP 14 | `MISSING_REQLEVEL_BOILERPLATE`, `MISSING_REQLEVEL_REF`, `INVALID_REQLEVEL_KEYWORD`, `INCORRECT_KEYWORD_SPELLING` |
| Example values | `INVALID_DOMAIN_TLD`, `POSSIBLE_INVALID_TLD`, `INVALID_ARPA_DOMAIN`, `INVALID_IPV4_ADDRESS`, `INVALID_IPV6_ADDRESS` |
| Text formatting | `LINE_TOO_LONG`, `PAGE_TOO_LONG`, `RAGGED_RIGHT` |
| References | `DOWNREF`, obsolete and unused references |

Two details matter for this style:

- **idnits checks IP address _validity_, not documentation-range membership.**
  `ip.mjs` tests whether a string parses as an IPv4 or IPv6 address and raises
  `INVALID_IPV4_ADDRESS` when it does not. It does not ask whether the address
  comes from RFC 5737 or RFC 3849. That gap is real and belongs to this style.
- **idnits does enforce reserved example domains**, via
  `FQDN_EXAMPLE_RE = /^(?:[a-z0-9_-]+\.)*example(?:\.(?:com|org|net))?$/i`. A
  rule here would duplicate it, so `IETF-Draft.ExampleDomains` exists only to
  catch placeholder domains at authoring time, in Markdown, before a draft is
  ever submitted.

## DraftForge

DraftForge[^draftforge] is the RFC Production Center's VS Code extension for
writing and reviewing drafts. Its **validation checks**[^df-checks] are the
closest published statement of what prose issues IETF tooling flags today:

| DraftForge check | What it looks for |
|---|---|
| Articles | `a` before a vowel, `an` before a consonant, with an exception list; `an LF` |
| Hyphenation | Inconsistent hyphenation of the same term (`abc-def` beside `abcdef`) |
| Inclusive Language | 8 dictionary entries covering 56 terms, with suggested alternates |
| Names | Author names that may be misspelled or not in the author's preferred form |
| Non-ASCII | Any character outside `x00-x7F` |
| Placeholders | `RFCTBD`, `RFCTBA`, `RFCXX`, `RFCYY`, `RFCNN`, `RFCMM`, `RFC0000`, `RFCTODO` |
| Typos | 62 dictionary entries covering 79 terms, including invalid BCP 14 terms and duplicated punctuation |

Its `src/commands/` directory carries several more prose checks than the
documentation site lists, and they are equally worth parity:

| DraftForge command | What it looks for |
|---|---|
| `repeated-words` | A word repeated across a line break, via `/\b(\w+)(?=\s)[^\S\n]*\n?[^\S\n]*\1\b/gi` |
| `rfc-terms` | IETF terms whose usage and capitalization authors get wrong: `RFC series`, `working group`, `standards track`, `last call`, `area director`, `chair`, `shepherd`, `proposed standard`, `internet draft`, and the stream names |
| `inconsistent-capitalization` | The same term capitalized two ways in one document |
| `inconsistent-formatting` | The same term formatted two ways in one document |
| `abbreviations` | Abbreviations in the document, listed against the RPC data |
| `surround-bcp14-keywords` | Marking BCP 14 key words up correctly in RFCXML |

DraftForge states plainly that a match is advisory: for inclusive language, "It's
up to the draft authors and RPC staff to determine whether a term should be left
as-is or find a better suited alternative"; for names, "A match doesn't
automatically signify an error".

**DraftForge is an editor extension, not a command-line checker.** An author who
does not use VS Code, or who wants these checks in CI or a pre-commit hook, has
no way to run them. That is the gap this style fills, which is why parity with
these checks is the draft style's first priority rather than an afterthought.

## RPC data files

`rfc-editor-drafts/common`[^rpc-common] publishes the data the RPC's own tooling
runs on, as JSON, in git:

- `abbreviations.json`[^rpc-abbrev-json] holds 3277 entries of
  `{term, full, wellknown?, note?}`, of which 279 are marked well-known. This
  supersedes scraping the wiki page: it is the same data, versioned, with the
  well-known flag machine-readable.
- `names.json`[^rpc-names-json] holds `{description, match[]}` entries recording
  author names whose preferred form differs from the obvious one.

Rules generated from these files are regenerated, never hand-edited, and the
generator records the upstream commit it read.

## xml2rfc and kramdown-rfc

Own document validity: schema conformance, reference resolution, generated
boilerplate. Nothing here overlaps.

# Division of labor

| Concern | Owner | This style |
|---|---|---|
| Required sections, boilerplate, expiry, copyright, downrefs, unused references, line and page length | idnits | Nothing. Running idnits is a prerequisite, not a substitute |
| BCP 14 boilerplate presence, key word spelling | idnits | A prose-level prompt only where idnits is silent, such as the IESG guidance on using SHOULD merely as a preference |
| Reserved example domains | idnits | Authoring-time placeholders in Markdown, before submission |
| IP addresses outside the documentation ranges | **nobody** | `IETF-Draft.ExampleIPv4`, `IETF-Draft.ExampleIPv6` |
| AS numbers, telephone numbers, MAC addresses outside the documentation ranges | **nobody** | `IETF-Draft.ExampleASN`, `IETF-Draft.ExamplePhone`, `IETF-Draft.ExampleMAC` |
| Articles, hyphenation consistency, inclusive language, names, non-ASCII, placeholders, typos | DraftForge, in VS Code only | Parity, as a CLI usable in any editor and in CI |
| Abbreviation expansion on first use | DraftForge (List Abbreviations) | `IETF-Draft.Abbreviations`, generated from `abbreviations.json` |
| RFC-specific term spellings, RFC Editor RECOMMENDED guidance | RPC, at AUTH48 | The rules in the draft style, including the narrow `ProtocolNames` check |
| Document validity | xml2rfc, kramdown-rfc | Nothing |

# Running them together

The point of the coverage map is that an author runs **both**. A planned
`make check-draft` runs idnits and Vale over the same document and prints one
report, so the structural and prose halves arrive together. Vale alone is not a
submission check, and the draft style's README says so.

# Implementation constraint: the patterns cannot be copied

Both official tools are JavaScript, whose regular expressions support lookbehind,
lookahead, and backreferences. **Vale's `script` rules use Go's `regexp`, which
is RE2 and supports none of the three.** Three upstream patterns are therefore
off limits as written, and each needs a different mechanism to reach the same
result:

| Upstream pattern | Why it cannot be copied | What to use instead |
|---|---|---|
| idnits `IPV4_LOOSE_RE`, guarded by `(?<![0-9a-zA-Z]+\.)` to skip section numbers such as `7.2.1.1` | Lookbehind | Match candidates, then filter them in Tengo |
| DraftForge `articles`, five patterns built almost entirely from lookarounds and exception lists | Lookbehind and lookahead | Tengo, with the exception lists carried over verbatim |
| DraftForge `repeated-words`, `/\b(\w+)...\1\b/` | Backreference | Vale's native `repetition` extension point, which exists for exactly this |

A rule that silently drops one of these guards is a **rule defect**, not a
difference of opinion: without the lookbehind, the IPv4 rule flags every section
number in the document. Each guard belongs in the fixture as a near-miss, so the
protection is pinned rather than assumed.

[^idnits]: [ietf-tools/idnits](https://github.com/ietf-tools/idnits), v3 branch.
[^draftforge]: [IETF DraftForge](https://draftforge.ietf.org/).
[^df-checks]: [DraftForge validation checks](https://draftforge.ietf.org/checks/).
[^rpc-common]: [rfc-editor-drafts/common](https://github.com/rfc-editor-drafts/common).
[^rpc-abbrev-json]: [abbreviations.json](https://github.com/rfc-editor-drafts/common/blob/main/abbreviations.json).
[^rpc-names-json]: [names.json](https://github.com/rfc-editor-drafts/common/blob/main/names.json).
