package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// realtimeCmd represents the realtime command
var realtimeCmd = &cobra.Command{
	Use:   "realtime",
	Short: "Provide real-time command suggestions",
	Long:  `The realtime command listens for user input in real-time and provides suggestions for CLI commands based on the input.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Real-time command suggestion mode. Type 'exit' to quit.")
		reader := bufio.NewReader(os.Stdin)

		for {
			fmt.Print("> ")
			input, err := reader.ReadString('\n')
			if err != nil {
				fmt.Printf("Error reading input: %v\n", err)
				continue
			}

			input = strings.TrimSpace(input)
			if input == "exit" {
				fmt.Println("Exiting real-time mode.")
				break
			}

			suggestion := generateSuggestion(input)
			fmt.Printf("Suggested command for '%s': %s\n", input, suggestion)
		}
	},
}

func generateSuggestion(input string) string {
	if strings.Contains(input, "repo") {
		return "gh repo view"
	} else if strings.Contains(input, "history") {
		return "gh history view"
	}
	return "[No suggestion available]"
}

func init() {
	rootCmd.AddCommand(realtimeCmd)
}
