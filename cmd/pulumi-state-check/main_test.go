// main_test.go
package main

import (
	"bytes"
	"os"
	"testing"
)

func TestRunWithAllFlag(t *testing.T) {
	oldStdout := os.Stdout
	defer func() { os.Stdout = oldStdout }()
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("Failed to create pipe: %v", err)
	}
	os.Stdout = w

	exitCode := run([]string{"--all"})

	w.Close()
	var buf bytes.Buffer
	os.Stdout = oldStdout
	_, err = buf.ReadFrom(r)
	if err != nil {
		t.Fatalf("Failed to read stdout: %v", err)
	}

	output := buf.String()
	expectedOutput := "Recursively scanning the current directory."

	if !bytes.Contains(buf.Bytes(), []byte(expectedOutput)) {
		t.Errorf("Expected stdout to contain %q, got %q", expectedOutput, output)
	}

	if exitCode != 0 {
		t.Errorf("Expected exit code 0, got %d", exitCode)
	}
}
