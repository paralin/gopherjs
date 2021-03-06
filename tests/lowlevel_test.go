package tests_test

import (
	"bytes"
	"io/ioutil"
	"os/exec"
	"path/filepath"
	"runtime"
	"testing"
)

// Test for internalization/externalization of time.Time/Date when time package is imported
// but time.Time is unused, causing it to be DCEed (or time package not imported at all).
//
// See https://github.com/gopherjs/gopherjs/issues/279.
func TestTimeInternalizationExternalization(t *testing.T) {
	if runtime.GOARCH == "js" {
		t.Skip("test meant to be run using normal Go compiler (needs os/exec)")
	}

	got, err := exec.Command("gopherjs", "run", filepath.Join("testdata", "time_inexternalization.go")).Output()
	if err != nil {
		t.Fatalf("%v:\n%s", err, got)
	}

	want, err := ioutil.ReadFile(filepath.Join("testdata", "time_inexternalization.out"))
	if err != nil {
		t.Fatalf("error reading .out file: %v", err)
	}

	if !bytes.Equal(got, want) {
		t.Fatalf("got != want:\ngot:\n%s\nwant:\n%s", got, want)
	}
}

func TestDeferBuiltin(t *testing.T) {
	if runtime.GOARCH == "js" {
		t.Skip("test meant to be run using normal Go compiler (needs os/exec)")
	}

	got, err := exec.Command("gopherjs", "run", filepath.Join("testdata", "defer_builtin.go")).CombinedOutput()
	if err != nil {
		t.Fatalf("%v:\n%s", err, got)
	}
}
