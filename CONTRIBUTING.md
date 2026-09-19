# Contributing

Thanks for helping build an IETF style for [Vale](https://vale.sh).

Every rule here changes what a warning means for someone writing an
Internet-Draft or posting to a working group list, so the bar is: **each rule
cites published IETF guidance, and its exact behavior is recorded in a test a
reviewer can read as a diff.**

Start with the specifications in [`docs/`](docs/index.md). A rule that is not in
[draft style](docs/draft-style.md) or [email style](docs/email-style.md) needs a
specification change first. That is where the argument about whether a rule
belongs happens, before anyone writes YAML. The
[evidence policy](docs/evidence-policy.md) decides its severity level and lists
the evidence your PR has to show.

## Before you open a PR

Install [Vale](https://vale.sh/docs/install) and Go 1.21+, then:

```bash
make test
make lint
make lint-docs   # only if you changed docs/
```

## Adding a rule

A rule is four files. Scaffold them with:

```bash
make new RULE=Ellipses                # in the IETF-Draft style
make new RULE=Slang STYLE=IETF-Email  # in the IETF-Email style
```

That writes the stubs below; fill them in. Using `IETF-Draft.Ellipses` as the
example:

1. **The rule**, `IETF-Draft/Ellipses.yml`. It must set `extends`, `message`,
   `level`, and a `link` pointing at the guidance it implements (an RFC section,
   or a page of the [RFC Editor style guide](https://www.rfc-editor.org/styleguide/)).
   `make test` fails without all four.

2. **A fixture**, `fixtures/IETF-Draft/Ellipses/`, with a `.vale.ini` that
   enables only this rule, plus at least one sample file:

   ```ini
   StylesPath = ../../../

   MinAlertLevel = suggestion

   [*.md]
   IETF-Draft.Ellipses = YES
   ```

   Include both text that should be flagged **and** text that should not. The
   near-misses are what stop a rule from over-matching later.

3. **The expectation**, `testdata/IETF-Draft.Ellipses.ct`. It opens with the
   two commands every case shares, and `make update` fills in the rest:

   ```
   $ cdf ${ROOTDIR}/fixtures/IETF-Draft/Ellipses
   $ vale --output=line --sort --normalize --relative --no-global --no-exit .
   ```

   ```bash
   make update          # go test ./... -update
   ```

   Then read the generated file. It is the record of what the rule actually
   does, line and column included, and it is what reviewers will read. Never
   hand-edit the recorded alerts; change the rule or the fixture and rerun.

4. **Coverage**: flip the matching key in `coverage/IETF-Draft/` to `true` and
   name the rule in a comment above it:

   ```yaml
   # Ellipses.yml
   punctuation: true
   ```

   The coverage test fails if a comment names a rule that no longer exists, so
   the manifest can't drift away from reality.

## Changing an existing rule

Run `make update` and commit the resulting `testdata/` diff **in the same
commit** as the rule change. That diff is the review: it shows every alert that
appeared, disappeared, or moved. A rule change with no testdata diff means the
fixture doesn't cover the change, so extend the fixture.

## Review expectations

The [evidence policy](docs/evidence-policy.md) lists the seven things a rule PR
has to show. On top of those, reviewers will ask:

- **Is it in the guidance?** Follow the `link`. A rule that encodes a personal
  preference, however reasonable, belongs in a user's own style, not here.
- **Is the source current?** Several obvious sources are obsolete: cite RFC 9245
  not RFC 3005, RFC 9945 not RFC 3934, RFC 9542 not RFC 7042.
- **Is it in the right style?** `IETF-Draft` rules must not fire on mail, and
  `IETF-Email` rules must not fire on drafts.
- **Is the level right?** `error` is for things the RFC Editor will change;
  `warning` for guidance with exceptions; `suggestion` for preferences.
- **Does the message tell the writer what to do?** Name the fix, not the
  offense.
- **What does it do to real drafts?** For anything broad, run it over a few
  published RFCs or drafts and say what you saw in the PR description. False
  positives at IETF scale are expensive.
- **Does the testdata diff match the description?** Unexplained changes are
  blocking.

## Vale version

CI pins a Vale version (`VALE_VERSION` in `.github/workflows/ci.yml`) because
recorded expectations are only meaningful against a known engine. Bump it in its
own PR, with the resulting `testdata/` diff and nothing else.
