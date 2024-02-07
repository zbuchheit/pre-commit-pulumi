package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/pulumi/pulumi/sdk/v3/go/common/apitype"
)

func main() {
	args := os.Args[1:]
	for _, filePath := range args {
		if isPulumiStateFile(filePath) {
			fmt.Printf("Error: Attempt to commit Pulumi state file detected: %s\n", filePath)
			os.Exit(1)
		}
	}
}

func isPulumiStateFile(filePath string) bool {
	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error reading file: %s\n", err)
		return false
	}

	var untypedDeployment apitype.UntypedDeployment
	if err := json.Unmarshal(data, &untypedDeployment); err != nil {
		return heuristicCheckForPulumiStateFile(data)
	}
	switch untypedDeployment.Version {
	case 1:
		return validateDeploymentV1(untypedDeployment.Deployment)
	case 2:
		return validateDeploymentV2(untypedDeployment.Deployment)
	case 3:
		return validateDeploymentV3(untypedDeployment.Deployment)
	default:
		return heuristicCheckForPulumiStateFile(untypedDeployment.Deployment)
	}
}

func validateDeploymentV1(deployment json.RawMessage) bool {
	var deploymentV1 apitype.DeploymentV1
	if err := json.Unmarshal(deployment, &deploymentV1); err != nil {
		fmt.Fprintf(os.Stderr, "Error validating deployment version 1: %v\n", err)
		return false
	}
	if deploymentV1.Manifest.Time.IsZero() {
		fmt.Fprintf(os.Stderr, "Error validating deployment version 1: Manifest time is zero\n")
		return false
	}
	return true
}

func validateDeploymentV2(deployment json.RawMessage) bool {
	var deploymentV2 apitype.DeploymentV2
	if err := json.Unmarshal(deployment, &deploymentV2); err != nil {
		fmt.Fprintf(os.Stderr, "Error validating deployment version 2: %v\n", err)
		return false
	}
	if deploymentV2.Manifest.Time.IsZero() {
		fmt.Fprintf(os.Stderr, "Error validating deployment version 2: Manifest time is zero\n")
		return false
	}
	return true
}

func validateDeploymentV3(deployment json.RawMessage) bool {
	var deploymentV3 apitype.DeploymentV3
	if err := json.Unmarshal(deployment, &deploymentV3); err != nil {
		fmt.Fprintf(os.Stderr, "Error validating deployment version 3: %v\n", err)
		return false
	}
	if deploymentV3.Manifest.Time.IsZero() {
		fmt.Fprintf(os.Stderr, "Error validating deployment version 3: Manifest time is zero\n")
		return false
	}
	return true
}

func heuristicCheckForPulumiStateFile(fileData []byte) bool {
	var jsonData map[string]interface{}
	if err := json.Unmarshal(fileData, &jsonData); err != nil {
		return false
	}

	// Check for keys that are expected to exist in Pulumi state files
	if _, versionExists := jsonData["version"]; versionExists {
		if deploymentData, deploymentExists := jsonData["deployment"]; deploymentExists {
			if deploymentMap, ok := deploymentData.(map[string]interface{}); ok {
				if _, manifestExists := deploymentMap["manifest"]; manifestExists {
					return true
				}
			}
		}
	}
	return false
}
