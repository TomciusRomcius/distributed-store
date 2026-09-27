package main

import (
	"net/http"

	"github.com/tomciusromcius/distributed-store/internal/query"
)

func main() {
	fileDataStorage := query.NewFileDataStorage()
	memtable := query.NewMemtable()
	queryParser := &query.QueryParser{}

	keyValStore := query.NewKeyvalStore(fileDataStorage, memtable, queryParser)
	keyValStore.InitFromLog()
	http.HandleFunc("/query", query.QueryController(keyValStore))
	http.ListenAndServe(":8080", nil)
}
