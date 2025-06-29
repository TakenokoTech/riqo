package history

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
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

	timestamp := time.Now().Format(time.RFC3339)
	line := fmt.Sprintf("%s|%s\n", command, timestamp)
	_, err = file.WriteString(line)
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

func (h *HistoryManager) ExportHistory(filename string) error {
	content, err := os.ReadFile(h.HistoryFile)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("no history found")
		}
		return err
	}

	lines := strings.Split(string(content), "\n")
	historyData := make([]map[string]string, 0)

	for _, line := range lines {
		if line == "" {
			continue
		}
		parts := strings.SplitN(line, "|", 2)
		cmd := parts[0]
		ts := ""
		if len(parts) > 1 {
			ts = parts[1]
		}
		historyData = append(historyData, map[string]string{
			"command":   cmd,
			"timestamp": ts,
		})
	}

	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("error creating export file: %v", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(historyData); err != nil {
		return fmt.Errorf("error writing JSON to export file: %v", err)
	}

	fmt.Printf("History exported to %s\n", filename)
	return nil
}
