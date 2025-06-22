package cmd

import (
	"fmt"
	"riqo/internal/history"

	"github.com/spf13/cobra"
)

// suggestCmd represents the suggest command
var suggestCmd = &cobra.Command{
	Use:   "suggest",
	Short: "Suggest a CLI command based on natural language input",
	Long:  `The suggest command takes a natural language input from the user and provides a corresponding CLI command suggestion.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Fprintln(cmd.OutOrStdout(), "Please provide a natural language input.")
			return
		}
		fmt.Fprintf(cmd.OutOrStdout(), "Processing input: %s\n", args[0])
		fmt.Fprintln(cmd.OutOrStdout(), "Suggested CLI command: gh repo view")

		history := history.HistoryManager{HistoryFile: "gh_history.txt"}
		history.AppendToHistory(fmt.Sprintf("suggest %s", args[0]))
	},
}

func init() {
	rootCmd.AddCommand(suggestCmd)
}
