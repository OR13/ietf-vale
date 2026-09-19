#!/usr/bin/env bash
# Scaffold a new rule: the rule file, an isolated fixture, and a .ct case
# holding the canonical vale invocation. Fill in the stubs, then run
# `make update` to record what the rule actually does.
set -euo pipefail

style="${STYLE:-IETF-Draft}"
name="${1:-}"

if [[ ! "$name" =~ ^[A-Z][A-Za-z0-9]*$ ]]; then
  echo "usage: make new RULE=UpperCamelCaseName [STYLE=IETF-Draft|IETF-Email]" >&2
  exit 1
fi

root="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "$root"

if [[ ! -f "$style/meta.json" ]]; then
  echo "unknown style '$style' (a style is a directory containing meta.json)" >&2
  exit 1
fi

for path in "$style/$name.yml" "fixtures/$style/$name" "testdata/$style.$name.ct"; do
  if [[ -e "$path" ]]; then
    echo "$path already exists" >&2
    exit 1
  fi
done

cat > "$style/$name.yml" <<EOF
extends: existence
message: "TODO: tell the writer what to do instead."
link: https://www.rfc-editor.org/rfc/rfc7322.html
level: suggestion
ignorecase: true
tokens:
  - TODO
EOF

mkdir -p "fixtures/$style/$name"
cat > "fixtures/$style/$name/.vale.ini" <<EOF
StylesPath = ../../../

MinAlertLevel = suggestion

[*.md]
$style.$name = YES
EOF

cat > "fixtures/$style/$name/test.md" <<EOF
# $name

TODO: text this rule should flag.

TODO: near-misses this rule must leave alone.
EOF

cat > "testdata/$style.$name.ct" <<EOF
\$ cdf \${ROOTDIR}/fixtures/$style/$name
\$ vale --output=line --sort --normalize --relative --no-global --no-exit .
EOF

cat <<EOF
Scaffolded:
  $style/$name.yml
  fixtures/$style/$name/.vale.ini
  fixtures/$style/$name/test.md
  testdata/$style.$name.ct

Next:
  1. Write the rule and the fixture samples.
  2. make update    # records the alerts in testdata/$style.$name.ct
  3. Read that diff, then flip the matching key in coverage/$style/ to true.
EOF
