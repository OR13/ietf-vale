# IETF

> **NOTE**: This project is neither maintained nor endorsed by the IETF or the RFC Editor.

This repository contains a [Vale-compatible](https://vale.sh) implementation of IETF
editorial guidance, primarily the [*RFC Style Guide* (RFC 7322)](https://www.rfc-editor.org/rfc/rfc7322.html)
and the [RFC Editor's web portion of the style guide](https://www.rfc-editor.org/styleguide/part2/).

**This repository currently ships no rules.** It is the scaffolding — tests, CI, and
review conventions — that every rule must pass through. See [CONTRIBUTING.md](CONTRIBUTING.md).

## Getting Started

Add the package to your configuration file and run `vale sync`:

```ini
StylesPath = styles
MinAlertLevel = suggestion

Packages = https://github.com/OR13/ietf-vale/releases/latest/download/IETF.zip

[*.md]
BasedOnStyles = Vale, IETF
```

See [Packages](https://vale.sh/docs/keys/packages) for more information.

## Repository Structure

<dl>
  <dt><code>/IETF</code></dt>
  <dd>The <a href="https://yaml.org/">YAML</a>-based rule implementations that make up the style, plus <code>meta.json</code>. This directory is what gets packaged and released.</dd>

  <dt><code>/fixtures</code></dt>
  <dd>The individual unit tests. Each directory is named after a rule in <code>/IETF</code> and includes its own <code>.vale.ini</code> that enables only that rule.</dd>

  <dt><code>/testdata</code></dt>
  <dd>The expected Vale output for each fixture, one <code>&lt;Rule&gt;.ct</code> file per fixture. We use <a href="https://github.com/google/go-cmdtest">go-cmdtest</a> to run Vale against each fixture and compare its output. Run the suite with <code>go test ./...</code>; regenerate expectations after an intentional change with <code>go test ./... -update</code>.</dd>

  <dt><code>/coverage</code></dt>
  <dd>How much of the source guidance we implement, tracked topic by topic. Each key is a subtopic set to <code>true</code> or <code>false</code>, optionally followed by a comment naming the rules that implement it. Run <code>go test -v -run TestCoverage ./...</code> to print the figures; the same test fails if a named rule no longer exists, so a topic can't silently claim coverage it has lost.</dd>
</dl>

## Local Development

Requires [Vale](https://vale.sh/docs/install) and Go 1.21+.

```bash
make new RULE=Ellipses  # scaffold a rule, its fixture, and its test case
make test               # fixture snapshots, rule contract checks, coverage
make lint               # yamllint over IETF/ and coverage/
make update             # regenerate testdata/*.ct from current rule behavior
make package            # build IETF.zip locally
```

## License

The rules and tooling in this repository are released under the [MIT License](LICENSE).
RFCs and the RFC Style Guide are published by the IETF and the RFC Editor under the
[IETF Trust Legal Provisions](https://trustee.ietf.org/documents/trust-legal-provisions/);
quoted guidance remains subject to those terms.
