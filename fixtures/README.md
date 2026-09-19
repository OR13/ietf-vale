# Fixtures

One directory per rule, named exactly after the rule file in `../IETF`.

Each directory holds a `.vale.ini` that enables only that rule, plus sample
files covering both the text the rule should flag and the near-misses it should
leave alone.

`go test ./...` fails if a rule has no fixture here, or if a directory here has
no matching rule.
