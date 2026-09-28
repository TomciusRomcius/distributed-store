package query

import (
	"encoding/json"
	"errors"
	"fmt"
	"sync"

	"github.com/tomciusromcius/distributed-store/internal/query/types"
)

type KeyvalStore struct {
	fileDataStorage FileDataStorage
	memtable        Memtable
	queryParser     QueryParser
	mutex           sync.Mutex
}

func NewKeyvalStore(fileDataStorage *FileDataStorage, memtable *Memtable, queryParser *QueryParser) *KeyvalStore {
	return &KeyvalStore{
		memtable:        *memtable,
		fileDataStorage: *fileDataStorage,
		queryParser:     *queryParser,
	}
}

func (r *KeyvalStore) InitFromLog() {
	channel := make(chan types.LogEntry)
	go r.fileDataStorage.RetrieveEntries(channel)
	r.memtable.PopulateFromChannel(channel)
	fmt.Println("Finished initializing from disk")
}

func (r *KeyvalStore) ExecuteQuery(query string) (string, error) {
	queryCmd, err := r.queryParser.ParseQuery(query)
	if err != nil {
		return "", err
	}

	switch queryCmd.Operation {
	case types.QueryOperationSet:
		r.set(queryCmd.Key, queryCmd.Val)
	case types.QueryOperationDel:
		r.del(queryCmd.Key)
	case types.QueryOperationGet:
		return r.get(queryCmd.Key)
	default:
		return "", errors.New("Invalid operation")
	}
	return "", nil
}

func (r *KeyvalStore) DumpDataJson() (string, error) {
	channel := make(chan *types.Pair)
	go r.memtable.DumpToStream(channel)
	pairs := []*types.Pair{}
	for pair := range channel {
		pairs = append(pairs, pair)
	}

	jsonBytes, err := json.Marshal(pairs)
	if err != nil {
		return "", err
	}
	jsonStr := string(jsonBytes)
	return jsonStr, nil
}

func (r *KeyvalStore) get(key string) (string, error) {
	fmt.Printf("key %s", key)
	val, err := r.memtable.get(key)
	if err != nil {
		return "", err
	}
	return val, nil
}

func (r *KeyvalStore) set(key string, value string) {
	fmt.Printf("key %s val %s", key, value)
	r.mutex.Lock()
	r.memtable.set(key, value)
	r.mutex.Unlock()
}

func (r *KeyvalStore) del(key string) {
	r.mutex.Lock()
	r.memtable.del(key)
	r.mutex.Unlock()
}
