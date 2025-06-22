package cmd

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestViewHistory(t *testing.T) {
	historyFile = "test_history.txt"
	defer os.Remove(historyFile)

	content := "gh repo view\n"
	err := os.WriteFile(historyFile, []byte(content), 0644)
	if err != nil {
		t.Fatalf("Failed to set up test history file: %v", err)
	}

	output := captureOutput(viewHistory)
	if output != "Command History:\ngh repo view\n\n" {
		t.Errorf("Expected history content, got: %s", output)
	}
}

func TestClearHistory(t *testing.T) {
	historyFile = "test_history.txt"
	defer os.Remove(historyFile)

	err := os.WriteFile(historyFile, []byte("gh repo view\n"), 0644)
	if err != nil {
		t.Fatalf("Failed to set up test history file: %v", err)
	}

	captureOutput(clearHistory)

	if _, err := os.Stat(historyFile); !os.IsNotExist(err) {
		t.Errorf("Expected history file to be deleted, but it still exists")
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
