package utils

import (
	"pre-commit-pulumi/internal/testutils"
	"testing"
)

func TestDirectoryScanner_Scan(t *testing.T) {

	files := map[string]string{
		"file1.json": `{"version":3,"deployment":{"manifest":{"time":"2020-06-01T14:00:00Z"}}}`,
		"file2.txt":  "This is not a Pulumi state file",
		"file3.json": `{"version":3,"deployment":{"manifest":{"time":"2020-06-01T14:00:00Z"}}}`,
		"file4.yaml": "This is not a Pulumi state file",
	}

	dir, cleanup := testutils.CreateTempDirWithFiles(t, files)
	defer cleanup()

	evaluator := &testutils.MockFileEvaluator{ShouldReturnPulumiStateFile: true}

	ds := NewDirectoryScanner(evaluator)

	detectedFiles, err := ds.Scan(dir)
	if err != nil {
		t.Fatalf("Expected no error, but got %v", err)
	}

	expectedFiles := []string{"file1.json", "file3.json"}
	if len(detectedFiles) != len(expectedFiles) {
		t.Errorf("Expected %d files to be detected, got %d", len(expectedFiles), len(detectedFiles))
	}

}
