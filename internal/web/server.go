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
		req, err := validateRequest(w, r)
		if err != nil {
			return
		}

		result, err := lists.GetList(req.ApiKey, req.Ts)
		if err != nil {
			switch {
			case errors.Is(err, data.ErrOldTimestamp):
				http.Error(w, err.Error(), http.StatusConflict)
			case errors.Is(err, data.ErrdClientOldTimestampUpdate):
				w.WriteHeader(http.StatusOK)
				w.Write(result)
			default:
				http.Error(w, err.Error(), http.StatusNotFound)
			}
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(dataResponse{Data: string(result)})
	})

	mux.HandleFunc("/update", func(w http.ResponseWriter, r *http.Request) {
	})

	return mux
}

func validateRequest(w http.ResponseWriter, r *http.Request) (dataRequest, error) {
	var req dataRequest
	if r.Method != http.MethodPost {
		http.Error(w, ErrMethodNotAllowed.Error(), http.StatusMethodNotAllowed)
		return req, ErrMethodNotAllowed
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, ErrInvalidJsonBody.Error(), http.StatusBadRequest)
		return req, ErrInvalidJsonBody
	}

	if req.ApiKey == "" {
		http.Error(w, ErrNoApiKey.Error(), http.StatusBadRequest)
		return req, ErrNoApiKey
	}

	return req, nil
}
