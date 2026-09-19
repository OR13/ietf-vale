---
type: Policy
title: Validation methodology
description: What evidence shows a rule is correct and worth supporting, and which corpus signals are not evidence at all.
tags: [policy, validation, testing, corpus]
status: draft
timestamp: 2026-09-19T00:00:00Z
generated: { by: claude/opus-5, at: 2026-09-19T00:00:00Z }
sources:
  - id: rfc7322
    resource: https://www.rfc-editor.org/rfc/rfc7322.html
    title: "RFC 7322: RFC Style Guide"
  - id: styleguide
    resource: https://www.rfc-editor.org/styleguide/
    title: Updates to the RFC Style Guide (living document)
  - id: rpc-terms
    resource: https://rpc-wiki.rfc-editor.org/doku.php?id=terms
    title: RFC-Specific Terms list
  - id: authors-required
    resource: https://authors.ietf.org/required-content
    title: "Internet-Draft Author Resources: Required content"
  - id: rfc-index
    resource: https://www.rfc-editor.org/rfc-index.xml
    title: RFC index, which maps each RFC to the draft version it was published from
---

This governs how a rule in either style is shown to be correct. It applies on top
of the [evidence policy](/evidence-policy.md), which governs where a rule's
authority comes from.

# The rule that governs everything here

**Rules come from policies, procedures, and official tools. A rule MUST NOT be
derived from, added because of, or removed because of the corpus.**

The corpus exists to demonstrate that a rule works -- that it fires on text a
human reviewer would accept as a genuine finding, and stays quiet elsewhere. It
cannot tell us whether the style guide contains the right rules, because it is a
sample of what people wrote, not a statement of what the guidance says. A rule's
source is the [evidence policy](/evidence-policy.md) chain: an RFC, an RFC Editor
tier, an IESG statement, or a check that official tooling implements. Corpus
behavior may send us back to read a source more carefully. It is never itself the
source.

# What the styles are accountable to

**A rule is accountable to documented guidance, not to observed editor
behavior.** The [draft style](/draft-style.md) exists to catch deviations from
what the RFC Editor, the IESG, and authors.ietf.org have written down. Whether
any particular deviation was caught, argued about, or fixed in any particular
document is a separate question, and not one this project answers.

This distinction decides what counts as evidence.

# Why document history is not authority

An Internet-Draft passes through many hands before publication: the authors
revising it, working group participants, directorate and early reviewers, the
responsible Area Director, and the IESG. The RFC Production Center edits the text
only at the end, during AUTH48. Review comments along the way are frequently
subjective, are weighed rather than obeyed, and authors and editors may decline
them.

Two inferences from a `-00`-to-RFC diff are therefore invalid, and neither may
appear in a rule PR:

1. **"The pattern disappeared before publication, so the rule is right."** The
   change may have come from any actor, for any reason, including reasons
   unrelated to style.
2. **"The pattern survived into a published RFC, so the rule is a false
   positive."** This is the more damaging error. A documented requirement that
   nobody insisted on is still a documented requirement. Publication is not
   ratification of every sentence in a document.

A corollary: **a rule that finds nothing in a corpus has not been disproven.** A
low hit rate says the sampled documents were clean, which affects how urgent a
rule is, never whether it is correct.

# The three evidence classes

Every rule PR shows all three. Only the first establishes correctness; the other
two establish that the rule is implemented well and is supportable.

## 1. Source fidelity -- establishes correctness

The fixture pins what the rule flags and what it leaves alone, and the `link`
plus quoted sentence establish that the flagged text is what the source actually
proscribes. This is the repository's existing contract: a rule with a fixture,
a recorded expectation, and a verbatim quote is a rule whose correctness a
reviewer can check by reading, without any corpus.

A near-miss in the fixture is worth more than a thousand corpus lines, because it
pins the boundary of the claim deliberately rather than incidentally.

## 2. Mechanical correctness -- catches rule defects

A rule can be faithful to its source and still be broken as a pattern. This is
checked by inspecting what the rule matches in real text and classifying each
alert against the cited source:

| Classification | Meaning | Action |
|---|---|---|
| **Deviation** | The text departs from the cited guidance | Rule working as specified, whether or not anyone would have fixed it |
| **Rule defect** | The match is not what the source describes at all | Fix the rule before merge |
| **Legitimate exception** | The source itself permits this case | Add an exception, and a fixture near-miss for it |

The classification is made by reading the source, not by checking what an editor
did. Running the style over this repository's own prose caught two rule defects
before v0.1 shipped, both of which a corpus count would have reported as
"working":

- `AbbreviationVerbs` matched `RPC's` and `OKF's`. The Style Guide sentence is
  about verb suffixes, `XOR'ed` and `NAT-ed`; a possessive is correct English.
  The token alternation dropped `s`, and the possessive is now a fixture
  near-miss.
- `CitationSpacing` matched `[RFC 7322](https://...)`, which is a Markdown link
  whose text happens to name an RFC, not a citation tag. Distinguishing the two
  needs a negative lookahead that RE2 does not have, so the rule became a
  `script` rule that filters candidates in Tengo. At `error` level this was the
  most expensive defect in the set.

Two further findings from the first corpus run illustrate the same point:

- A dotted-quad pattern for `IETF-Draft.ExampleIPv4` matched 53 strings in
  published RFCs, every one of them a **section number** such as `7.2.1.1`. That
  is a rule defect: a section number is not an IPv4 address under any reading of
  RFC 5737, and the rule must require address-shaped context before it ships.
- A naive substitution of "Acknowledgements" to "Acknowledgments" is a **rule
  defect** too, but the proof is in the Terms list, which says to leave either
  spelling alone when the document is internally consistent[^rpc-terms] -- not in
  the fact that the spelling appears in published RFCs.

## 3. Human adjudication and noise budget -- establishes whether it is supportable

Two numbers, measured over the pinned corpus and reported in the PR.

**Accepted or dismissed.** A sample of the rule's alerts is shown to someone
reading as an IETF reviewer would, who marks each one as a finding they would
raise or one they would wave through. This is the closest thing to proof that the
tool works for its purpose: an alert a reviewer would dismiss is noise, whatever
its provenance. A high dismissal rate sends the rule back to be narrowed -- the
rule stays, its pattern changes.

**Alerts per 1000 lines**, per rule. A correct rule that fires on every third
line will be disabled by its users no matter how well grounded it is.

Neither number can add a rule, and neither can delete one. A high rate is a
reason to narrow the scope, lower the level, or ship the rule disabled by
default. A zero rate is not a reason to do anything at all.

# The corpus

Built from the RFC index's draft-to-RFC mapping[^rfc-index], which names the
exact draft version each RFC was published from, so documents can be sampled at
any stage. The corpus list is pinned in the repository; the documents are not
committed.

It samples three populations, for three different purposes:

| Population | Purpose | What it cannot show |
|---|---|---|
| Active Internet-Drafts, early versions | The target user's text. Accept/dismiss adjudication, noise budget, and the most realistic source of rule defects | Nothing about authority, and nothing about which rules should exist |
| Published RFCs | Text that survived the full process. Useful for finding rule defects, since a match here deserves scrutiny | Whether a surviving pattern is permitted -- read the source to decide |
| Final draft version against its published RFC | The narrow window in which the RFC Production Center actually edits | What the rules require; only what was enforced in those cases |

The third population is the only one where a change is attributable to the RFC
Editor, and even there it is **corroboration, never authority**. A rule may cite
it as "the RPC made this change in N of M sampled documents". A rule may not cite
it as the reason the guidance exists, and a rule is never rejected because the
RPC left a pattern alone.

# Reporting

`make corpus-report` fetches the pinned corpus, runs each style over it, and
reports per rule: alerts per 1000 lines for each population, and a sample of
alerts for classification and accept/dismiss adjudication. The sample is what a
reviewer reads; the rate is what they budget against.

The report never proposes a rule. A gap it reveals is a prompt to go read the
guidance and the official tools, and any rule that results cites those, not the
report.

The report is regenerated per PR for the rules that PR touches. It is not a gate:
no threshold automatically passes or fails a rule, because the judgment that
matters -- deviation, defect, or exception -- is made by reading the source.

[^rpc-terms]: [RFC-Specific Terms list](https://rpc-wiki.rfc-editor.org/doku.php?id=terms), which states for Acknowledgments: "If consistent (with or without the 'e'), leave it. If inconsistent, delete the 'e'."
