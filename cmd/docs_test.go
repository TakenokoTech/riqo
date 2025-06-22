package cmd

import (
	"os"
	"testing"
)

func TestUpdateDocs(t *testing.T) {
	historyFile := "gh_history.txt"
	docsFile = "test_docs.md"
	defer os.Remove(historyFile)
	defer os.Remove(docsFile)

	// Set up test history file
	historyContent := "suggest Show repository details\nsuggest Show repository details\nhistory view\n"
	err := os.WriteFile(historyFile, []byte(historyContent), 0644)
	if err != nil {
		t.Fatalf("Failed to set up test history file: %v", err)
	}

	// Run updateDocs function
	err = updateDocs()
	if err != nil {
		t.Fatalf("updateDocs failed: %v", err)
	}

	// Verify the generated docs file
	docsContent, err := os.ReadFile(docsFile)
	if err != nil {
		t.Fatalf("Failed to read generated docs file: %v", err)
	}

	expectedContent := "# よく使われるコマンド\n\n- suggest Show repository details （使用回数: 2 回）\n- history view （使用回数: 1 回）\n"
	if string(docsContent) != expectedContent {
		t.Errorf("Generated docs content does not match expected content.\nExpected:\n%s\nGot:\n%s", expectedContent, string(docsContent))
	}
}
