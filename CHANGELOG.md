# Changelog

## Unreleased

- Add the narrow, suggestion-level `IETF-Draft.ProtocolNames` check for common
  protocol capitalization forms.

## 0.4.4 - 2026-09-19

- Publish the opt-in `IETF-Draft-Optional` package with conservative article and
  repeated-word checks.
- Ignore inline code and fenced code blocks when checking citation references.
- Document the optional package and its installation workflow.

## 0.4.3 - 2026-09-19

- Add RPC-backed abbreviation awareness to `IETF-Email`.
- Refresh the README examples with IETF-specific draft and working-group prose.

## 0.4.2 - 2026-09-19

- Add section-title, spelling-consistency, and BCP 14 keyword review rules.
- Generate abbreviation exceptions from a pinned RPC source revision.
- Track specification areas that require structural or contextual tooling in
  GitHub issues.

## 0.4.1 - 2026-09-19

- Add `IETF-Draft.AngleBracketsURI` for bare URIs.
- Extend `IETF-Draft.ExamplePhone` with the documented UK fictitious-number range.
- Fix documentation-bundle links so the OKF validation job passes its warning budget.
- Make package creation isolated and validate release archives before upload.
- Update the reproducible-installation example to use the 0.4.1 release.
