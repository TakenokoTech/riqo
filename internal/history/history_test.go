package history

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestViewHistory(t *testing.T) {
	historyFile := "test_history.txt"
	defer os.Remove(historyFile)

	content := "gh repo view\n"
	err := os.WriteFile(historyFile, []byte(content), 0644)
	if err != nil {
		t.Fatalf("Failed to set up test history file: %v", err)
	}

	historyManager := HistoryManager{HistoryFile: historyFile}
	output := captureOutput(historyManager.ViewHistory)
	if output != "Command History:\ngh repo view\n\n" {
		t.Errorf("Expected history content, got: %s", output)
	}
}

func TestClearHistory(t *testing.T) {
	historyFile := "test_history.txt"
	defer os.Remove(historyFile)

	err := os.WriteFile(historyFile, []byte("gh repo view\n"), 0644)
	if err != nil {
		t.Fatalf("Failed to set up test history file: %v", err)
	}

	historyManager := HistoryManager{HistoryFile: historyFile}
	captureOutput(historyManager.ClearHistory)

	if _, err := os.Stat(historyFile); !os.IsNotExist(err) {
		t.Errorf("Expected history file to be deleted, but it still exists")
	}
}

func TestAppendToHistory(t *testing.T) {
	historyFile := "test_history.txt"
	defer os.Remove(historyFile)

	historyManager := HistoryManager{HistoryFile: historyFile}
	historyManager.AppendToHistory("gh repo clone")

	content, err := os.ReadFile(historyFile)
	if err != nil {
		t.Fatalf("Failed to read history file: %v", err)
	}

	expected := "gh repo clone\n"
	if string(content) != expected {
		t.Errorf("Expected '%s', got '%s'", expected, string(content))
	}
}

func TestSearchHistory(t *testing.T) {
	historyFile := "test_history.txt"
	defer os.Remove(historyFile)

	content := "gh repo view\ngh repo clone\n"
	err := os.WriteFile(historyFile, []byte(content), 0644)
	if err != nil {
		t.Fatalf("Failed to set up test history file: %v", err)
	}

	historyManager := HistoryManager{HistoryFile: historyFile}
	output := captureOutput(func() {
		historyManager.SearchHistory("clone")
	})

	expected := "Search results for keyword 'clone':\ngh repo clone\n"
	if output != expected {
		t.Errorf("Expected '%s', got '%s'", expected, output)
	}
}

func captureOutput(f func()) string {
	r, w, _ := os.Pipe()
	originalStdout := os.Stdout
	os.Stdout = w

	f()

	w.Close()
	os.Stdout = originalStdout
	var buf bytes.Buffer
	io.Copy(&buf, r)
	return buf.String()
}
