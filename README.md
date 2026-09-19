# IETF styles for Vale

> **NOTE**: This project is neither maintained nor endorsed by the IETF or the RFC Editor.

This repository contains [Vale](https://vale.sh)-compatible implementations of
published IETF editorial guidance. It provides two independent styles because
Internet-Drafts and IETF discussion posts follow different guidance:

- [`IETF-Draft`](IETF-Draft/) lints Internet-Draft and RFC-bound prose. Its
  sources include [RFC 7322](https://www.rfc-editor.org/rfc/rfc7322.html), the
  [RFC Editor Style Guide](https://www.rfc-editor.org/styleguide/),
  [authors.ietf.org](https://authors.ietf.org/), and
  [BCP 14](https://www.rfc-editor.org/info/bcp14).
- [`IETF-Email`](IETF-Email/) lints mailing-list and issue-tracker prose. Its
  sources include [BCP 54](https://www.rfc-editor.org/info/bcp54),
  [BCP 45](https://www.rfc-editor.org/info/bcp45), and
  [BCP 245](https://www.rfc-editor.org/info/bcp245).

Each style is specified before it is implemented. Start with the
[documentation index](docs/index.md), then read the [draft-style
specification](docs/draft-style.md), the [email-style
specification](docs/email-style.md), and the shared [evidence
policy](docs/evidence-policy.md), which explains how each rule earns its
severity level.

The repository currently ships 24 draft rules and 2 email rules. Every rule
traces to published guidance. The email style is intentionally narrow: testing
against real IETF list mail showed that much of [BCP
54](https://www.rfc-editor.org/info/bcp54) and [BCP
45](https://www.rfc-editor.org/info/bcp45) concerns context rather than words,
so three proposed checks were withdrawn instead of becoming noisy alerts. See
[what the email style can enforce](docs/email-style.md#what-this-style-can-enforce).

The coverage manifests currently represent 14 of 34 draft topics (41.2%) and
2 of 9 email topics (22.2%). Run `go test -v -run TestCoverage ./...` for the
current report. To propose or implement another rule, see
[CONTRIBUTING.md](CONTRIBUTING.md).

`IETF-Draft` is **not a submission check**. It lints prose; [idnits](https://github.com/ietf-tools/idnits)
owns structure, boilerplate, and references. Run both tools: the [tool
landscape](docs/tool-landscape.md) explains which checks belong to which tool.

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
`https://github.com/OR13/ietf-vale/releases/download/v0.4.1/IETF-Draft.zip`.

## Repository Structure

- [`IETF-Draft/`](IETF-Draft/) and [`IETF-Email/`](IETF-Email/) contain the
  YAML rules and each style's `meta.json`. These directories are packaged for
  release.
- [`docs/`](docs/) contains the [OKF](https://github.com/GoogleCloudPlatform/knowledge-catalog/blob/main/okf/SPEC.md)
  specifications. A rule is specified here before it is implemented.
- [`fixtures/`](fixtures/) contains one isolated Vale fixture per rule, with
  examples that should alert and near-misses that should remain clean.
- [`testdata/`](testdata/) contains the expected output for every fixture. The
  tests use [go-cmdtest](https://github.com/google/go-cmdtest); regenerate the
  snapshots with `go test ./... -update` after an intentional rule change.
- [`coverage/`](coverage/) records which topics from each source document are
  implemented. The coverage test also prevents comments from referring to
  rules that no longer exist.

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
