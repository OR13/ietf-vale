package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/google/go-cmdtest"
)

var update = flag.Bool("update", false, "replace test file contents with output")

// Test runs every testdata/<Rule>.ct case: each one cds into the matching
// fixture directory, runs vale with only that rule enabled, and compares the
// output against the recorded expectation. A rule change that shifts an alert's
// position, level, or message shows up as a diff in review.
func Test(t *testing.T) {
	cases, err := filepath.Glob(filepath.Join("testdata", "*.ct"))
	if err != nil {
		t.Fatal(err)
	}
	if len(cases) == 0 {
		t.Skip("no testdata/*.ct cases yet")
	}

	ts, err := cmdtest.Read("testdata")
	if err != nil {
		t.Fatal(err)
	}

	ts.Setup = func(_ string) error {
		_, testFileName, _, ok := runtime.Caller(0)
		if !ok {
			return fmt.Errorf("failed get real working directory from caller")
		}

		// ROOTDIR is the repository root, so each case can cd into fixtures/<Rule>.
		projectRootDir := filepath.Dir(testFileName)
		if err := os.Setenv("ROOTDIR", projectRootDir); err != nil {
			return fmt.Errorf("failed change 'ROOTDIR' to caller working directory: %v", err)
		}

		return nil
	}

	path, err := exec.LookPath("vale")
	if err != nil {
		t.Fatal(err)
	}

	ts.Commands["vale"] = cmdtest.Program(path)
	ts.Commands["cdf"] = cmdtest.InProcessProgram("cdf", cdf)

	ts.Run(t, *update)
}
