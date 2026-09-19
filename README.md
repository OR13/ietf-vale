# IETF styles for Vale

> **NOTE**: This project is neither maintained nor endorsed by the IETF or the RFC Editor.

This repository contains [Vale-compatible](https://vale.sh) implementations of
published IETF editorial guidance. It ships **two styles**, because drafts and
mailing list posts are governed by different documents and share no rules:

| Style | Lints | Governed by |
|---|---|---|
| `IETF-Draft` | Internet-Drafts and RFC-bound prose | [RFC 7322](https://www.rfc-editor.org/rfc/rfc7322.html), the [RFC Editor Style Guide](https://www.rfc-editor.org/styleguide/), [authors.ietf.org](https://authors.ietf.org/), BCP 14 |
| `IETF-Email` | Mailing list and issue-tracker prose | BCP 54, BCP 45, BCP 245 |

Each style is specified before it is built. Read the specifications in
[`docs/`](docs/index.md) — [draft style](docs/draft-style.md),
[email style](docs/email-style.md), and the shared
[evidence policy](docs/evidence-policy.md) that decides how a rule earns its
severity level.

**Neither style ships any rules yet.** What exists is the scaffolding — the
specifications, the tests, and the CI — that every rule must pass through. See
[CONTRIBUTING.md](CONTRIBUTING.md).

## Getting Started

Add the package you want to your configuration file and run `vale sync`:

```ini
StylesPath = styles
MinAlertLevel = suggestion

Packages = https://github.com/OR13/ietf-vale/releases/latest/download/IETF-Draft.zip

[*.md]
BasedOnStyles = Vale, IETF-Draft
```

To lint mail alongside drafts, install both packages and scope them separately.
Vale has no rule-name wildcards, and a section's `BasedOnStyles` replaces rather
than extends earlier ones, so name every style you want in each section:

```ini
Packages = https://github.com/OR13/ietf-vale/releases/latest/download/IETF-Draft.zip, \
           https://github.com/OR13/ietf-vale/releases/latest/download/IETF-Email.zip

[*.md]
BasedOnStyles = IETF-Draft

[posts/*.{txt,eml}]
BasedOnStyles = IETF-Email
```

See [Packages](https://vale.sh/docs/keys/packages) for more information.

## Repository Structure

<dl>
  <dt><code>/IETF-Draft</code>, <code>/IETF-Email</code></dt>
  <dd>The <a href="https://yaml.org/">YAML</a>-based rule implementations, plus each style's <code>meta.json</code>. These directories are what get packaged and released.</dd>

  <dt><code>/docs</code></dt>
  <dd>The specifications, as an <a href="https://github.com/GoogleCloudPlatform/knowledge-catalog/blob/main/okf/SPEC.md">OKF</a> bundle. A rule is specified here before it is implemented. The bundle deliberately omits OKF's optional <code>log.md</code>: git already records what changed and when.</dd>

  <dt><code>/fixtures</code></dt>
  <dd>The individual unit tests, one directory per rule at <code>fixtures/&lt;Style&gt;/&lt;Rule&gt;/</code>, each with a <code>.vale.ini</code> that enables only that rule.</dd>

  <dt><code>/testdata</code></dt>
  <dd>The expected Vale output for each fixture, one <code>&lt;Style&gt;.&lt;Rule&gt;.ct</code> file per fixture. We use <a href="https://github.com/google/go-cmdtest">go-cmdtest</a> to run Vale against each fixture and compare its output. Run the suite with <code>go test ./...</code>; regenerate expectations after an intentional change with <code>go test ./... -update</code>.</dd>

  <dt><code>/coverage</code></dt>
  <dd>How much of the source guidance each style implements, tracked topic by topic under <code>coverage/&lt;Style&gt;/</code>. Each key is a subtopic set to <code>true</code> or <code>false</code>, optionally followed by a comment naming the rules that implement it. Run <code>go test -v -run TestCoverage ./...</code> to print the figures; the same test fails if a named rule no longer exists, so a topic can't silently claim coverage it has lost.</dd>
</dl>

## Local Development

Requires [Vale](https://vale.sh/docs/install), Go 1.21+, and Node (for the
specification linter).

```bash
make new RULE=Ellipses               # scaffold a rule in IETF-Draft
make new RULE=Slang STYLE=IETF-Email # ... or in IETF-Email
make test                            # fixture snapshots, rule contracts, coverage
make lint                            # yamllint over the styles, vale over our prose
make lint-docs                       # okf-lint over the specification bundle
make update                          # regenerate testdata/*.ct from current behavior
make package                         # build IETF-Draft.zip and IETF-Email.zip
```

## License

The rules and tooling in this repository are released under the [MIT License](LICENSE).
RFCs and the RFC Style Guide are published by the IETF and the RFC Editor under the
[IETF Trust Legal Provisions](https://trustee.ietf.org/documents/trust-legal-provisions/);
quoted guidance remains subject to those terms.
