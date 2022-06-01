package handler

import (
	"encoding/json"
	"net/http"
	"whyslove/wbl2/dev11/core/domain"
)

func newErrorResponse(w http.ResponseWriter, r *http.Request, err error, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(`{"error": "` + err.Error() + "\"}"))
}

func okReposponse(w http.ResponseWriter, r *http.Request, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(`{"message": "ok"}`))
}

func sliceResponse(w http.ResponseWriter, r *http.Request, status int, slice []domain.Event) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if len(slice) == 0 {
		w.Write([]byte(`{"message": "no events"}`))
	} else {
		json, _ := json.Marshal(slice)
		w.Write(json)
	}
}
