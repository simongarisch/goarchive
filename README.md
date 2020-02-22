[![Build Status](https://travis-ci.com/simongarisch/goarchive.svg?branch=master)](https://travis-ci.com/simongarisch/goarchive)

# goarchive

## Installation
```bash
go get github.com/simongarisch/goarchive
```

## Usage
```go
package main

import (
	"strings"

	"github.com/simongarisch/goarchive"
)

func archiveCsv(filePath string) bool {
	if strings.HasSuffix(filePath, ".csv") {
		return true
	}
	return false
}

func main() {
	archive := goarchive.Archive{
		SourceFolderPath:  "test_folder",
		ArchiveFolderName: "archive",
		FileFilterFunc:    archiveCsv,
	}

	archive.Run()
}
```