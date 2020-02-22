package goarchive

import (
	"os"
	"testing"
)

func fileFilterFunc(filePath string) bool {
	return true
}

func TestSourceFolderDoesntExist(t *testing.T) {

	archive := Archive{
		sourceFolderPath:  "not_a_folder",
		archiveFolderName: "archive",
		fileFilterFunc:    fileFilterFunc,
	}

	err := archive.validate()
	actual := err.Error()
	expected := "source folder path doesn't exist"
	if actual != expected {
		t.Errorf("Error actual = %v, and Expected = %v.", actual, expected)
	}

}

func TestSourceFolderExists(t *testing.T) {
	sourceFolder := "test_folder"

	exists := folderExists(sourceFolder)

	if exists {
		err := os.RemoveAll(sourceFolder)
		if err != nil {
			t.Error("Unable to remove test folder")
		}
	}

	err := os.Mkdir(sourceFolder, os.ModePerm)
	if err != nil {
		t.Error("Unable to create test folder")
	}

	archive := Archive{
		sourceFolderPath:  sourceFolder,
		archiveFolderName: "archive",
		fileFilterFunc:    fileFilterFunc,
	}

	err = archive.validate()
	if err != nil {
		t.Error(err)
	}

	os.RemoveAll(sourceFolder)
}
