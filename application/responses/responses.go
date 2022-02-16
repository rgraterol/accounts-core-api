package responses

import (
	"encoding/json"
	"net/http"
)

func OK(w http.ResponseWriter, response interface{}) {
	answer(w, http.StatusOK, response)
}

func answer(w http.ResponseWriter, status int, response interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(response)
}
