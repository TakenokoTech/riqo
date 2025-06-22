package cmd

import (
	"fmt"
	"os"
	"riqo/internal/docs"

	"github.com/spf13/cobra"
)

var docsFile = "gh_docs.md"

// docsCmd represents the docs command
var docsCmd = &cobra.Command{
	Use:   "docs",
	Short: "Automatically organize and update documentation",
	Long:  `The docs command analyzes command usage history and updates the documentation file to prioritize frequently used commands.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := updateDocs()
		if err != nil {
			fmt.Printf("Error updating documentation: %v\n", err)
			return
		}
		fmt.Println("Documentation updated successfully.")
	},
}

func init() {
	rootCmd.AddCommand(docsCmd)
}

func updateDocs() error {
	historyContent, err := os.ReadFile("gh_history.txt")
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("no history found to update documentation")
		}
		return err
	}

	updater := docs.DocsUpdater{DocsFile: docsFile}
	return updater.UpdateDocs(string(historyContent))
}
