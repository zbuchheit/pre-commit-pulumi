package utils

import (
	"path/filepath"
	"pre-commit-pulumi/internal/testutils"
	"testing"
)

func TestIsPulumiStateFileWithTestData(t *testing.T) {
	testDataDir := testutils.TestDataPath()

	tests := []struct {
		name        string
		filePath    string
		want        bool
		expectError bool
	}{
		{
			name:        "Valid Pulumi V1 State",
			filePath:    filepath.Join(testDataDir, "/valid/deployment_v1.json"),
			want:        true,
			expectError: false,
		},
		{
			name:        "Valid Pulumi V2 State",
			filePath:    filepath.Join(testDataDir, "/valid/deployment_v2.json"),
			want:        true,
			expectError: false,
		},
		{
			name:        "Valid Pulumi V3 State",
			filePath:    filepath.Join(testDataDir, "/valid/deployment_v3.json"),
			want:        true,
			expectError: false,
		},
		// {
		// 	name:        "Malformed Pulumi State",
		// 	filePath:    filepath.Join(testDataDir, "/valid/malformed_state_file.json"),
		// 	want:        true,
		// 	expectError: false,
		// },
		{
			name:        "Version and Deployment file",
			filePath:    filepath.Join(testDataDir, "/valid/version.json"),
			want:        true,
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			evaluator := NewPulumiFileEvaluator()
			got, err := evaluator.IsPulumiStateFile(tt.filePath)

			if tt.expectError && err == nil {
				t.Errorf("%s: expected an error but got none", tt.name)
			} else if !tt.expectError && err != nil {
				t.Errorf("%s: did not expect an error but got one: %v", tt.name, err)
			}

			if got != tt.want {
				t.Errorf("%s: IsPulumiStateFile(%s) = %t, want %t", tt.name, tt.filePath, got, tt.want)
			}
		})
	}
}
