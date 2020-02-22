package goarchive

import (
	"os"
	"testing"
)

func fileFilterFunc(filePath string) bool {
	return true
}

func TestSourceFolderDoesntExist(t *testing.T) {

	options := ArchiveOptions{
		sourceFolderPath:  "not_a_folder",
		archiveFolderName: "archive",
		fileFilterFunc:    fileFilterFunc,
	}

	err := validateOptions(options)
	actual := err.Error()
	expected := "source folder path doesn't exist"
	if actual != expected {
		t.Errorf("Error actual = %v, and Expected = %v.", actual, expected)
	}

}

func TestSourceFolderExists(t *testing.T) {
	sourceFolder := "test_folder"

	exists, err := folderExists(sourceFolder)

	if exists {
		err := os.RemoveAll(sourceFolder)
		if err != nil {
			t.Error("Unable to remove test folder")
		}
	}

	err = os.Mkdir(sourceFolder, os.ModePerm)
	if err != nil {
		t.Error("Unable to create test folder")
	}

	options := ArchiveOptions{
		sourceFolderPath:  sourceFolder,
		archiveFolderName: "archive",
		fileFilterFunc:    fileFilterFunc,
	}

	err = validateOptions(options)
	if err != nil {
		t.Error(err.Error())
	}

	os.RemoveAll(sourceFolder)
}
