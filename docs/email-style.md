---
type: Style Specification
title: Email style
description: Normative specification for the IETF-Email Vale style, which lints mailing list and issue-tracker prose.
resource: https://github.com/OR13/ietf-vale
tags: [vale, mailing-lists, conduct, style]
status: draft
timestamp: 2026-09-19T00:00:00Z
generated: { by: claude/opus-5, at: 2026-09-19T00:00:00Z }
sources:
  - id: bcp54
    resource: https://www.rfc-editor.org/rfc/rfc7154.html
    title: "RFC 7154 (BCP 54): IETF Guidelines for Conduct"
    last_modified: 2014-03-01T00:00:00Z
  - id: bcp45
    resource: https://www.rfc-editor.org/rfc/rfc9245.html
    title: "RFC 9245 (BCP 45): IETF Discussion List Charter"
  - id: bcp245
    resource: https://www.rfc-editor.org/rfc/rfc9945.html
    title: "RFC 9945 (BCP 245): IETF Community Moderation"
  - id: bcp25
    resource: https://www.rfc-editor.org/rfc/rfc8716.html
    title: "RFC 8716 (BCP 25): Update to the IETF Anti-Harassment Procedures"
---

This specifies the `IETF-Email` Vale style: prose rules for mailing list posts,
issue comments, and review notes. Severity levels follow the
[evidence policy](/evidence-policy.md). The companion
[draft style](/draft-style.md) covers Internet-Drafts and shares no rules with
this one.

# What this style is, and is not

The documents this style draws on govern **conduct**, not vocabulary. RFC 9945
states plainly that moderation "consists of subjective judgment calls" and that
its list of disruptive behavior is "non-exhaustive"[^bcp245]. No linter can
decide whether a message is uncivil, off-topic, or made in good faith.

What a linter can do is catch a small number of phrasings, before the message is
sent, that reliably read worse than their author intended.

Therefore:

- **Every rule in this style is `suggestion`.** No rule in this style will ever
  be `error` or `warning`.
- **An alert is not a finding.** It is a prompt to reread one sentence. It is not
  evidence that a message violated BCP 54, BCP 45, or BCP 245, and it must never
  be cited in a moderation discussion, a complaint, or a response to one.
- **The style lints your own drafts, not other people's mail.** Running it over
  an archive to characterize another participant is a misuse of this repository,
  and the rule messages say so.
- **Silence means nothing.** A message with no alerts has not been judged
  civil, on-topic, or constructive.

# Scope

**In scope**: prose the author is about to send — mailing list posts, issue and
pull request comments, review notes — in plain text, `.eml`, and Markdown. Vale
lints all three directly, with no transform.

**Out of scope**: anything requiring thread context, history, or intent. These
are tracked as permanently uncovered in `coverage/IETF-Email/out-of-scope.yml`
so that they are visibly considered rather than silently missing:

| Concern | Why a linter cannot do it |
|---|---|
| Venue and topicality (RFC 9245 §2)[^bcp45] | Requires knowing the list's charter and what else is under discussion |
| Advertising and unsolicited bulk email (RFC 9245 §2)[^bcp45] | Indistinguishable from a legitimate announcement without context |
| Sealioning (RFC 9945 App. B)[^bcp245] | Defined by persistence across a thread, not by any single message |
| Arguing counter to an approved charter (RFC 9945 App. B)[^bcp245] | Requires the charter and the argument's history |
| Harassment (BCP 25)[^bcp25] | A matter for the Ombudsteam, not a text pattern |

# Conformance

There is no failing state. Because every rule is `suggestion`, a message
"conforms" whenever its author has read the prompts and decided. A reviewer
adopting this style should set `MinAlertLevel = suggestion` and treat the output
as a pre-send checklist.

# Normative basis

| Document | Standing | What it supports |
|---|---|---|
| RFC 7154[^bcp54] | **BCP 54**, current | "IETF participants extend respect and courtesy to their colleagues at all times"; "We dispute ideas by using reasoned argument rather than through intimidation or personal attack"; "limiting the use of slang" |
| RFC 9245[^bcp45] | **BCP 45**, current, obsoletes RFC 3005 | Inappropriate postings include "Uncivil commentary, regardless of the general subject" |
| RFC 9945[^bcp245] | **BCP 245**, current, obsoletes RFC 3934 and RFC 3683 | Non-normative examples of disruptive behavior, and the statement that moderation is subjective |
| RFC 8716[^bcp25] | **BCP 25**, current, updates RFC 7776 | Establishes that harassment is handled by people, not tooling — cited here for the boundary it sets |

Cite RFC 9245 and RFC 9945, never RFC 3005 or RFC 3934; both were obsoleted, and
RFC 3005's "Unprofessional commentary" became "Uncivil commentary" in the
current text.

# Rule inventory

Each rule is one PR, subject to the evidence policy. All are `suggestion`.

| Rule | Prompt | Basis |
|---|---|---|
| `IETF-Email.PersonalAttacks` | This sentence judges a person rather than an argument | RFC 7154 §2: "We dispute ideas by using reasoned argument rather than through intimidation or personal attack"[^bcp54] |
| `IETF-Email.Incivility` | This wording reads as contempt | RFC 9245 §2: "Uncivil commentary, regardless of the general subject"[^bcp45] |
| `IETF-Email.Slang` | This term may not be understood by participants who do not have English as a first language | RFC 7154 §2: "All participants, particularly those with English as a first language, attempt to accommodate the needs of other participants by communicating clearly, including speaking slowly and limiting the use of slang"[^bcp54] |
| `IETF-Email.Idioms` | This idiom does not survive translation | RFC 7154 §2, same sentence[^bcp54] |
| `IETF-Email.UnsupportedDismissals` | This dismisses a position without offering data | RFC 7154 §2: "Try to provide data and facts for your standpoints so the rest of the participants who are sitting on the sidelines watching the discussion can form an opinion"[^bcp54] |

`IETF-Email.Slang` is the most distinctly IETF rule in this repository and the
least contentious: BCP 54 names slang explicitly, and the guidance is addressed
to native English speakers writing for an international audience.

`IETF-Email.PersonalAttacks` and `IETF-Email.Incivility` carry the highest false
positive risk in the repository, because the same words are civil or uncivil
depending on who they are about. Both keep deliberately short token lists, and
both fail the evidence policy's overreach check if their message asserts that the
text *is* uncivil rather than asking the author to look again.

# Rules deliberately not proposed

- **Profanity lists beyond contempt markers.** Swearing about a protocol is not
  incivility toward a person, and a generic profanity filter would flag the
  former far more often than the latter.
- **Tone scoring or sentiment analysis.** Not reproducible, not citable, and not
  something a reviewer can argue with.
- **Anything keyed to who wrote the message.** This style has no notion of
  authorship and will not acquire one.

[^bcp54]: [RFC 7154](https://www.rfc-editor.org/rfc/rfc7154.html), BCP 54, IETF Guidelines for Conduct.
[^bcp45]: [RFC 9245](https://www.rfc-editor.org/rfc/rfc9245.html), BCP 45, IETF Discussion List Charter.
[^bcp245]: [RFC 9945](https://www.rfc-editor.org/rfc/rfc9945.html), BCP 245, IETF Community Moderation.
[^bcp25]: [RFC 8716](https://www.rfc-editor.org/rfc/rfc8716.html), BCP 25, Update to the IETF Anti-Harassment Procedures.
