---
okf_version: "0.2"
---

# Overview

Each style is specified before it is implemented. Read the [draft-style
specification](/draft-style.md), the [email-style
specification](/email-style.md), and the shared [evidence
policy](/evidence-policy.md), which explains how each rule earns its severity
level.

The repository ships 24 draft rules and 2 email rules. Every rule traces to
published guidance. The email style is intentionally narrow: testing against
real IETF list mail showed that much of [BCP
54](https://www.rfc-editor.org/info/bcp54) and [BCP
45](https://www.rfc-editor.org/info/bcp45) concerns context rather than words,
so three proposed checks were withdrawn instead of becoming noisy alerts. See
[what the email style can enforce](/email-style.md#what-this-style-can-enforce).

The coverage manifests currently represent 18 of 34 draft topics (52.9%) and
2 of 9 email topics (22.2%). Run `go test -v -run TestCoverage ./...` for the
current report. To propose or implement another rule, see
[`CONTRIBUTING.md`](https://github.com/OR13/ietf-vale/blob/main/CONTRIBUTING.md).

`IETF-Draft` is **not a submission check**. It lints prose; [idnits](https://github.com/ietf-tools/idnits)
owns structure, boilerplate, and references. Run both tools: the [tool
landscape](/tool-landscape.md) explains which checks belong to which tool.

# Style specifications

* [Draft style](/draft-style.md) - Normative specification for the `IETF-Draft` Vale style, which lints Internet-Draft and RFC-bound prose.
* [Email style](/email-style.md) - Normative specification for the `IETF-Email` Vale style, which lints mailing list and issue-tracker prose.

# Reference

* [Tool landscape](/tool-landscape.md) - What idnits, DraftForge, xml2rfc and this style each own, and how an author runs them together.

# Policy

* [Evidence policy](/evidence-policy.md) - How a rule earns its severity level, and what every rule PR must show.
* [Validation methodology](/validation.md) - What evidence shows a rule is correct, and which corpus signals are not evidence at all.
