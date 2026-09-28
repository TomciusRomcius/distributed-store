package debug

import (
	"net/http"

	"github.com/tomciusromcius/distributed-store/internal/query"
)

func DebugController(keyvalStore *query.KeyvalStore) func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		dataJson, err := keyvalStore.DumpDataJson()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(dataJson))
	}
}
