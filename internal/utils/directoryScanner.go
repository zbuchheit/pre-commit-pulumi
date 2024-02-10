package utils

import (
	"os"
	"path/filepath"
)

type DirectoryScanner struct {
	evaluator FileEvaluator
}

func NewDirectoryScanner(evaluator FileEvaluator) *DirectoryScanner {
	return &DirectoryScanner{evaluator: evaluator}
}

func (ds *DirectoryScanner) Scan(dir string) ([]string, error) {
	var filesFound []string
	err := filepath.Walk(dir, ds.makeWalkFunc(&filesFound))

	if err != nil {
		return nil, err
	}

	return filesFound, nil
}

func (ds *DirectoryScanner) makeWalkFunc(filesFound *[]string) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && filepath.Ext(path) == ".json" {
			if isPulumiStateFile, err := ds.evaluator.IsPulumiStateFile(path); err!= nil {
				return err
			} else if isPulumiStateFile {
				*filesFound = append(*filesFound, path)
			}
		}
		return nil
	}
}
