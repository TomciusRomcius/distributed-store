package query

type FileDataStorage struct {
	logPath string
}

func NewFileDataStorage() *FileDataStorage {
	return &FileDataStorage{}
}
