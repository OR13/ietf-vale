package main

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"

	"gopkg.in/yaml.v3"
)

// Every rule in this style has to arrive with a fixture, a recorded
// expectation, and enough metadata for a reviewer to check it against the
// source guidance. These tests enforce that contract mechanically so review can
// focus on whether the rule is *right* rather than on whether it is testable.

const styleName = "IETF"

// canonicalCmd is the vale invocation every .ct case uses. Keeping it identical
// across cases means a diff in testdata/ is always a behavior change, never a
// change in how the expectation was produced.
const canonicalCmd = "vale --output=line --sort --normalize --relative --no-global --no-exit ."

// ruleName matches the UpperCamelCase file names Vale surfaces as IETF.<Name>.
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

// ruleFiles returns the base names (without .yml) of every rule in the style.
func ruleFiles(t *testing.T) []string {
	t.Helper()

	paths, err := filepath.Glob(filepath.Join(styleName, "*.yml"))
	if err != nil {
		t.Fatalf("glob %s: %v", styleName, err)
	}

	names := make([]string, 0, len(paths))
	for _, p := range paths {
		names = append(names, strings.TrimSuffix(filepath.Base(p), ".yml"))
	}
	return names
}

func TestRuleMetadata(t *testing.T) {
	for _, name := range ruleFiles(t) {
		path := filepath.Join(styleName, name+".yml")

		t.Run(name, func(t *testing.T) {
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

func TestEveryRuleIsTested(t *testing.T) {
	for _, name := range ruleFiles(t) {
		t.Run(name, func(t *testing.T) {
			dir := filepath.Join("fixtures", name)
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
			if !strings.Contains(string(data), styleName+"."+name) {
				t.Errorf("%s: must enable %s.%s and nothing else", cfg, styleName, name)
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

			ct := filepath.Join("testdata", name+".ct")
			expected, err := os.ReadFile(ct)
			if err != nil {
				t.Fatalf("missing %s (create the fixture, then run: go test ./... -update)", ct)
			}

			body := string(expected)
			if !strings.Contains(body, "cdf ${ROOTDIR}/fixtures/"+name) {
				t.Errorf("%s: must cd into fixtures/%s", ct, name)
			}
			if !strings.Contains(body, canonicalCmd) {
				t.Errorf("%s: must invoke vale as:\n  %s", ct, canonicalCmd)
			}
		})
	}
}

func TestNoOrphanedTests(t *testing.T) {
	known := map[string]bool{}
	for _, name := range ruleFiles(t) {
		known[name] = true
	}

	entries, err := os.ReadDir("fixtures")
	if err != nil {
		t.Fatal(err)
	}
	for _, e := range entries {
		if !e.IsDir() {
			continue
		}
		if !known[e.Name()] {
			t.Errorf("fixtures/%s/ has no matching rule %s/%s.yml", e.Name(), styleName, e.Name())
		}
	}

	cases, err := filepath.Glob(filepath.Join("testdata", "*.ct"))
	if err != nil {
		t.Fatal(err)
	}
	for _, c := range cases {
		name := strings.TrimSuffix(filepath.Base(c), ".ct")
		if !known[name] {
			t.Errorf("%s has no matching rule %s/%s.yml", c, styleName, name)
		}
	}
}

// TestStyleIsPackageable guards the files that make the release archive usable
// as a Vale package.
func TestStyleIsPackageable(t *testing.T) {
	meta := filepath.Join(styleName, "meta.json")
	if _, err := os.Stat(meta); err != nil {
		t.Fatalf("missing %s: Vale needs it to install the package", meta)
	}
}
