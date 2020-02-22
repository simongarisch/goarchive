package goarchive

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
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
	err := archive.validate()
	if err != nil {
		return err
	}

	files, err := ioutil.ReadDir(archive.sourceFolderPath)
	if err != nil {
		return err
	}

	if len(files) == 0 {
		return nil
	}

	// create the archive folder if it doesn't exist
	archiveFolderPath := path.Join(archive.sourceFolderPath, archive.archiveFolderName)
	if !folderExists(archiveFolderPath) {
		err := os.MkdirAll(archiveFolderPath, os.ModePerm)
		if err != nil {
			return err
		}
	}

	// files contains os.FileInfo types, which provide the methods
	// https://golang.org/pkg/os/#FileInfo
	for _, f := range files {
		fileName := f.Name()
		mustArchiveFile := archive.fileFilterFunc(fileName)
		if mustArchiveFile {
			oldPath := path.Join(archive.sourceFolderPath, fileName)
			newPath := path.Join(archiveFolderPath, fileName)
			err := os.Rename(oldPath, newPath)
			if err != nil {
				return err
			}
		}
	}

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
