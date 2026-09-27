package main

import (
	"net/http"

	"github.com/tomciusromcius/distributed-store/src/packages/query"
)

func main() {
	fileDataStorage := query.NewFileDataStorage()
	memtable := query.NewMemtable()

	keyValStore := query.NewKeyvalStore(fileDataStorage, memtable)
	keyValStore.InitFromLog()
	http.HandleFunc("/query", query.QueryController(keyValStore))
	http.ListenAndServe(":8080", nil)
}
