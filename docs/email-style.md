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

**In scope**: prose the author is about to send (mailing list posts, issue and
pull request comments, review notes) in plain text, `.eml`, and Markdown. Vale
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
| RFC 8716[^bcp25] | **BCP 25**, current, updates RFC 7776 | Establishes that harassment is handled by people, not tooling, and is cited here for the boundary it sets |

Cite RFC 9245 and RFC 9945, never RFC 3005 or RFC 3934; both were obsoleted, and
RFC 3005's "Unprofessional commentary" became "Uncivil commentary" in the
current text.

## Sources considered and excluded

**RFC 7704, "An IETF with Much Diversity and Professional Conduct"**, describes
harassment and bullying in terms close to what this style would want: name
calling, belittling or disrespectful comments, and language of condescension.
It is an **Independent Submission, Informational**, so it does not carry IETF
consensus, and this style draws only on consensus documents. It is recorded here
so that its absence is a decision rather than an oversight.

# Rule inventory

**Shipped**: `Slang` and `Idioms`. `PersonalAttacks` and `UnsupportedDismissals`
shipped in v0.1 and were withdrawn; `Incivility` was specified and will not be
implemented. See [what this style can enforce](#what-this-style-can-enforce).

Each rule is one PR, subject to the evidence policy. All are `suggestion`.

| Rule | Prompt | Basis |
|---|---|---|
| `IETF-Email.Incivility` | This wording reads as contempt | RFC 9245 §2: "Uncivil commentary, regardless of the general subject"[^bcp45] |
| `IETF-Email.Slang` | This term may not be understood by participants who do not have English as a first language | RFC 7154 §2: "All participants, particularly those with English as a first language, attempt to accommodate the needs of other participants by communicating clearly, including speaking slowly and limiting the use of slang"[^bcp54] |
| `IETF-Email.Idioms` | This idiom does not survive translation | RFC 7154 §2, same sentence[^bcp54] |

`IETF-Email.Slang` is the most distinctly IETF rule in this repository and the
least contentious: BCP 54 names slang explicitly, and the guidance is addressed
to native English speakers writing for an international audience.

## What this style can enforce

Three of the five rules originally specified here were withdrawn or abandoned
after measurement. They failed for one reason, and the reason generalizes:

> **A rule is enforceable when its source names a property of the words. It is
> not enforceable when its source names a property of the relationship between
> the words and their context.**

Applying that test to BCP 54, BCP 45 and BCP 245:

| Requirement | What the source names | Enforceable |
|---|---|---|
| "limiting the use of slang" | a vocabulary category | **yes**: `Slang` |
| "communicating clearly" for participants without English as a first language | fixed expressions that do not translate | **yes**: `Idioms` |
| "dispute ideas ... rather than through intimidation or personal attack" | whether an argument addresses a person or a position | no |
| "provide data and facts for your standpoints" | whether the message contains support | no |
| "Uncivil commentary, regardless of the general subject" | how words land, given who they are about | no |

The measurements behind the three "no" rows are below. The same test is now part
of the [validation methodology](/validation.md), because it applies to any rule
in either style.

What the moderation community discusses confirms the split. Across 14,265 lines
of `mod-discuss`, the recurring subjects are disruption (188 mentions), threads
(55), tone (23), bad faith (14), and repetition, volume and frequency. Word
choice barely appears: two mentions of insults, one of name calling, one of
swearing, and no mention of slang, idiom or jargon at all. The IETF's moderation
problems are structural, and a line-based linter does not see structure.

This is a narrow style by design. It catches two courtesies an author can check
before sending. It does not catch what moderators deal with, and nothing here
should be mistaken for a conduct tool.

## Why PersonalAttacks was withdrawn

BCP 54 asks that participants "dispute ideas by using reasoned argument rather
than through intimidation or personal attack"[^bcp54], and v0.1 implemented that
as a short list of phrasings. The list was written from the guidance and observed
from nothing, and measurement showed it detects nothing.

The archive was searched for messages in which a participant **themselves**
claimed a personal attack or an ad hominem: 158 such messages on the `tls` and
`mod-discuss` lists, containing 5,491 quoted lines of the material being objected
to. The v0.1 token list matched **once** in all of it.

Reading what the complainants actually said explains why. They describe
attribution, not vocabulary: the clearest formulation in the sample is that
someone "argues about a person's motives rather than the correctness of that
person's position". The most frequent second-person phrases in the objected-to
material are ordinary discussion, such as "you look at the" and "your point is
here". There is no phrase to match, because the same words are an attack or a
question depending on who they are about and what was said before.

A line-based linter cannot make that judgment. The requirement is therefore
recorded in `coverage/IETF-Email/out-of-scope.yml` as something this style does
not reach, alongside sealioning and topicality, rather than being covered by a
rule that produces nothing.

This is not the corpus deciding which rules exist. BCP 54 still states the
requirement, and it is still unmet. What measurement established is narrower and
entirely within the [validation methodology](/validation.md): the implementation
did not match the phenomenon its source describes.

## Why UnsupportedDismissals was withdrawn

BCP 54 asks participants to "provide data and facts for your standpoints so the
rest of the participants ... can form an opinion"[^bcp54]. Being unsupported is a
property of the whole message: a linter would have to judge whether the evidence
offered is sufficient for the claim made.

The archive was searched for threads where participants asked each other for
support. In 102 such messages, 19,014 lines, the v0.1 tokens fired **three**
times, all from one token. The phenomenon itself was plainly present, in the
replies rather than in the messages at fault: "any evidence" appears 69 times,
"unsubstantiated" 19 times, "citation needed" 8 times.

That is the diagnosis. The detectable signal is somebody else objecting, which
arrives in a different message, written by a different author, after the fact.
The rule was looking in the wrong place, and no token list fixes that.

`Incivility` was never implemented, for the same reason. RFC 9245 asks that
commentary not be uncivil, and RFC 9945 says judging that "consists of
subjective judgment calls"[^bcp245]. Both requirements are recorded as
uncovered in `coverage/IETF-Email/out-of-scope.yml`.

# Measured behavior

Measured over 1,428 messages fetched from the IETF's anonymous IMAP archive
(`imap.ietf.org:993`) across eleven lists, including `mod-discuss`, `ietf` and
`tls`, after stripping quoted text and signatures. Reproduce with
`make corpus-mail EMAIL=you@example.org`.

| Rule | Alerts | Per 1000 lines |
|---|---|---|
| `Slang` | 18 | 0.35 |
| `Idioms` | 2 | 0.04 |
| `UnsupportedDismissals` (withdrawn) | 1 | 0.02 |
| `PersonalAttacks` (withdrawn) | **0** | 0.00 |

What this does and does not establish:

- **No noise problem.** Roughly 0.4 alerts per 1000 lines across the whole
  style. Nothing here would make a participant turn the style off.
- **`PersonalAttacks` was withdrawn**, on the evidence described above. It never
  fired on any real message measured, including 158 messages in which a
  participant claimed a personal attack.
- **A token list is a sample, not a closure.** Six of ten `Slang` tokens and one
  of ten `Idioms` tokens ever fired. Silence from either rule means nothing,
  since it can only flag what it lists.
- **Both `Idioms` alerts were `bikeshed`.** It is a term of art here, and it is
  also exactly what BCP 54 describes: an expression a participant without
  English as a first language cannot decode from its parts. It is treated as a
  true positive.
- **The general sample is civil by construction.** It is the most recent traffic
  on technical lists, so a quiet result there says little. The targeted search
  described below was chosen because it contains the phenomenon; everything
  taken from it is de-identified, and it is used to study phrasing, never to
  characterize any participant.
- **Nothing here may change the rule set.** Per the
  [validation methodology](/validation.md), these numbers measure the
  implementation. Adding a token because it appears in the corpus would be
  fitting the style to the sample.

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
