package cli

import (
	"os"
	"testing"
)

func TestIntegration(t *testing.T) {
	root := buildRoot()
	os.Args = []string{"sim", "45", "100"}
	root.Execute()
}
