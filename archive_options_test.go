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
		SourceFolderPath:  "not_a_folder",
		ArchiveFolderName: "archive",
		FileFilterFunc:    fileFilterFunc,
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
		SourceFolderPath:  sourceFolder,
		ArchiveFolderName: "archive",
		FileFilterFunc:    fileFilterFunc,
	}

	err = archive.validate()
	if err != nil {
		t.Error(err)
	}

	os.RemoveAll(sourceFolder)
}
