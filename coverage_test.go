package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"testing"
)

// The coverage/ directory tracks, topic by topic, which parts of the IETF's
// published editorial guidance each style implements. It holds one directory
// per style; each file inside mirrors one source document or top-level section,
// and each key is a subtopic set to true or false, optionally followed by a
// comment naming the rules that implement it.
//
// This test reports the resulting metric and enforces the invariants that keep
// it honest: values must be canonical booleans, and every rule named in a
// comment must still exist. A rule that gets renamed or merged away otherwise
// leaves the manifest silently claiming coverage it no longer has.

var ruleRef = regexp.MustCompile(`([A-Za-z][A-Za-z0-9]*)\.yml`)

type tally struct{ covered, total int }

func (t tally) pct() float64 {
	if t.total == 0 {
		return 0
	}
	return 100 * float64(t.covered) / float64(t.total)
}

func parseManifest(t *testing.T, style, path string) tally {
	t.Helper()

	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read %s: %v", path, err)
	}

	var got tally
	for i, line := range strings.Split(string(data), "\n") {
		lineNo := i + 1

		// A comment names the implementing rules; check they still exist.
		if idx := strings.Index(line, "#"); idx >= 0 {
			for _, m := range ruleRef.FindAllStringSubmatch(line[idx:], -1) {
				rule := filepath.Join(style, m[1]+".yml")
				if _, err := os.Stat(rule); err != nil {
					t.Errorf("%s:%d: names %s.yml, which doesn't exist", path, lineNo, m[1])
				}
			}
			line = line[:idx]
		}

		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		key, value, found := strings.Cut(line, ":")
		if !found {
			t.Errorf("%s:%d: not a key/value pair: %q", path, lineNo, line)
			continue
		}

		switch strings.TrimSpace(value) {
		case "true":
			got.covered++
			got.total++
		case "false":
			got.total++
		default:
			t.Errorf("%s:%d: %q must be exactly true or false, got %q",
				path, lineNo, strings.TrimSpace(key), strings.TrimSpace(value))
		}
	}

	return got
}

func TestCoverage(t *testing.T) {
	for _, style := range styles(t) {
		paths, err := filepath.Glob(filepath.Join("coverage", style, "*.yml"))
		if err != nil {
			t.Fatal(err)
		}
		if len(paths) == 0 {
			t.Logf("%s: no coverage manifests yet", style)
			continue
		}
		sort.Strings(paths)

		var total tally
		var report []string

		width := 0
		for _, p := range paths {
			if n := len(filepath.Base(p)); n > width {
				width = n
			}
		}

		for _, p := range paths {
			got := parseManifest(t, style, p)
			report = append(report, fmt.Sprintf("  %-*s  %3d/%-3d  %5.1f%%",
				width, filepath.Base(p), got.covered, got.total, got.pct()))

			total.covered += got.covered
			total.total += got.total
		}

		summary := fmt.Sprintf("\n\n  %s overall:  %d/%d (%.1f%%)",
			style, total.covered, total.total, total.pct())

		t.Log(style + " coverage\n" + strings.Join(report, "\n") + summary)
	}
}
