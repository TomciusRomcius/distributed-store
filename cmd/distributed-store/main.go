package main

import (
	"net/http"

	"github.com/tomciusromcius/distributed-store/internal/debug"
	"github.com/tomciusromcius/distributed-store/internal/query"
	"github.com/tomciusromcius/distributed-store/internal/utils"
)

func main() {
	fileDataStorage := query.NewFileDataStorage()
	memtable := query.NewMemtable()
	queryParser := &query.QueryParser{}

	keyValStore := query.NewKeyvalStore(fileDataStorage, memtable, queryParser)
	keyValStore.InitFromLog()
	http.HandleFunc("/query", query.QueryController(keyValStore))
	if utils.IsDebug() {
		handler := debug.DebugController(keyValStore)
		http.HandleFunc("/debug", handler)
	}
	http.ListenAndServe(":8080", nil)
}
