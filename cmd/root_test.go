package cmd

import (
	"bytes"
	"testing"
)

func TestExecute(t *testing.T) {
	var output bytes.Buffer
	rootCmd.SetOut(&output)

	// Test case: Execute root command without arguments
	err := rootCmd.Execute()
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if !bytes.Contains(output.Bytes(), []byte("Usage:")) {
		t.Errorf("Expected usage information, got: %s", output.String())
	}
}
