# IETF styles for Vale

> **NOTE**: This project is neither maintained nor endorsed by the IETF or the RFC Editor.

This repository contains [Vale](https://vale.sh)-compatible implementations of
published IETF editorial guidance. The specifications in [`docs/`](docs/)
describe each style's scope, sources, rules, and limitations.

## IETF-Draft

The draft style checks Internet-Draft and RFC-bound prose. Its rules cover
terminology, abbreviations, citations, BCP 14 keywords, example values, URIs,
and several writing conventions. For example, one review can catch protocol
capitalization, a preference incorrectly expressed as a BCP 14 keyword, and an
HTTP URI that should be secure:

```diff
- The Http client SHOULD use http://example.com to select an algorithm.
+ The HTTP client should use <https://example.com> to select an algorithm.
```

It also checks that example addresses use documentation ranges and that
unfamiliar abbreviations are expanded:

```diff
- The XYZ sends a request from 10.0.0.1.
+ The Example Extension (XYZ) sends a request from 192.0.2.1.
```

Read the [draft-style specification](docs/draft-style.md) for the complete
rule set, sources, and boundaries.

## IETF-Email

The email style checks mailing-list and issue-tracker prose. It offers
suggestions for slang, idioms, and unfamiliar abbreviations that may be
difficult for international participants to interpret:

```diff
- The XYZ WG is gonna punt on this; let's boil the ocean.
+ The Example Working Group (XYZ) is going to defer this; let's solve the problem incrementally.
```

Read the [email-style specification](docs/email-style.md) for its deliberately
narrow scope and the guidance it can enforce.

## IETF-Draft-Optional

[`IETF-Draft-Optional`](IETF-Draft-Optional/) contains editorial prompts that
are useful during drafting but are not default IETF requirements. It currently
checks repeated words and a conservative set of common `a`/`an` errors:

```diff
- The the client sends sends a request.
+ The client sends a request.

- Use a abstract in the draft.
+ Use an abstract in the draft.
```

These rules run at suggestion level and are packaged separately, so existing
users do not receive them unless they opt in. Their rationale is documented in
the [draft-style specification](docs/draft-style.md#group-0b-tooling-conventions-with-no-documented-requirement).

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

To opt into the optional draft checks, add their package alongside the default
draft package and list the style in `BasedOnStyles`:

```ini
Packages = https://github.com/OR13/ietf-vale/releases/latest/download/IETF-Draft.zip, \
           https://github.com/OR13/ietf-vale/releases/latest/download/IETF-Draft-Optional.zip

[*.md]
BasedOnStyles = Vale, IETF-Draft, IETF-Draft-Optional
```

Because it is a separate package, the optional rules never appear when a user
installs only `IETF-Draft`.

For repeatable builds, replace the latest package URL with a tagged URL such as
`https://github.com/OR13/ietf-vale/releases/download/v0.4.6/IETF-Draft.zip`.

## Repository Structure

- [`IETF-Draft/`](IETF-Draft/), [`IETF-Email/`](IETF-Email/), and
  [`IETF-Draft-Optional/`](IETF-Draft-Optional/) contain the YAML rules and
  each style's `meta.json`. These directories are packaged for release.
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
make check-draft FILE=draft.txt      # run Vale and idnits over one draft
make package                         # build all style ZIP archives
```

`make lint` checks the YAML rules and runs Vale over the repository's
documentation. `make test` runs the rule fixtures and requires Vale on your
`PATH`. `make check-draft` expects the idnits v3 command to be installed; run
`npm install -g @ietf-tools/idnits` first. It is intended for plaintext
Internet-Drafts, which is the input format idnits validates. Build Markdown or
RFCXML sources to the format required by your document toolchain before running
the combined check.

## License

The rules and tooling in this repository are released under the [MIT License](LICENSE).
RFCs and the RFC Style Guide are published by the IETF and the RFC Editor under the
[IETF Trust Legal Provisions](https://trustee.ietf.org/documents/trust-legal-provisions/);
quoted guidance remains subject to those terms.
