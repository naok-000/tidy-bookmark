package cli

import (
	"fmt"
	"io"
	"os"
	"strconv"

	"tidy-bookmark/internal/bookmark"
	storepkg "tidy-bookmark/internal/store"

	"github.com/spf13/cobra"
)

func newRootCmd(store bookmark.Store) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "tidy-bookmark",
		Short: "Manage bookmarks from the command line",
		Long:  "Tidy Bookmark is a CLI for organizing and managing bookmarks.",
	}

	cmd.AddCommand(newAddCmd(store), newListCmd(store), newRemoveCmd(store))

	return cmd
}

func newAddCmd(store bookmark.Store) *cobra.Command {
	return &cobra.Command{
		Use:  "add URL",
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return bookmark.Add(store, args[0])
		},
	}
}

func newListCmd(store bookmark.Store) *cobra.Command {
	return &cobra.Command{
		Use:  "list",
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			list, err := bookmark.List(store)
			if err != nil {
				return err
			}

			_, err = io.WriteString(cmd.OutOrStdout(), list)
			return err
		},
	}
}

func newRemoveCmd(store bookmark.Store) *cobra.Command {
	return &cobra.Command{
		Use:  "remove ID",
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := strconv.Atoi(args[0])
			if err != nil {
				return fmt.Errorf("invalid ID: %s", args[0])
			}
			return bookmark.Remove(store, id)
		},
	}
}

func Execute() {
	err := newRootCmd(storepkg.FileStore{Path: "bookmarks.txt"}).Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
