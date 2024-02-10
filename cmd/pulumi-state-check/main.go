package main

import (
	"flag"
	"fmt"
	"os"
	"pre-commit-pulumi/internal/utils"
)

type Config struct {
	All       bool
	FilePaths []string
}

func parseFlags() Config {
	var config Config
	flag.BoolVar(&config.All, "all", false, "Recursively scan the current directory for Pulumi state files if no specific file paths are provided.")
	flag.BoolVar(&config.All, "a", false, "Alias for --all")
	flag.Parse()
	config.FilePaths = flag.Args()
	return config
}

func main() {
	config := parseFlags()

	if !config.All && len(config.FilePaths) == 0 {
		fmt.Println("No file paths provided.")
		flag.PrintDefaults()
		os.Exit(2)
	}

	scanner := utils.NewDirectoryScanner(utils.NewPulumiFileEvaluator())

	pathsToScan := []string{"."}

	if config.All {
		pathsToScan = config.FilePaths
	}

	detectedFiles, scanErrors := scanPaths(scanner, pathsToScan)

	for _, file := range detectedFiles {
		fmt.Printf("Detected Pulumi state file: %s\n", file)
	}
	for _, err := range scanErrors {
		fmt.Printf("Error: %v\n", err)
	}

	if len(detectedFiles) > 0 || len(scanErrors) > 0 {
		os.Exit(1)
	} else {
		fmt.Println("No Pulumi state files detected.")
		os.Exit(0)
	}
}

func scanPaths(scanner *utils.DirectoryScanner, paths []string) ([]string, []error) {
	var detectedFiles []string
	var scanErrors []error
	for _, filePath := range paths {
		files, err := scanner.Scan(filePath)
		if err != nil {
			scanErrors = append(scanErrors, err)
		}
		detectedFiles = append(detectedFiles, files...)
	}
	return detectedFiles, scanErrors
}
