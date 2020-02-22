package goarchive

import (
	"errors"
	"io/ioutil"
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

	err := os.MkdirAll(sourceFolderPath, os.ModePerm)
	if err != nil {
		return err
	}

	for _, fileName := range files {
		filePath := path.Join(sourceFolderPath, fileName)
		err := ioutil.WriteFile(filePath, []byte("Hello"), 0755)
		if err != nil {
			return err
		}
	}
	return nil
}

func removeTestData() error {
	exists := folderExists(sourceFolderPath)

	if exists {
		err := os.RemoveAll(sourceFolderPath)
		if err != nil {
			return errors.New("Unable to remove test folder")
		}
	}
	return nil
}

func resetTestData() error {
	err := removeTestData()
	if err != nil {
		return err
	}

	err = createTestData()
	if err != nil {
		return err
	}

	return nil
}

func getArchiveMap() (map[string]bool, error) {
	archiveMap := make(map[string]bool)
	for _, fileName := range files {
		filePath := path.Join(sourceFolderPath, fileName)
		archivePath := path.Join(sourceFolderPath, archiveFolderName, fileName)

		inSourceFolder := fileExists(filePath)
		inArchiveFolder := fileExists(archivePath)
		if inSourceFolder && inArchiveFolder {
			return archiveMap, errors.New("File is both in source and archive folder")
		}
		if inSourceFolder {
			archiveMap[fileName] = false
		} else {
			archiveMap[fileName] = true
		}
	}

	return archiveMap, nil
}

func TestArchiveNone(t *testing.T) {
	err := resetTestData()
	if err != nil {
		t.Error(err)
	}

	archive := Archive{
		sourceFolderPath:  sourceFolderPath,
		archiveFolderName: archiveFolderName,
		fileFilterFunc:    archiveNone,
	}

	err = archive.Run()
	if err != nil {
		t.Error(err)
	}

	archiveMap, err := getArchiveMap()
	if err != nil {
		t.Error(err)
	}

	// none should be archived, all false
	for _, isArchived := range archiveMap {
		if isArchived == true {
			t.Error("No files should be archived")
		}
	}

	removeTestData()
}

func TestArchiveAll(t *testing.T) {
	err := resetTestData()
	if err != nil {
		t.Error(err)
	}

	archive := Archive{
		sourceFolderPath:  sourceFolderPath,
		archiveFolderName: archiveFolderName,
		fileFilterFunc:    archiveAll,
	}

	err = archive.Run()
	if err != nil {
		t.Error(err)
	}

	archiveMap, err := getArchiveMap()
	if err != nil {
		t.Error(err)
	}

	// all should be archived
	for fileName, isArchived := range archiveMap {
		if isArchived == false {
			t.Errorf("%s - All files should be archived", fileName)
		}
	}

	removeTestData()
}

func TestArchiveCsv(t *testing.T) {
	err := resetTestData()
	if err != nil {
		t.Error(err)
	}

	archive := Archive{
		sourceFolderPath:  sourceFolderPath,
		archiveFolderName: archiveFolderName,
		fileFilterFunc:    archiveCsv,
	}

	err = archive.Run()
	if err != nil {
		t.Error(err)
	}

	archiveMap, err := getArchiveMap()
	if err != nil {
		t.Error(err)
	}

	// all should be archived
	for fileName, isArchived := range archiveMap {
		if isArchived == false && strings.HasSuffix(fileName, ".csv") {
			t.Errorf("%s - All csv files should be archived", fileName)
		}
	}

	removeTestData()
}
