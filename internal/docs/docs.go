package docs

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

type DocsUpdater struct {
	DocsFile string
}

func (d *DocsUpdater) UpdateDocs(historyContent string) error {
	historyLines := strings.Split(historyContent, "\n")
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

	return os.WriteFile(d.DocsFile, []byte(docsContent), 0644)
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
