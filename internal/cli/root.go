package cli

import (
	"fmt"
	"io"
	"os"
	"strconv"

	"tidy-bookmark/internal/bookmark"

	"github.com/spf13/cobra"
)

func newRootCmd() *cobra.Command {
	list := &bookmark.BookmarkList{}

	cmd := &cobra.Command{
		Use:   "tidy-bookmark",
		Short: "Manage bookmarks from the command line",
		Long:  "Tidy Bookmark is a CLI for organizing and managing bookmarks.",
	}

	cmd.AddCommand(newAddCmd(list), newListCmd(list), newRemoveCmd(list))

	return cmd
}

func newAddCmd(list *bookmark.BookmarkList) *cobra.Command {
	return &cobra.Command{
		Use:  "add URL",
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			list.Add(args[0])
			return nil
		},
	}
}

func newListCmd(list *bookmark.BookmarkList) *cobra.Command {
	return &cobra.Command{
		Use:  "list",
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			_, err := io.WriteString(cmd.OutOrStdout(), list.Show())
			return err
		},
	}
}

func newRemoveCmd(list *bookmark.BookmarkList) *cobra.Command {
	return &cobra.Command{
		Use:  "remove ID",
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := strconv.Atoi(args[0])
			if err != nil {
				return fmt.Errorf("invalid ID: %s", args[0])
			}
			list.Remove(id)
			return nil
		},
	}
}

func Execute() {
	err := newRootCmd().Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
