package cmd

import (
	"bytes"
	"strings"
	"testing"
)

func TestRealtimeCommand(t *testing.T) {
	var output bytes.Buffer
	mockInput := "test input\nexit\n"
	mockReader := strings.NewReader(mockInput)

	realtimeCmd.SetOut(&output)
	realtimeCmd.SetIn(mockReader)

	err := realtimeCmd.Execute()
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
}
