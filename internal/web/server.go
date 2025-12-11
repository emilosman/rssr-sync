package web

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/emilosman/rssr-sync/internal/data"
)

type dataRequest struct {
	ApiKey string `json:"apiKey"`
	Ts     int64  `json:"ts"`
}

type dataResponse struct {
	Data string `json:"data"`
}

func Server(lists *data.Lists) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/data", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var req dataRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "invalid JSON body", http.StatusBadRequest)
			return
		}

		if req.ApiKey == "" {
			http.Error(w, "apiKey is required", http.StatusBadRequest)
			return
		}

		result, err := lists.GetList(req.ApiKey, req.Ts)
		if err != nil {
			switch {
			case errors.Is(err, data.ErrOldTimestamp):
				http.Error(w, err.Error(), http.StatusConflict)
			default:
				http.Error(w, err.Error(), http.StatusNotFound)
			}
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(dataResponse{Data: string(result)})
	})

	return mux
}
