package goarchive

import (
	"fmt"
	"os"
)

// Archive settings when running the archive job.
type Archive struct {
	sourceFolderPath  string
	archiveFolderName string
	fileFilterFunc    func(string) bool
}

func (archive *Archive) validate() error {
	exists := folderExists(archive.sourceFolderPath)
	if !exists {
		return fmt.Errorf("source folder path doesn't exist")
	}
	return nil
}

// Run our archive job
func (archive *Archive) Run() error {
	return nil
}

func fileExists(filePath string) bool {
	stat, err := os.Stat(filePath)
	if err == nil && !stat.IsDir() {
		return true
	}
	return false
}

func folderExists(folderPath string) bool {
	stat, err := os.Stat(folderPath)
	if err == nil && stat.IsDir() {
		return true
	}
	return false
}
