package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/akitorahayashi/cbr-tmpl/internal"
)

func NewAddCmd(storage internal.Storage) *cobra.Command {
	var content string

	cmd := &cobra.Command{
		Use:     "add <id>",
		Aliases: []string{"a"},
		Short:   "Add a new item",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id := args[0]
			if err := storage.Add(id, content); err != nil {
				var existsErr *internal.ItemExistsError
				if errors.As(err, &existsErr) {
					fmt.Fprintln(cmd.ErrOrStderr(), errorStyle.Render(
						fmt.Sprintf("Error: item '%s' already exists", id)))
					return err
				}
				return err
			}
			fmt.Fprintln(cmd.OutOrStdout(), successStyle.Render(
				fmt.Sprintf("Added '%s'", id)))
			return nil
		},
	}

	cmd.Flags().StringVarP(&content, "content", "c", "", "Content of the item")
	cmd.MarkFlagRequired("content")

	return cmd
}
