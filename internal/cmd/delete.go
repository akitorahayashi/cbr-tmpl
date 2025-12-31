package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/akitorahayashi/cbr-tmpl/internal"
)

func NewDeleteCmd(storage internal.Storage) *cobra.Command {
	return &cobra.Command{
		Use:     "delete <id>",
		Aliases: []string{"rm"},
		Short:   "Delete an item",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id := args[0]
			if err := storage.Delete(id); err != nil {
				var notFoundErr *internal.ItemNotFoundError
				if errors.As(err, &notFoundErr) {
					fmt.Fprintln(cmd.ErrOrStderr(), errorStyle.Render(
						fmt.Sprintf("Error: item '%s' not found", id)))
					return err
				}
				return err
			}
			fmt.Fprintln(cmd.OutOrStdout(), successStyle.Render(
				fmt.Sprintf("Deleted '%s'", id)))
			return nil
		},
	}
}
