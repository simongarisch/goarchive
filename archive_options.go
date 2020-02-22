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

func validateOptions(options ArchiveOptions) error {
	stat, err := os.Stat(options.sourceFolderPath)
	if err != nil || !stat.IsDir() {
		return fmt.Errorf("source folder path doesn't exist")
	}
	return err
}
