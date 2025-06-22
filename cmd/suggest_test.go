package cmd

import (
	"bytes"
	"testing"
)

func TestSuggestCommand(t *testing.T) {
	var output bytes.Buffer
	rootCmd.SetOut(&output)

	// Test case: No arguments provided
	rootCmd.SetArgs([]string{"suggest"})
	err := rootCmd.Execute()
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if output.String() != "Please provide a natural language input.\n" {
		t.Errorf("Expected error message for missing input, got: %s", output.String())
	}

	// Reset output buffer
	output.Reset()

	// Test case: Valid argument provided
	rootCmd.SetArgs([]string{"suggest", "Show repository details"})
	err = rootCmd.Execute()
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if output.String() != "Processing input: Show repository details\nSuggested GitHub CLI command: gh repo view\n" {
		t.Errorf("Expected suggestion, got: %s", output.String())
	}
}
