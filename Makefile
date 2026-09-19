.PHONY: new test lint update package clean

# Scaffold a new rule, its fixture, and its .ct case: make new RULE=Ellipses
new:
	@./scripts/new-rule.sh $(RULE)

# Full suite: fixture snapshots, rule contract checks, coverage manifests.
test:
	go test ./...

lint:
	yamllint -c .yamllint.yml IETF/ coverage/

# Regenerate testdata/*.ct from current rule behavior. Review the diff.
update:
	go test ./... -update

# Build the archive users install, license included.
package:
	cp LICENSE IETF/LICENSE
	zip -r IETF.zip IETF -x "*.DS_Store"

clean:
	rm -f IETF.zip IETF/LICENSE
