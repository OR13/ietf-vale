#!/usr/bin/env bash
set -u

file="${1:-}"
idnits="${2:-idnits}"

if [[ -z "$file" ]]; then
  echo "usage: make check-draft FILE=path/to/draft.txt" >&2
  exit 2
fi

if [[ ! -f "$file" ]]; then
  echo "check-draft: file not found: $file" >&2
  exit 2
fi

if ! command -v vale >/dev/null 2>&1; then
  echo "check-draft: vale is not installed or not on PATH" >&2
  exit 2
fi

if ! command -v "$idnits" >/dev/null 2>&1; then
  echo "check-draft: idnits is not installed or not on PATH" >&2
  echo "Install it with: npm install -g @ietf-tools/idnits" >&2
  exit 2
fi

echo "== Vale =="
vale_status=0
vale --output=line "$file" || vale_status=$?

echo "== idnits =="
idnits_status=0
"$idnits" --no-color --no-progress "$file" || idnits_status=$?

if (( vale_status != 0 || idnits_status != 0 )); then
  exit 1
fi
