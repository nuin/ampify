// Public Domain (-) 2010-2011 The Ampify Authors.
// See the Ampify UNLICENSE file for details.

package runtime

import (
	"amp/command"
	"os"
	"strings"
	"strconv"
	"testing"
)

func TestCPUCount(t *testing.T) {
	cpus := GetCPUCount()
	output, err := command.GetOutput(
		[]string{
			os.Getenv("AMPIFY_ROOT") + "/environ/local/bin/python",
			"-c",
			"import multiprocessing; print multiprocessing.cpu_count()"})
	if err != nil {
		t.Errorf("Couldn't call Python:\n%v", err)
		return
	}
	expected, err := strconv.Atoi(strings.TrimSpace(output))
	if err != nil {
		t.Errorf("Couldn't parse the output from Python:\n%v", err)
		return
	}
	if cpus != expected {
		t.Errorf("Got mis-matched CPU Counts: %d vs. %d", cpus, expected)
	}
}
