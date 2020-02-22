package goarchive

import "testing"

func fileFilterFunc(filePath string) bool {
	return true
}

func TestValidateOptions(t *testing.T) {

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
