package main

import (
	"path/filepath"
	"testing"
)

func TestIsPulumiStateFile_InvalidTSConfig(t *testing.T) {
    filename := filepath.Join("testdata/invalid", "invalid_tsconfig.json")
	testIsPulumiStateFile(t, filename, false)
}
func TestIsPulumiStateFile_InvalidMalformedJSON(t *testing.T) {
    filename := filepath.Join("testdata/invalid", "invalid_malformed.json")
	testIsPulumiStateFile(t, filename, false)

}
func TestIsPulumiStateFile_InvalidEmpty(t *testing.T) {
    filename := filepath.Join("testdata/invalid", "invalid_empty.json")
	testIsPulumiStateFile(t, filename, false)
}
func TestIsPulumiStateFile_InvalidNotJSON(t *testing.T) {
    filename := filepath.Join("testdata/invalid", "invalid_not_json.text")
	testIsPulumiStateFile(t, filename, false)
}

func TestIsPulumiStateFileV1(t *testing.T) {
	filePath := filepath.Join("testdata/valid", "deployment_v1.json")
	testIsPulumiStateFile(t, filePath, true) // true indicates it is expected to be a valid state file
}

func TestIsPulumiStateFileV2(t *testing.T) {
	filePath := filepath.Join("testdata/valid", "deployment_v2.json")
	testIsPulumiStateFile(t, filePath, true)
}

func TestIsPulumiStateFileV3(t *testing.T) {
	filePath := filepath.Join("testdata/valid", "deployment_v3.json")
	testIsPulumiStateFile(t, filePath, true)
}

// testIsPulumiStateFile is a helper function to test isPulumiStateFile with different state files.
func testIsPulumiStateFile(t *testing.T, filePath string, expected bool) {
	result := isPulumiStateFile(filePath)
	if result != expected {
		t.Errorf("Expected %v, got %v for file %s", expected, result, filePath)
	}
}
