package history

import (
	"fmt"
	"os"
	"strings"
)

type HistoryManager struct {
	HistoryFile string
}

func (h *HistoryManager) ViewHistory() {
	content, err := os.ReadFile(h.HistoryFile)
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

func (h *HistoryManager) ClearHistory() {
	err := os.Remove(h.HistoryFile)
	if err != nil && !os.IsNotExist(err) {
		fmt.Printf("Error clearing history: %v\n", err)
		return
	}
	fmt.Println("History cleared.")
}

func (h *HistoryManager) AppendToHistory(command string) {
	file, err := os.OpenFile(h.HistoryFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
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

func (h *HistoryManager) SearchHistory(keyword string) error {
	content, err := os.ReadFile(h.HistoryFile)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("no history found")
		}
		return err
	}

	lines := strings.Split(string(content), "\n")
	fmt.Printf("Search results for keyword '%s':\n", keyword)
	found := false
	for _, line := range lines {
		if strings.Contains(line, keyword) {
			fmt.Println(line)
			found = true
		}
	}

	if !found {
		fmt.Println("No matching commands found.")
	}
	return nil
}
