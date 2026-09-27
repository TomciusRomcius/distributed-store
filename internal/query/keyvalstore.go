package query

import (
	"fmt"
	"sync"

	"github.com/tomciusromcius/distributed-store/internal/query/types"
)

type KeyvalStore struct {
	fileDataStorage FileDataStorage
	memtable        Memtable
	mutex           sync.Mutex
}

func NewKeyvalStore(fileDataStorage *FileDataStorage, memtable *Memtable) *KeyvalStore {
	return &KeyvalStore{
		memtable:        *memtable,
		fileDataStorage: *fileDataStorage,
	}
}

func (r *KeyvalStore) InitFromLog() {
	channel := make(chan types.LogEntry)
	go r.fileDataStorage.RetrieveEntries(channel)
	r.memtable.PopulateFromChannel(channel)
	fmt.Println("Finished initializing from disk")
}

func (r *KeyvalStore) get(key string) (string, error) {
	val, err := r.memtable.get(key)
	if err != nil {
		return "", err
	}
	return val, nil
}

func (r *KeyvalStore) set(key string, value string) {
	r.mutex.Lock()
	r.memtable.set(key, value)
	r.mutex.Unlock()
}

func (r *KeyvalStore) del(key string) {
	r.mutex.Lock()
	r.memtable.del(key)
	r.mutex.Unlock()
}
