package query

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/tomciusromcius/distributed-store/internal/query/types"
	"github.com/tomciusromcius/distributed-store/internal/utils"
)

type FileDataStorage struct {
	logPath            string // Directory of log files
	headOperationCount int
	operationsPerFile  int
	logHead            string // Log head file name without an extension
}

func NewFileDataStorage() *FileDataStorage {
	return &FileDataStorage{
		logPath:           "./data",
		operationsPerFile: 2,
	}
}

func (r *FileDataStorage) AppendToLog(entry *types.LogEntry) error {
	if r.logHead == "" || r.headOperationCount >= r.operationsPerFile {
		if err := r.CreateNewHead(); err != nil {
			log.Printf("create log head: %v", err)
			return err
		}
	}

	file, err := os.OpenFile(r.logHeadPath(), os.O_WRONLY|os.O_APPEND, 0)
	if err != nil {
		fmt.Println("hit")
		return err
	}

	entryText := r.logEntryToLogString(entry)
	_, err = file.WriteString(entryText)
	if err != nil {

		fmt.Println("hit")
		return err
	}

	r.headOperationCount++
	return nil
}

func (r *FileDataStorage) CreateNewHead() error {
	timestamp := time.Now().UTC().UnixNano()
	if r.logPath == "" {
		r.logPath = "./data"
	}

	_, err := os.ReadDir(r.logPath)
	if err != nil {
		if err := os.Mkdir(r.logPath, 0755); err != nil {
			return err
		}
	}
	r.logHead = strconv.FormatUint(uint64(timestamp), 10)
	file, err := os.Create(r.logHeadPath())
	if err != nil {
		return err
	}
	r.headOperationCount = 0
	return file.Close()
}

func (r *FileDataStorage) logEntryToLogString(entry *types.LogEntry) string {
	var entryText string
	switch entry.Operation {
	case types.OperationSet:
		entryText = fmt.Sprintf("%s=%s\n", entry.Key, entry.Val)
	case types.OperationRemove:
		entryText = fmt.Sprintf("?%s\n", entry.Key)
	default:
		log.Fatalf("invalid operation type: %v", entry.Operation)
	}
	return entryText
}

func (r *FileDataStorage) logHeadPath() string {
	return fmt.Sprintf("%s/%s.data", r.logPath, r.logHead)
}

func (r *FileDataStorage) formatLogPath(timestamp int64) string {
	return fmt.Sprintf("%s/%d.data", strings.TrimRight(r.logPath, "/"), timestamp)
}

func (r *FileDataStorage) RetrieveEntries(entryChannel chan types.LogEntry) {
	defer close(entryChannel)
	entries, err := os.ReadDir(r.logPath)
	if err != nil {
		return
	}

	for _, e := range entries {
		if e.IsDir() {
			continue
		}
		r.headOperationCount = 0
		fmt.Println(r.logHead + e.Name())
		file, err := os.OpenFile(r.logPath+"/"+e.Name(), os.O_RDONLY, 0)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(e.Name())
		r.getEntryFromFile(file, entryChannel)
		r.logHead = strings.TrimSuffix(e.Name(), filepath.Ext(e.Name()))
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
		if strLen == 0 {
			continue
		}
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
		r.headOperationCount++
	}

	return nil
}
