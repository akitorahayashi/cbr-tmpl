package cmd

import (
	"github.com/spf13/cobra"

	"github.com/akitorahayashi/cbr-tmpl/internal"
)

var version = "0.1.0"

func NewRootCmd(storage internal.Storage) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:     "cbr-tmpl",
		Short:   "A minimal Go CLI template using Cobra",
		Version: version,
	}

	rootCmd.AddCommand(NewAddCmd(storage))
	rootCmd.AddCommand(NewListCmd(storage))
	rootCmd.AddCommand(NewDeleteCmd(storage))

	return rootCmd
}
