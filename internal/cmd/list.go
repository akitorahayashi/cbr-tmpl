package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/akitorahayashi/cbr-tmpl/internal"
)

func NewListCmd(storage internal.Storage) *cobra.Command {
	return &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "List all items",
		RunE: func(cmd *cobra.Command, args []string) error {
			items, err := storage.List()
			if err != nil {
				return err
			}
			if len(items) == 0 {
				fmt.Fprintln(cmd.OutOrStdout(), dimStyle.Render("No items found"))
				return nil
			}
			for _, id := range items {
				fmt.Fprintf(cmd.OutOrStdout(), "  %s\n", id)
			}
			return nil
		},
	}
}
