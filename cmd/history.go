package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var historyFile = "gh_history.txt"

// historyCmd represents the history command
var historyCmd = &cobra.Command{
	Use:   "history",
	Short: "Manage command history",
	Long:  `The history command allows you to view and manage the history of executed GitHub CLI commands.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Please provide a subcommand (e.g., view, clear).")
			return
		}

		switch args[0] {
		case "view":
			viewHistory()
		case "clear":
			clearHistory()
		default:
			fmt.Printf("Unknown subcommand: %s\n", args[0])
		}
	},
}

func init() {
	rootCmd.AddCommand(historyCmd)
}

func viewHistory() {
	content, err := os.ReadFile(historyFile)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("No history found.")
		} else {
			fmt.Printf("Error reading history: %v\n", err)
		}
		return
	}
	fmt.Println("Command History:")
	fmt.Println(string(content))
}

func clearHistory() {
	err := os.Remove(historyFile)
	if err != nil && !os.IsNotExist(err) {
		fmt.Printf("Error clearing history: %v\n", err)
		return
	}
	fmt.Println("History cleared.")
}

func appendToHistory(command string) {
	file, err := os.OpenFile(historyFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Error appending to history: %v\n", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(command + "\n")
	if err != nil {
		fmt.Printf("Error writing to history: %v\n", err)
	}
}
