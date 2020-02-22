# goarchive

## Installation
```bash
go get github.com/simongarisch/goarchive
```

## Usage
```go
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

func main(){
	archive := Archive{
		SourceFolderPath:  "test_folder",
		ArchiveFolderName: "archive",
		FileFilterFunc:    archiveCsv,
    }
    
    archive.Run()
}
```