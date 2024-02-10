package testutils

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

func CreateTestFiles(t *testing.T, dir string, files map[string]string) {
	for filename, content := range files {
		path := filepath.Join(dir, filename)
		if err := os.WriteFile(path, []byte(content), 0644); err != nil {
			t.Fatalf("Failed to create file %q: %v", filename, err)
		}
	}
}

func CreateTempDirWithFiles(t *testing.T, files map[string]string) (string, func()) {
	tempDir, err := os.MkdirTemp("", "test-directory")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}

	CreateTestFiles(t, tempDir, files)

	return tempDir, func() { os.RemoveAll(tempDir) }
}

func TestDataPath() string {
    _, filename, _, _ := runtime.Caller(0)
    dir := filepath.Dir(filename)
    return filepath.Join(dir, "..", "..", "testdata")
}
