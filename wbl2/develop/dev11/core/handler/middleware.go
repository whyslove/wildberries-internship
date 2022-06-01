package handler

import (
	"log"
	"net/http"
)

func (h *Handler) MiddlewareLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r)
		next.ServeHTTP(w, r)
	})
}
