package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// suggestCmd represents the suggest command
var suggestCmd = &cobra.Command{
	Use:   "suggest",
	Short: "Suggest a GitHub CLI command based on natural language input",
	Long:  `The suggest command takes a natural language input from the user and provides a corresponding GitHub CLI command suggestion.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Fprintln(cmd.OutOrStdout(), "Please provide a natural language input.")
			return
		}
		fmt.Fprintf(cmd.OutOrStdout(), "Processing input: %s\n", args[0])
		fmt.Fprintln(cmd.OutOrStdout(), "Suggested GitHub CLI command: gh repo view")
		appendToHistory(fmt.Sprintf("suggest %s", args[0]))
	},
}

func init() {
	rootCmd.AddCommand(suggestCmd)
}
