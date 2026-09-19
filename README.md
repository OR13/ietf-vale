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
[`docs/`](docs/index.md): the [draft style](docs/draft-style.md), the
[email style](docs/email-style.md), and the shared
[evidence policy](docs/evidence-policy.md) that decides how a rule earns its
severity level.

**23 rules in `IETF-Draft` and 2 in `IETF-Email`**, each tracing to a published
requirement. The email style is deliberately narrow: measurement against real
IETF list mail showed that most of what BCP 54 and BCP 45 ask for is contextual
rather than lexical, so three rules were withdrawn rather than shipped as
alerts that mean nothing. See
[what this style can enforce](docs/email-style.md#what-this-style-can-enforce).
Coverage of the specified rule set is 18% for the draft style and 22% for the
email style; `go test -v -run TestCoverage ./...` prints
the current figures. See [CONTRIBUTING.md](CONTRIBUTING.md) to add one.

`IETF-Draft` is **not a submission check**. It lints prose; idnits owns
structure, boilerplate and references. Run both; the
[tool landscape](docs/tool-landscape.md) maps who checks what.

## Installation

Install [Vale](https://vale.sh/docs/install) with the package manager for your
platform. For example, macOS users with Homebrew can run:

```bash
brew install vale
```

Other platforms and installation methods are listed in Vale's [installation
guide](https://vale.sh/docs/install). Confirm that Vale is available:

```bash
vale --version
```

Then create a `.vale.ini` file in the root of the project you want to lint.
The following configuration installs the draft style into `styles/` and enables
it for Markdown files:

```ini
StylesPath = styles
MinAlertLevel = suggestion

Packages = https://github.com/OR13/ietf-vale/releases/latest/download/IETF-Draft.zip

[*.md]
BasedOnStyles = Vale, IETF-Draft
```

Install the style package and lint a document:

```bash
vale sync
vale draft.md
```

The first command downloads the package into `styles/`. The second command
prints any matching alerts with their file, line, column, severity, and message.
Replace `draft.md` with a file or directory in your project.

To lint mailing list or issue tracker prose, install both packages and scope
them separately. Vale does not support rule-name wildcards, and each section's
`BasedOnStyles` value lists the styles enabled for that section:

```ini
Packages = https://github.com/OR13/ietf-vale/releases/latest/download/IETF-Draft.zip, \
           https://github.com/OR13/ietf-vale/releases/latest/download/IETF-Email.zip

[*.md]
BasedOnStyles = IETF-Draft

[posts/*.{txt,eml}]
BasedOnStyles = IETF-Email
```

Now `vale posts/` uses `IETF-Email` for matching text files, while Markdown
files use `IETF-Draft`. See Vale's [Packages
documentation](https://vale.sh/docs/keys/packages) for package configuration
options.

### A complete example

With the first configuration above, save this as `draft.md`:

```markdown
# Example

This document is written in order to explain an HTTP service.
The service is available at http://example.com.
```

Run `vale draft.md`. Vale reports the phrases that the draft style recommends
revising, including the verbose phrase and the HTTP URI. The suggestions are
prompts for review; they do not replace idnits or the RFC Editor's review.

For repeatable builds, replace the latest package URL with a tagged URL such as
`https://github.com/OR13/ietf-vale/releases/download/v0.2.0/IETF-Draft.zip`.

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

Requires [Vale](https://vale.sh/docs/install), Go 1.21+, Python with
`yamllint`, and Node.js (for the specification linter). To install the Python
lint dependency, run `python3 -m pip install yamllint`.

```bash
make new RULE=Ellipses               # scaffold a rule in IETF-Draft
make new RULE=Slang STYLE=IETF-Email # ... or in IETF-Email
make test                            # fixture snapshots, rule contracts, coverage
make lint                            # yamllint over the styles, vale over our prose
make lint-docs                       # okf-lint over the specification bundle
make update                          # regenerate testdata/*.ct from current behavior
make package                         # build IETF-Draft.zip and IETF-Email.zip
```

`make lint` checks the YAML rules and runs Vale over the repository's
documentation. `make test` runs the rule fixtures and requires Vale on your
`PATH`.

## License

The rules and tooling in this repository are released under the [MIT License](LICENSE).
RFCs and the RFC Style Guide are published by the IETF and the RFC Editor under the
[IETF Trust Legal Provisions](https://trustee.ietf.org/documents/trust-legal-provisions/);
quoted guidance remains subject to those terms.
