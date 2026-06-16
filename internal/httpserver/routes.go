package httpserver

import (
	"encoding/json"
	"net/http"
)

func NewRouter() http.Handler{
	mux := http.NewServeMux()
	mux.HandleFunc("/health", healthHandler)

	return mux
}

func healthHandler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	resp := map[string]string{
		"status": "ok",
	}

	json.NewEncoder(w).Encode(resp)
}
