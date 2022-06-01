package handler

import (
	"net/http"
	"whyslove/wbl2/dev11/core/usecase"
	// gin-swagger middleware
	// swagger embed files
)

type Handler struct {
	usecase *usecase.UseCase
}

func NewHandler(u *usecase.UseCase) *Handler {
	return &Handler{usecase: u}
}

func (h *Handler) InitRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle("/create-event", h.MiddlewareLogger(http.HandlerFunc(h.CreateEvent)))
	mux.Handle("/update-event", h.MiddlewareLogger(http.HandlerFunc(h.UpdateEvent)))
	mux.Handle("/delete-event", h.MiddlewareLogger(http.HandlerFunc(h.DeleteEvent)))
	mux.Handle("/events_for_day", h.MiddlewareLogger(http.HandlerFunc(h.EventsForDay)))
	mux.Handle("/events_for_week", h.MiddlewareLogger(http.HandlerFunc(h.EventsForWeek)))
	mux.Handle("/events_for_month", h.MiddlewareLogger(http.HandlerFunc(h.EventsForMonth)))

	return mux
}
