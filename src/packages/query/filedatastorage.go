package query

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/tomciusromcius/distributed-store/src/packages/query/types"
	"github.com/tomciusromcius/distributed-store/src/packages/utils"
)

type FileDataStorage struct {
	logPath string
}

func NewFileDataStorage() *FileDataStorage {
	return &FileDataStorage{}
}

func (r *FileDataStorage) RetrieveEntries(entryChannel chan types.LogEntry) {
	defer close(entryChannel)
	entries, err := os.ReadDir("./data")
	if err != nil {
		log.Fatal(err)
		return
	}

	for _, e := range entries {
		if e.IsDir() {
			continue
		}
		file, err := os.OpenFile("./data/"+e.Name(), os.O_RDONLY, 0)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(e.Name())
		r.getEntryFromFile(file, entryChannel)
	}
}

func (r *FileDataStorage) getEntryFromFile(file *os.File, entryChannel chan types.LogEntry) error {
	fileInfo, err := file.Stat()
	if err != nil {
		return err
	}

	fileBytes := make([]byte, fileInfo.Size())
	_, err = file.Read(fileBytes)
	if err != nil {
		return err
	}

	fileDataStr := string(fileBytes)
	lines := strings.Split(fileDataStr, "\n")
	for _, line := range lines {
		strLen := len(line)
		fmt.Println(line)
		utils.Assert(strLen != 0, "Line length is 0")
		if line[0] == '?' {
			utils.Assert(strLen != 1, "Line length is 1")
			key := line[1:strLen]
			entryChannel <- *types.NewRemoveLogEntry(key)
			continue
		}
		lineParts := strings.SplitN(line, "=", 2)
		key := lineParts[0]
		value := lineParts[1]
		entryChannel <- *types.NewAddLogEntry(key, value)
	}

	return nil
}
