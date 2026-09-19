---
type: Policy
title: Evidence policy
description: How a rule earns its severity level, what every rule PR must show, and how cited sources are kept current.
tags: [policy, review, provenance]
status: draft
timestamp: 2026-09-19T00:00:00Z
generated: { by: claude/opus-5, at: 2026-09-19T00:00:00Z }
sources:
  - id: styleguide
    resource: https://www.rfc-editor.org/styleguide/
    title: "Updates to the RFC Style Guide (RFC Editor, living document)"
  - id: rfc-index
    resource: https://www.rfc-editor.org/rfc-index.xml
    title: RFC index (authoritative status and obsoletion data)
  - id: rpc-terms
    resource: https://rpc-wiki.rfc-editor.org/doku.php?id=terms
    title: RFC-Specific Terms list (RFC Production Center wiki)
  - id: rpc-abbrev
    resource: https://rpc-wiki.rfc-editor.org/doku.php?id=abbrev_list
    title: Abbreviations list (RFC Production Center wiki)
  - id: iesg-inclusive
    resource: https://datatracker.ietf.org/doc/statement-iesg-iesg-statement-on-inclusive-language-20210511
    title: IESG Statement on Inclusive Language (2021-05-11)
---

This policy applies to both the [draft style](/draft-style.md) and the
[email style](/email-style.md), and pairs with the
[validation methodology](/validation.md), which governs what evidence shows a
rule is correct. It exists to answer one question before any rule
merges: when a participant asks "who says so?", what do we point at?

# Severity tiers

A rule's level is a function of how strongly its source states the guidance, not
of how strongly the author feels about it. The RFC Editor already tiers its own
guidance into MUSTs, RECOMMENDED, and Author Choice[^styleguide]; this policy
maps those tiers, plus IETF document status, onto Vale levels.

| Tier | Source class | Vale level |
|---|---|---|
| A | IETF consensus (BCP or Standards Track), or an RFC Editor **MUST**, mechanically checkable in prose | `error` |
| B | RFC Editor **RECOMMENDED**, or an authors.ietf.org "must"/"should" for Internet-Drafts | `warning` |
| C | RFC Editor **Author Choice**, an IESG statement, or guidance the IETF points to but did not produce | `suggestion` |
| D | Guidance about participant **behavior**, which a linter can only approximate | `suggestion`, opt-in |

# Official tooling is not a source

idnits and DraftForge implement checks; they do not define requirements. A check
in one of those tools is **corroboration that a requirement is enforced in
practice, and a floor for coverage** -- never the authority for a rule. Three
things follow, and each has already bitten this specification:

1. **A tool may under-cover a requirement.** DraftForge's inclusive-language
   dictionary holds 7 entries; NISTIR 8366 Table 1, which the IESG statement
   actually points to, holds roughly two dozen. A rule built for parity would
   cover under a third of the guidance. The rule is built from the requirement.
2. **A tool may check something no requirement states.** DraftForge's typo
   dictionary is 62 observed mistakes, and its article checks are English
   grammar. Both are useful; neither traces to IETF guidance. Such a rule is
   Tier C at best, ships disabled by default, and says in its own documentation
   that it implements a tooling convention rather than a requirement.
3. **A tool encodes the workflow it was built for.** DraftForge flags `RFCXXXX`
   as a placeholder because it is used in final review, where such a placeholder
   must be resolved. In an Internet-Draft the same string is correct: numbers are
   not assigned yet, and authors leave a note for the RFC Editor. Copying the
   check without its context turns correct text into an error.

Where a tool states its own confidence, that caps ours: DraftForge calls its
inclusive-language and names matches advisory, so no rule derived from the same
guidance ships above `suggestion`.

Two constraints override the table:

1. **No rule reaches `error` unless its source uses mandatory language _and_ the
   check has no plausible false positive in a well-formed document.** A Tier A
   source with a fuzzy implementation ships as `warning`.
2. **A rule never claims more than its source.** A rule that flags a lowercase
   "should" does not assert the text is wrong; its message says what to check,
   not what to fix.

# Currency of sources

Citing an obsolete document is the fastest way to lose an argument about a rule.
Source status is verified against the RFC index[^rfc-index], never from memory.
The following successors were confirmed on 2026-09-19 and are the versions this
project cites:

| Superseded | Current |
|---|---|
| RFC 3005 (IETF Discussion List Charter) | **RFC 9245**, BCP 45 |
| RFC 3934 (mailing list management) | **RFC 9945**, BCP 245 |
| RFC 7042 (IEEE OUI/EUI usage) | **RFC 9542**, BCP 141 |
| RFC 3849 alone (IPv6 documentation prefix) | RFC 3849 **and** RFC 9637 |
| RFC 7776 alone (anti-harassment) | **RFC 8716**, BCP 25 |
| The former "I-D Guidelines" page | authors.ietf.org content guidelines |

A scheduled `check-sources` job re-resolves every `link:` in every rule against
the RFC index and fails when a cited document has been obsoleted. The same job
hashes the two RFC Production Center wiki pages[^rpc-terms][^rpc-abbrev] and
fails when they change, because rules generated from them are then stale.

# Generated rules

Rules derived from an upstream list are **generated, not transcribed**. The
generator writes a header naming the source URL and the fetch date. Hand-editing
a generated rule file is a review failure: the fix belongs in the generator or
upstream.

# Evidence required in every rule PR

This extends the checklist in `CONTRIBUTING.md`. A rule PR shows:

1. **Source**, by stable URL, with the **verbatim sentence** the rule implements.
   A tool's source code is not this sentence. If no documented sentence exists,
   say so, and the rule ships disabled by default as a tooling convention.
2. **Standing**: the BCP number and current status from the RFC index, or the
   RFC Editor tier (MUST / RECOMMENDED / Author Choice), or "IESG statement", or
   "non-IETF guidance the IETF points to". State it; do not imply it.
3. **Not obsoleted**: confirmed, naming the successor if the obvious source has
   one.
4. **Tier and level**: which row of the table above, and why the level matches.
5. **Audience**: draft style, email style, or draft-stage-only, and why the other
   audience is unaffected.
6. **Behavior in real text**: a sample of corpus alerts classified as deviation,
   rule defect, or legitimate exception, plus the alert rate per 1000 lines and
   the near-misses in the fixture that pin the boundary. See the
   [validation methodology](/validation.md), which also records the two corpus
   inferences that are not admissible.
7. **Overreach check**: one sentence naming what the rule does *not* claim.

# Guidance the IETF points to but did not write

Some guidance reaches IETF participants by reference rather than by consensus.
The clearest case is inclusive language: the RFC Editor marks it RECOMMENDED
because each stream has chosen to follow the IESG statement[^iesg-inclusive],
and that statement "encourages IETF participants to use the guidance in
[NISTIR8366]" -- a NIST document, not an IETF consensus document.

Such guidance is Tier C. Rules implementing it ship in their own file so they can
be disabled in one line, and their documentation states the chain of authority
plainly rather than presenting NIST guidance as IETF consensus.

[^styleguide]: [Updates to the RFC Style Guide](https://www.rfc-editor.org/styleguide/), which states "This page reflects current usage."
[^rfc-index]: [RFC index](https://www.rfc-editor.org/rfc-index.xml).
[^rpc-terms]: [RFC-Specific Terms list](https://rpc-wiki.rfc-editor.org/doku.php?id=terms).
[^rpc-abbrev]: [Abbreviations list](https://rpc-wiki.rfc-editor.org/doku.php?id=abbrev_list).
[^iesg-inclusive]: [IESG Statement on Inclusive Language](https://datatracker.ietf.org/doc/statement-iesg-iesg-statement-on-inclusive-language-20210511), 11 May 2021.
