STYLES := IETF-Draft IETF-Email IETF-Draft-Optional

# Pinned so a linter release cannot change what CI accepts without a commit.
OKF_LINT_VERSION ?= 0.1.0

.PHONY: new test lint lint-docs update update-abbreviations update-email-abbreviations package package-check clean corpus-mail check-sources

# Scaffold a new rule, its fixture, and its .ct case:
#   make new RULE=Ellipses               (defaults to the IETF-Draft style)
#   make new RULE=Slang STYLE=IETF-Email
new:
	@./scripts/new-rule.sh $(RULE)

# Full suite: fixture snapshots, rule contract checks, coverage manifests.
test:
	go test ./...

lint:
	yamllint -c .yamllint.yml $(STYLES) coverage/
	vale README.md CONTRIBUTING.md docs/

# The specifications are an OKF bundle; okf-lint checks that structure.
# Two warnings are expected: we target OKF v0.2 against a v0.1 linter, and the
# bundle has no log.md because git already records the history.
lint-docs:
	npx --yes @thisismydesign/okf-lint@$(OKF_LINT_VERSION) docs --max-warnings 2

# Fail if a rule cites an RFC that has since been obsoleted.
check-sources:
	@python3 scripts/check-sources.py

# Measure IETF-Email against real list mail from the IETF's anonymous IMAP
# archive. Measures behavior; never decides which rules exist.
corpus-mail:
	@python3 scripts/corpus-mail.py --email "$(EMAIL)"

# Regenerate testdata/*.ct from current rule behavior. Review the diff.
update:
	go test ./... -update

update-abbreviations:
	@python3 scripts/update-abbreviations.py

update-email-abbreviations:
	@python3 scripts/update-abbreviations.py --output IETF-Email/Abbreviations.yml --level suggestion

# Build the archives users install, license included. Build from a temporary
# directory so packaging never leaves generated LICENSE files in the styles.
package:
	@tmp=$$(mktemp -d); \
	trap 'rm -rf "$$tmp"' EXIT; \
	for s in $(STYLES); do \
		cp -R $$s "$$tmp/$$s"; \
		cp LICENSE "$$tmp/$$s/LICENSE"; \
		(cd "$$tmp" && zip -qr "$(CURDIR)/$$s.zip" "$$s" -x "*.DS_Store"); \
		echo "built $$s.zip"; \
	done

# Validate the files that will be uploaded by the release workflow.
package-check: package
	@for s in $(STYLES); do \
		unzip -t $$s.zip >/dev/null; \
		unzip -l $$s.zip | grep -q "$$s/meta.json"; \
		unzip -l $$s.zip | grep -q "$$s/LICENSE"; \
		echo "checked $$s.zip"; \
	done

clean:
	rm -f $(addsuffix .zip,$(STYLES)) $(addsuffix /LICENSE,$(STYLES))
