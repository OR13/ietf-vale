package main

import (
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"testing"

	"gopkg.in/yaml.v3"
)

// Every rule in every style has to arrive with a fixture, a recorded
// expectation, and enough metadata for a reviewer to check it against the
// source guidance. These tests enforce that contract mechanically so review can
// focus on whether the rule is *right* rather than on whether it is testable.
//
// This repository ships more than one style: IETF-Draft lints Internet-Draft
// prose, IETF-Email lints mailing list prose, and they share no rules. Vale has
// no way to enable part of a style -- rule-name wildcards do not work -- so the
// split has to be separate style directories, each packaged on its own.

// canonicalCmd is the vale invocation every .ct case uses. Keeping it identical
// across cases means a diff in testdata/ is always a behavior change, never a
// change in how the expectation was produced.
const canonicalCmd = "vale --output=line --sort --normalize --relative --no-global --no-exit ."

// ruleName matches the UpperCamelCase file names Vale surfaces as <Style>.<Name>.
var ruleName = regexp.MustCompile(`^[A-Z][A-Za-z0-9]*$`)

// extendsPoints are the Vale extension points a rule may build on.
var extendsPoints = map[string]bool{
	"capitalization": true,
	"conditional":    true,
	"consistency":    true,
	"existence":      true,
	"metric":         true,
	"occurrence":     true,
	"readability":    true,
	"repetition":     true,
	"script":         true,
	"sequence":       true,
	"spelling":       true,
	"substitution":   true,
}

var levels = map[string]bool{"suggestion": true, "warning": true, "error": true}

type rule struct {
	Extends string `yaml:"extends"`
	Message string `yaml:"message"`
	Level   string `yaml:"level"`
	Link    string `yaml:"link"`
}

// styles returns every style directory: any directory holding a meta.json,
// which is what makes it installable as a Vale package.
func styles(t *testing.T) []string {
	t.Helper()

	metas, err := filepath.Glob(filepath.Join("*", "meta.json"))
	if err != nil {
		t.Fatalf("glob styles: %v", err)
	}
	if len(metas) == 0 {
		t.Fatal("no style directories found: a style is a directory containing meta.json")
	}

	var names []string
	for _, m := range metas {
		names = append(names, filepath.Dir(m))
	}
	sort.Strings(names)
	return names
}

// rulesIn returns the base names (without .yml) of every rule in one style.
func rulesIn(t *testing.T, style string) []string {
	t.Helper()

	paths, err := filepath.Glob(filepath.Join(style, "*.yml"))
	if err != nil {
		t.Fatalf("glob %s: %v", style, err)
	}

	names := make([]string, 0, len(paths))
	for _, p := range paths {
		names = append(names, strings.TrimSuffix(filepath.Base(p), ".yml"))
	}
	return names
}

func TestRuleMetadata(t *testing.T) {
	for _, style := range styles(t) {
		for _, name := range rulesIn(t, style) {
			path := filepath.Join(style, name+".yml")

			t.Run(style+"/"+name, func(t *testing.T) {
				if !ruleName.MatchString(name) {
					t.Errorf("%s: name must be UpperCamelCase with no separators", path)
				}

				data, err := os.ReadFile(path)
				if err != nil {
					t.Fatalf("read %s: %v", path, err)
				}

				var r rule
				if err := yaml.Unmarshal(data, &r); err != nil {
					t.Fatalf("%s: invalid YAML: %v", path, err)
				}

				if !extendsPoints[r.Extends] {
					t.Errorf("%s: extends %q is not a Vale extension point", path, r.Extends)
				}
				if strings.TrimSpace(r.Message) == "" {
					t.Errorf("%s: needs a message", path)
				}
				if !levels[r.Level] {
					t.Errorf("%s: level must be suggestion, warning, or error, got %q", path, r.Level)
				}
				// The link is what a reviewer follows to confirm the rule matches
				// published guidance, so it is not optional here.
				if !strings.HasPrefix(r.Link, "http") {
					t.Errorf("%s: needs a link to the guidance it implements, got %q", path, r.Link)
				}
			})
		}
	}
}

func TestEveryRuleIsTested(t *testing.T) {
	for _, style := range styles(t) {
		for _, name := range rulesIn(t, style) {
			t.Run(style+"/"+name, func(t *testing.T) {
				dir := filepath.Join("fixtures", style, name)
				if _, err := os.Stat(dir); err != nil {
					t.Fatalf("missing fixture directory %s/", dir)
				}

				cfg := filepath.Join(dir, ".vale.ini")
				data, err := os.ReadFile(cfg)
				if err != nil {
					t.Fatalf("missing %s", cfg)
				}
				// The fixture must isolate its rule; otherwise an unrelated rule's
				// alerts leak into this rule's recorded expectation.
				if !strings.Contains(string(data), style+"."+name) {
					t.Errorf("%s: must enable %s.%s and nothing else", cfg, style, name)
				}

				entries, err := os.ReadDir(dir)
				if err != nil {
					t.Fatal(err)
				}
				samples := 0
				for _, e := range entries {
					if !e.IsDir() && !strings.HasPrefix(e.Name(), ".") {
						samples++
					}
				}
				if samples == 0 {
					t.Errorf("%s/: needs at least one sample file for vale to lint", dir)
				}

				ct := filepath.Join("testdata", style+"."+name+".ct")
				expected, err := os.ReadFile(ct)
				if err != nil {
					t.Fatalf("missing %s (create the fixture, then run: go test ./... -update)", ct)
				}

				body := string(expected)
				if !strings.Contains(body, "cdf ${ROOTDIR}/fixtures/"+style+"/"+name) {
					t.Errorf("%s: must cd into fixtures/%s/%s", ct, style, name)
				}
				if !strings.Contains(body, canonicalCmd) {
					t.Errorf("%s: must invoke vale as:\n  %s", ct, canonicalCmd)
				}
			})
		}
	}
}

func TestNoOrphanedTests(t *testing.T) {
	known := map[string]bool{}
	for _, style := range styles(t) {
		for _, name := range rulesIn(t, style) {
			known[style+"."+name] = true
		}
	}

	fixtures, err := filepath.Glob(filepath.Join("fixtures", "*", "*"))
	if err != nil {
		t.Fatal(err)
	}
	for _, f := range fixtures {
		info, err := os.Stat(f)
		if err != nil || !info.IsDir() {
			continue
		}
		style, name := filepath.Base(filepath.Dir(f)), filepath.Base(f)
		if !known[style+"."+name] {
			t.Errorf("%s/ has no matching rule %s/%s.yml", f, style, name)
		}
	}

	cases, err := filepath.Glob(filepath.Join("testdata", "*.ct"))
	if err != nil {
		t.Fatal(err)
	}
	for _, c := range cases {
		id := strings.TrimSuffix(filepath.Base(c), ".ct")
		if !known[id] {
			style, name, ok := strings.Cut(id, ".")
			if !ok {
				t.Errorf("%s: name a case <Style>.<Rule>.ct", c)
				continue
			}
			t.Errorf("%s has no matching rule %s/%s.yml", c, style, name)
		}
	}
}

// TestStylesArePackageable guards the files that make each release archive
// usable as a Vale package.
func TestStylesArePackageable(t *testing.T) {
	for _, style := range styles(t) {
		// A style name reaches users as the <Style>.<Rule> identifier and as the
		// value of BasedOnStyles. A dot in the name makes that identifier
		// ambiguous, and Vale silently ignores the style.
		if strings.Contains(style, ".") {
			t.Errorf("%s: a style name must not contain a dot", style)
		}

		meta := filepath.Join(style, "meta.json")
		if _, err := os.Stat(meta); err != nil {
			t.Fatalf("missing %s: Vale needs it to install the package", meta)
		}
	}
}
