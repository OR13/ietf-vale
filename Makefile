STYLES := IETF-Draft IETF-Email

# Pinned so a linter release cannot change what CI accepts without a commit.
OKF_LINT_VERSION ?= 0.1.0

.PHONY: new test lint lint-docs update package clean corpus-mail

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

# Measure IETF-Email against real list mail from the IETF's anonymous IMAP
# archive. Measures behavior; never decides which rules exist.
corpus-mail:
	@python3 scripts/corpus-mail.py --email "$(EMAIL)"

# Regenerate testdata/*.ct from current rule behavior. Review the diff.
update:
	go test ./... -update

# Build the archives users install, license included.
package:
	@for s in $(STYLES); do \
		cp LICENSE $$s/LICENSE; \
		zip -qr $$s.zip $$s -x "*.DS_Store"; \
		echo "built $$s.zip"; \
	done

clean:
	rm -f $(addsuffix .zip,$(STYLES)) $(addsuffix /LICENSE,$(STYLES))
