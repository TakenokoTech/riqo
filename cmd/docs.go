package cmd

import (
	"fmt"
	"os"
	"sort"
	"strings"

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
	historyContent, err := os.ReadFile(historyFile)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("no history found to update documentation")
		}
		return err
	}

	historyLines := strings.Split(string(historyContent), "\n")
	commandFrequency := make(map[string]int)

	for _, line := range historyLines {
		if line != "" {
			commandFrequency[line]++
		}
	}

	sortedCommands := sortCommandsByFrequency(commandFrequency)

	docsContent := "# よく使われるコマンド\n\n"
	for _, cmd := range sortedCommands {
		docsContent += fmt.Sprintf("- %s （使用回数: %d 回）\n", cmd.Command, cmd.Frequency)
	}

	return os.WriteFile(docsFile, []byte(docsContent), 0644)
}

type commandUsage struct {
	Command   string
	Frequency int
}

func sortCommandsByFrequency(freqMap map[string]int) []commandUsage {
	var commands []commandUsage
	for cmd, freq := range freqMap {
		commands = append(commands, commandUsage{Command: cmd, Frequency: freq})
	}

	sort.Slice(commands, func(i, j int) bool {
		return commands[i].Frequency > commands[j].Frequency
	})

	return commands
}
