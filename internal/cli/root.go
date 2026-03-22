package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tidy-bookmark",
	Short: "Manage bookmarks from the command line",
	Long:  "Tidy Bookmark is a CLI for organizing and managing bookmarks.",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Hello World")
		return nil
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
