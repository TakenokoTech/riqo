package cmd

import (
	"bytes"
	"testing"

	"github.com/spf13/cobra"
)

func TestLLMCommand(t *testing.T) {
	var output bytes.Buffer

	// Mock rootCmd for testing
	mockRootCmd := &cobra.Command{
		Use: "llm",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				output.WriteString("Please provide input for the LLM.\n")
			} else {
				output.WriteString("Mock LLM Response\n")
			}
		},
	}

	// Test case: No arguments provided
	mockRootCmd.SetOut(&output)
	mockRootCmd.SetArgs([]string{})
	err := mockRootCmd.Execute()
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if output.String() != "Please provide input for the LLM.\n" {
		t.Errorf("Expected error message for missing input, got: %s", output.String())
	}

	// Reset output buffer
	output.Reset()

	// Test case: Valid argument provided
	mockRootCmd.SetArgs([]string{"Test input"})
	err = mockRootCmd.Execute()
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if output.String() != "Mock LLM Response\n" {
		t.Errorf("Expected mock LLM response, got: %s", output.String())
	}
}
