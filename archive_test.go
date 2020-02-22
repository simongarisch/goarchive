package goarchive

import (
	"errors"
	"os"
	"path"
	"strings"
	"testing"
)

const (
	sourceFolderPath  = "test_folder"
	archiveFolderName = "archive"
)

func archiveAll(filePath string) bool {
	return true
}

func archiveNone(filePath string) bool {
	return false
}

func archiveCsv(filePath string) bool {
	if strings.HasSuffix(filePath, ".csv") {
		return true
	}
	return false
}

var files [5]string

func createTestData() error {
	files = [...]string{
		"csv1.csv",
		"csv2.csv",
		"csv3.csv",
		"txt1.txt",
		"txt2.txt",
	}
	for _, fileName := range files {
		filePath := path.Join(sourceFolderPath, fileName)
		err := os.MkdirAll(filePath, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}

func removeTestData() error {
	exists, err := folderExists(sourceFolderPath)

	if exists {
		err := os.RemoveAll(sourceFolderPath)
		if err != nil {
			return errors.New("Unable to remove test folder")
		}
	}
	return nil
}

func isArchived() (map[string]bool, error) {
	for _, filePath := range files {

	}
}

func TestArchiveJob(t *testing.T) {
	// create our archive options
	options := ArchiveOptions{
		sourceFolderPath:  sourceFolderPath,
		archiveFolderName: archiveFolderName,
		fileFilterFunc:    fileFilterFunc,
	}
	err := validateOptions(options)
	if err != nil {
		t.Error(err)
	}

	// remove our test folder (if exists)
	err = removeTestData()
	if err != nil {
		t.Error(err)
	}

	// create test files
	err = createTestData()
	if err != nil {
		t.Error(err)
	}

	// check archiving
	err := RunArchiveJob(options)

}
