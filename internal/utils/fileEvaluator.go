package utils

import (
	"encoding/json"
	"fmt"
	"os"
)

type FileEvaluator interface {
	IsPulumiStateFile(filePath string) (bool, error)
}

func NewPulumiFileEvaluator() FileEvaluator {
	return &PulumiFileEvaluator{}
}

type PulumiFileEvaluator struct{}

func (p *PulumiFileEvaluator) IsPulumiStateFile(filePath string) (bool, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return false, fmt.Errorf("error reading file %s: %w", filePath, err)
	}

	var parsedData map[string]interface{}
	err = json.Unmarshal(data, &parsedData)
	if err != nil {
		return false, nil
	}
	requiredKeys := []string{"version", "deployment"}
	deploymentKeys := []string{"manifest", "secrets_providers", "resources"}
	score := 0

	for _, key := range requiredKeys {
		if _, exists := parsedData[key]; exists {
			score++
		}
	}

	if deployment, exists := parsedData["deployment"].(map[string]interface{}); exists {
		for _, key := range deploymentKeys {
			if _, exists := deployment[key]; exists {
				score++
			}
		}
	}

	//TODO: This is a placeholder for a more sophisticated scoring system and also account for the possibility of a malformed state file

	const threshold = 2
	return score >= threshold, nil
}
