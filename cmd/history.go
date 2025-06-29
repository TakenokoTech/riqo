package cmd

import (
	"fmt"

	"riqo/internal/history"

	"github.com/spf13/cobra"
)

var historyFile = "gh_history.txt"
var searchKeyword string

// historyCmd represents the history command
var historyCmd = &cobra.Command{
	Use:   "history",
	Short: "Manage command history",
	Long:  `The history command allows you to view, clear, and search the history of executed CLI commands.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Please provide a subcommand (e.g., view, clear, search, export).")
			return
		}

		manager := history.HistoryManager{HistoryFile: historyFile}

		switch args[0] {
		case "view":
			manager.ViewHistory()
		case "clear":
			manager.ClearHistory()
		case "search":
			if searchKeyword == "" {
				fmt.Println("Please provide a keyword to search using the --keyword flag.")
				return
			}
			err := manager.SearchHistory(searchKeyword)
			if err != nil {
				fmt.Printf("Error searching history: %v\n", err)
			}
		case "export":
			err := manager.ExportHistory("history_export.json")
			if err != nil {
				fmt.Printf("Error exporting history: %v\n", err)
			}
		default:
			fmt.Printf("Unknown subcommand: %s\n", args[0])
		}
	},
}

func init() {
	historyCmd.Flags().StringVarP(&searchKeyword, "keyword", "k", "", "Keyword to search in history")
	rootCmd.AddCommand(historyCmd)
}
