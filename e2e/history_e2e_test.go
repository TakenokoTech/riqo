package e2e

import (
	"encoding/json"
	"os"
	"os/exec"
	"testing"
)

func TestSuggestAndExportE2E(t *testing.T) {
	histFile := "gh_history.txt"
	exportFile := "history_export.json"
	os.Remove(histFile)
	os.Remove(exportFile)

	cmd1 := exec.Command("go", "run", "../main.go", "suggest", "e2e test 1")
	out1, err1 := cmd1.CombinedOutput()
	if err1 != nil {
		t.Fatalf("failed to run suggest 1: %v\nstdout/stderr:\n%s", err1, string(out1))
	}
	cmd2 := exec.Command("go", "run", "../main.go", "suggest", "e2e test 2")
	out2, err2 := cmd2.CombinedOutput()
	if err2 != nil {
		t.Fatalf("failed to run suggest 2: %v\nstdout/stderr:\n%s", err2, string(out2))
	}
	cmd3 := exec.Command("go", "run", "../main.go", "history", "export")
	out3, err3 := cmd3.CombinedOutput()
	if err3 != nil {
		t.Fatalf("failed to run history export: %v\nstdout/stderr:\n%s", err3, string(out3))
	}

	data, err := os.ReadFile(exportFile)
	if err != nil {
		t.Fatalf("failed to read export file: %v", err)
	}
	var arr []map[string]string
	if err := json.Unmarshal(data, &arr); err != nil {
		t.Fatalf("failed to unmarshal export json: %v", err)
	}
	if len(arr) != 2 {
		t.Fatalf("expected 2 entries, got %d", len(arr))
	}
	if arr[0]["command"] != "suggest e2e test 1" || arr[1]["command"] != "suggest e2e test 2" {
		t.Fatalf("unexpected commands: %+v", arr)
	}
}
