package goarchive

import (
	"fmt"
	"os"
)

// ArchiveOptions when running the archive job.
type ArchiveOptions struct {
	sourceFolderPath  string
	archiveFolderName string
	fileFilterFunc    func(string) bool
}

func folderExists(folderPath string) (bool, error) {
	stat, err := os.Stat(folderPath)
	if err == nil && stat.IsDir() {
		return true, err
	}
	return false, err
}

func validateOptions(options ArchiveOptions) error {
	exists, err := folderExists(options.sourceFolderPath)
	if err != nil || !exists {
		return fmt.Errorf("source folder path doesn't exist")
	}
	return err
}
