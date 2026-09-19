#!/usr/bin/env bash
# Scaffold a new rule: the rule file, an isolated fixture, and a .ct case
# holding the canonical vale invocation. Fill in the stubs, then run
# `make update` to record what the rule actually does.
set -euo pipefail

name="${1:-}"
if [[ ! "$name" =~ ^[A-Z][A-Za-z0-9]*$ ]]; then
  echo "usage: make new RULE=UpperCamelCaseName" >&2
  exit 1
fi

root="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "$root"

for path in "IETF/$name.yml" "fixtures/$name" "testdata/$name.ct"; do
  if [[ -e "$path" ]]; then
    echo "$path already exists" >&2
    exit 1
  fi
done

cat > "IETF/$name.yml" <<EOF
extends: existence
message: "TODO: tell the writer what to do instead."
link: https://www.rfc-editor.org/rfc/rfc7322.html
level: suggestion
ignorecase: true
tokens:
  - TODO
EOF

mkdir -p "fixtures/$name"
cat > "fixtures/$name/.vale.ini" <<EOF
StylesPath = ../../

MinAlertLevel = suggestion

[*.md]
IETF.$name = YES
EOF

cat > "fixtures/$name/test.md" <<EOF
# $name

TODO: text this rule should flag.

TODO: near-misses this rule must leave alone.
EOF

cat > "testdata/$name.ct" <<EOF
\$ cdf \${ROOTDIR}/fixtures/$name
\$ vale --output=line --sort --normalize --relative --no-global --no-exit .
EOF

cat <<EOF
Scaffolded:
  IETF/$name.yml
  fixtures/$name/.vale.ini
  fixtures/$name/test.md
  testdata/$name.ct

Next:
  1. Write the rule and the fixture samples.
  2. make update    # records the alerts in testdata/$name.ct
  3. Read that diff, then flip the matching key in coverage/ to true.
EOF
