package utils

import (
	"encoding/json"
	"net/http"
)

func ResponseSuccess(w http.ResponseWriter, r *http.Request, getList func() (any, error)) {
	list, err := getList()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(list)
}
