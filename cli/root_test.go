package cli

import (
	"os"
	"testing"
)

func TestIntegration(t *testing.T) {
	root := buildRoot()
	os.Args = []string{"protra", "sim", "45", "100", "--thrust", "angle:0"}
	root.Execute()
}
