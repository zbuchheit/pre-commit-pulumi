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

func parseFlags(args []string) (Config, error) {
	var config Config
	flagSet := flag.NewFlagSet("pulumi-state-check", flag.ContinueOnError)
	flagSet.BoolVar(&config.All, "all", false, "Recursively scan the current directory for Pulumi state files if no specific file paths are provided.")
	flagSet.BoolVar(&config.All, "a", false, "Alias for --all")
	if err := flagSet.Parse(args); err != nil {
		return config, err
	}

	config.FilePaths = flagSet.Args()
	return config, nil
}

func run(args []string) int {
	config, err := parseFlags(args)
	if err != nil {
		fmt.Println(err)
		return 1
	}

	if !config.All && len(config.FilePaths) == 0 {
		fmt.Println("No file paths provided.")
		flag.PrintDefaults()
		return 0
	}

	scanner := utils.NewDirectoryScanner(utils.NewPulumiFileEvaluator())

	pathsToScan := determinePathsToScan(config)

	detectedFiles, scanErrors := scanPaths(scanner, pathsToScan)

	for _, file := range detectedFiles {
		fmt.Printf("Detected Pulumi state file: %s\n", file)
	}
	for _, err := range scanErrors {
		fmt.Printf("Error: %v\n", err)
	}

	if len(detectedFiles) > 0 || len(scanErrors) > 0 {
		return 1
	} else {
		fmt.Println("No Pulumi state files detected.")
		return 0
	}
}

func determinePathsToScan(config Config) []string {
	if len(config.FilePaths) > 0 {
		fmt.Println("Scanning specific file paths.")
		return config.FilePaths
	}
	fmt.Println("Recursively scanning the current directory.")
	return []string{"."}
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

func main() {
	os.Exit(run(os.Args[1:]))
}
