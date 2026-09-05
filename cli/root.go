package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func Execute() {
	root := buildRoot()

	if err := root.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'\n", err)
		os.Exit(1)
	}
}

func buildRoot() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:     "protra",
		Version: "0.0.1",
	}

	rootCmd.AddCommand(NewSimulateCommand())

	return rootCmd
}