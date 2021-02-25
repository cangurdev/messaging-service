package handler

import (
	"encoding/json"
	"net/http"
)

func errorRespond(w http.ResponseWriter, r Response) {
	w.WriteHeader(r.StatusCode)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	json.NewEncoder(w).Encode(r)
}
