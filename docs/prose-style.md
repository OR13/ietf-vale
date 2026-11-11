---
type: Style Specification
title: Optional prose style
description: Research and experimental editorial checks for formulaic prose.
resource: https://github.com/OR13/ietf-vale
tags: [vale, prose, research]
status: draft
timestamp: 2026-11-11T00:00:00Z
sources:
  - id: graphite
    resource: https://graphite.io/research/ai-tells
    title: AI Tells (Graphite, 2026)
  - id: claudisms
    resource: https://claudisms.ai/
    title: Claudisms (living editorial catalog)
  - id: processing
    resource: https://arxiv.org/abs/2610.29162v1
    title: Why is AI-generated language so hard to understand?
  - id: farrell-feng
    resource: https://www.ietf.org/archive/id/draft-fengfar-led-01.html#section-2
    title: Dealing with LLMs in IETF Discussions (individual Internet-Draft, revision 01)
  - id: openai-style
    resource: https://developers.openai.com/api/docs/guides/latest-model?model=gpt-6-astra
    title: OpenAI GPT-6 writing guidance (living documentation)
  - id: anthropic-density
    resource: https://platform.claude.com/docs/en/build-with-claude/prompt-engineering/prompting-claude-fable-5-1#writing-density
    title: Anthropic Fable 5.1 writing density guidance
  - id: anthropic-reports
    resource: https://platform.claude.com/docs/en/build-with-claude/prompt-engineering/prompting-claude-fable-5-2#make-reports-easy-to-follow
    title: Anthropic Fable 5.2 report and style guidance
  - id: anthropic-prompts
    resource: https://platform.claude.com/docs/en/build-with-claude/prompt-engineering/claude-prompting-best-practices
    title: Anthropic prompting best practices (living documentation)
---

# Purpose and standing

`Prose-Optional` offers experimental editorial suggestions for technical prose,
including prose drafted with an assistant. It is separately installed and
explicitly enabled. These checks have no IETF standing and do not measure IETF
style coverage. Their messages ask for review; they never identify an author or
model, declare a claim false, or automatically rewrite technical meaning.

Research reviewed on 2026-11-11 supports testing small, explainable patterns.
There is no defensible ranking of the "most frustrating" defects across all
writers or tasks. The priorities below reflect relevance to this repository
and the cost of a mistaken alert.

# Evidence and priorities

[Graphite's original study][graphite] compares 10,000 older human web articles
with 90,000 generated articles. It reports model-dependent contrast formulas,
importance framing, and assurances about avoiding tradeoffs. Its web-article
baseline and shared generation prompt limit transfer to standards prose;
frequency is not evidence that an individual sentence is defective. Some theme
labels and prose ratings are model-generated. This supports review prompts,
not bans on ordinary words or punctuation.

The [Claudisms catalog][claudisms] provides first-hand editorial complaints
about staged reflection and instructions to the reader to feel impressed.
It is a changing personal preference list, not a controlled study or an IETF
source. We use a small phrase subset, not its blanket prohibitions on ordinary
verbs, metaphors, or emoji.

An [October 2026 preprint][processing] measures processing-related features
across nine models, including unusual word combinations, noun stacking, and
long dependencies. This supports investigating readability beyond buzzwords.
It does not validate a simple regex for grammatical complexity or justify
banning technical compound nouns.

| Candidate defect | Decision | Reason |
|---|---|---|
| Rehearsed contrasts | Implement a bounded review prompt | A small set of explicit sentence shapes is testable. |
| Announcing importance | Implement a phrase-level prompt | Ask for the actual consequence. |
| Reader-directed filler | Implement a narrow phrase list | Removing the setup often exposes the observation. |
| Claims of no tradeoff | Ask for conditions or evidence | Matching cannot determine whether evidence is already present. |
| Dense noun stacks and invented compounds | Defer | Protocol names and established technical terms would trigger broad patterns. |
| Excessive praise and emphatic agreement | Implement a narrow phrase list | Prompt a direct answer or explanation; sincere praise remains a legitimate exception. |
| Geometric framing | Implement two bounded phrase forms | Ask whether the property is concrete; do not ban geometric vocabulary. |
| Counted labels around a colon | Implement a bounded sentence pattern | Ask for the relationship; useful compact summaries can still match. |
| Excessive apologies, misplaced certainty, and artificial balancing | Defer | Context and the author's intended position determine whether these are defects. |
| Stock summary lead-ins | Implement three sentence-initial forms | Prompt a direct statement; summaries themselves remain useful. |
| Decorative metaphors | Implement two phrase families | Prefer literal meaning where appropriate; do not ban metaphors generally. |
| Repetitive conclusions and excessive lists | Defer | Document structure and audience matter more than a fixed count. |
| Fabricated citations or unsupported conclusions | Human or external verification | Vale cannot verify factual support. |
| Em dashes, "delve", "robust", "leverage", "load-bearing" | No blanket ban | Authorship guesses and word replacement do not establish clarity. |

The existing `IETF-Draft.Concision` already covers several verbose wind-ups.
The new style avoids duplicating that list.

[Farrell and Feng's individual draft][farrell-feng], revised August 6, 2026,
records excessive praise, geometric framing, and counted-label constructions
in Section 2. These examples already appeared in revision 00. Revision 01 adds
community discussion feedback. The draft concerns discussions rather than RFC
text and is work in progress with no IETF consensus standing. Its examples
motivate three experimental checks here; the additional praise variants and
regex boundaries are our choices. None establishes that a writer used AI.

# Lessons from vendor guidance

Official documentation checked on 2026-11-11 provides additional motivation.
These are living pages; the review date does not assert a publication date.
Vendor prompting advice is neither a writing standard nor evidence that a
particular expression identifies an AI author.

[OpenAI's writing guidance][openai-style] favors direct explanations and calls
out stock summary introductions and unnecessary contrast. `SummaryLabels`
checks three explicit lead-ins. `ContrastFraming` already covers a subset of
the contrast advice. We retain legitimate technical negations and compounds.

[Anthropic's density guidance][anthropic-density] recommends literal expression
when a metaphor obscures meaning. `ManneredMetaphors` checks its dial and
argument-value examples with a few documented variants. It does not attempt
to classify all figurative language. The same page warns that older rules
against formatting can remove useful structure in a newer model.

[Anthropic's newer report guidance][anthropic-reports] asks writers to supply
the context readers lack and avoid invented terminology. It also warns that
rigid sentence-length limits can make prose choppy. Consequently, we do not
add hard word-count, bullet-count, or hyphen bans. Existing abbreviation checks
help with a small part of the context problem; they cannot establish that a
report is understandable on its own.

[Anthropic's general guidance][anthropic-prompts] favors actionable instructions
and relevant, varied examples. Our application is to keep diagnostics focused
on the revision to consider and to test both intended matches and exceptions.
This is a design choice, not a vendor validation of these Vale rules.

For assisted editing, pair the linter with this project-specific instruction:

```text
Treat Vale suggestions as requests for review. For each alert, identify the
claim and its intended reader before revising. Prefer a direct explanation
when it preserves the meaning. Keep required terms, uncertainty, evidence,
and useful comparisons. Keep lists and headings when they aid comprehension.
Do not add claims or remove qualifications merely to silence a suggestion.
If the original is clearer, retain it and explain the exception to the editor.
```

Judge revisions against the original meaning as well as alert counts. Fewer
alerts alone do not show better prose. Revisit these choices with new models
and real documents instead of accumulating unconditional prohibitions.

# Rule contract

All nine rules use `existence`, case-insensitive matching, and `suggestion`
severity. `CountedLabels` and `SummaryLabels` use `scope: sentence & ~heading`;
the other rules use
`scope: ~code`. Both exclude Markdown code in the regression fixtures. No automatic replacement is provided. Each YAML `link`
points to the research or editorial catalog motivating the check; the exact
matching boundary is a project choice documented here and in fixtures.

| Rule | Boundary | Legitimate exception |
|---|---|---|
| `ContrastFraming` | Paired "this/it/that is not about ...; this/it/that is about" forms, "is not just/merely/simply ... but", and "less like a/an ... and more like"; straight or curly contractions; variable spans limited to 100 characters and exclude sentence punctuation and newlines | A contrast may correct a misconception or distinguish two architectures. |
| `ImportanceClaims` | "this matters because", "the key/main takeaway is", "the only thing that matters" | A summary may be appropriate for the audience. |
| `ReaderPrompts` | "let that sink in", "sit with that", "here is/here's the thing" followed by punctuation, "let's unpack this/that" | Quoted dialogue or a literal instruction may match. |
| `ExcessivePraise` | "you have/you've asked exactly the right question", "you are/you're absolutely/completely/entirely right", and "that is/that's/this is [a/an] excellent/great/fantastic/brilliant question"; straight or curly apostrophes | Sincere praise or emphatic agreement may be appropriate. |
| `GeometricFraming` | "the seam is" or "there are [count] dimensions", where count is 1-99 in digits or one through twenty in words | A physical seam or literal matrix dimensions can match. |
| `CountedLabels` | An entire prose sentence containing count plus one or two ASCII words, a colon, then count plus one or two ASCII words; the same count range as `GeometricFraming`; optional final punctuation; horizontal whitespace only | A useful summary of interface inputs and outputs can match. |
| `SummaryLabels` | Sentence-initial "bottom line:" or "in short:" (also comma), and "the simplest mental model is"; headings excluded | A summary introduction can be useful in a long document. |
| `ManneredMetaphors` | "a/the/this/that dial worth turning" and "point/argument/analogy/metaphor earns its keep" | A real control dial, or an intentional discussion of rhetoric, can match. |
| `TradeoffClaims` | "without sacrificing/compromising" followed by security, performance, reliability, scalability, flexibility, simplicity, correctness, or privacy | A nearby measurement may already substantiate the claim. |

Counts and labels are lexical patterns, not grammatical analysis: the counted
words need not actually be nouns. Counted headings are outside the sentence
scope. These limits keep the check small; it does not flag arbitrary colons.

Patterns intentionally miss other phrasings. Fixed phrase whitespace accepts
line wrapping. The variable spans in contrasts do not cross a plaintext line
break; Markdown preprocessing can join wrapped lines. This precision tradeoff
is preferable to matching across unrelated sentences or paragraphs.

Markdown inline and fenced code are excluded. Plaintext has no equivalent
code markup. Prose quotations can still trigger suggestions. Disable an
individual rule when its exceptions dominate a document; an alert is a request
to inspect, not an instruction to delete.

# Validation

Fixtures cover positive matches in Markdown and plaintext, curly and straight
apostrophes, a wrapped phrase, technical near-misses, and Markdown code
exclusions. The recorded `.ct` files pin every emitted alert. Keep intentionally
matched legitimate exceptions in the fixtures as well as examples that stay
quiet: a regex cannot distinguish a supported tradeoff claim from an unsupported
one with the same words.

A small negative-corpus smoke test on 2026-11-11 used all nine rules with
Vale 3.21.0 and the RFC
Editor's full plaintext files, including front matter and appendices:

| File | Lines | Alerts | Alerts per 1000 lines |
|---|---:|---:|---:|
| RFC 8446 | 8963 | 0 | 0 |
| RFC 9000 | 8485 | 0 | 0 |
| RFC 9110 | 10785 | 0 | 0 |

There were no alerts to classify. This is evidence of low noise on these three
files only; it establishes neither recall nor a general false-positive rate.
Synthetic fixtures demonstrate matching behavior, not editorial effectiveness.
Before expanding the patterns, review alerts on contemporary drafts and human
mail as well as assistant-generated samples.

To repeat the smoke test, download `rfc8446.txt`, `rfc9000.txt`, and `rfc9110.txt`
from `https://www.rfc-editor.org/rfc/`. Use a temporary configuration with an
absolute `StylesPath` pointing at this checkout, `MinAlertLevel = suggestion`,
and `BasedOnStyles = Prose-Optional` under `[*.txt]`. Run:

```sh
vale --config=/path/to/corpus/.vale.ini --no-global --no-exit --output=JSON /path/to/corpus/
```

[graphite]: <https://graphite.io/research/ai-tells>
[claudisms]: <https://claudisms.ai/>
[processing]: <https://arxiv.org/abs/2610.29162v1>

[farrell-feng]: <https://www.ietf.org/archive/id/draft-fengfar-led-01.html#section-2>

[openai-style]: <https://developers.openai.com/api/docs/guides/latest-model?model=gpt-6-astra#personality-and-writing-style>
[anthropic-density]: <https://platform.claude.com/docs/en/build-with-claude/prompt-engineering/prompting-claude-fable-5-1#writing-density>
[anthropic-reports]: <https://platform.claude.com/docs/en/build-with-claude/prompt-engineering/prompting-claude-fable-5-2#make-reports-easy-to-follow>
[anthropic-prompts]: <https://platform.claude.com/docs/en/build-with-claude/prompt-engineering/claude-prompting-best-practices#control-the-format-of-responses>
