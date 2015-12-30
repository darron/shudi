package cmd

import (
	"fmt"
	"testing"
)

func TestBuildPath(t *testing.T) {
	t.Log("Testing the path from BuildPath().")
	Prefix = "testing"
	Exec = "w"
	hostname := GetHostname()
	sha := "50e721e49c013f00c62cf59f2163542a9d8df02464efeb615d31051b0fddc326"
	path := fmt.Sprintf("%s/%s/%s", Prefix, sha, hostname)
	comparisonPath := BuildPath()
	if comparisonPath != path {
		t.Errorf("The path should have matched.")
	}
}
