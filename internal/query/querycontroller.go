package query

import (
	"encoding/json"
	"net/http"
)

type QueryRequest struct {
	Query string `json:"query"`
}

func QueryController(store *KeyvalStore) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		queryRequest := &QueryRequest{}
		err := json.NewDecoder(req.Body).Decode(queryRequest)
		if err != nil {
			http.Error(w, "invalid params", http.StatusBadRequest)
			return
		}

		result, err := store.ExecuteQuery(queryRequest.Query)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write([]byte(result))
	}
}
