package handler

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"
	"whyslove/wbl2/dev11/core/domain"

	"github.com/sirupsen/logrus"
)

func (h *Handler) CreateEvent(w http.ResponseWriter, r *http.Request) {
	// r.URL.Query()
	event, err := ParseBodyToEvent(r.Body)
	if err != nil {
		logrus.Infof("bad json body %s", err)
		newErrorResponse(w, r, err, http.StatusBadRequest)
		return
	}
	logrus.Info(event)

	err = h.usecase.CreateEvent(event)
	if err != nil {
		newErrorResponse(w, r, err, http.StatusInternalServerError)
		return
	}
	okReposponse(w, r, http.StatusOK)

}

func (h *Handler) UpdateEvent(w http.ResponseWriter, r *http.Request) {
	event, err := ParseBodyToEvent(r.Body)
	if err != nil {
		logrus.Infof("bad json body %s", err)
		newErrorResponse(w, r, err, http.StatusBadRequest)
		return
	}
	logrus.Info(event)
	h.usecase.UpdateEvent(event)
	if err != nil {
		newErrorResponse(w, r, err, http.StatusInternalServerError)
		return
	}
	okReposponse(w, r, http.StatusOK)
}

func (h *Handler) DeleteEvent(w http.ResponseWriter, r *http.Request) {
	idInJson := struct {
		Id int `json:"id"`
	}{}
	byteBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		newErrorResponse(w, r, err, http.StatusBadRequest)
		return
	}
	json.Unmarshal(byteBody, &idInJson)
	if idInJson.Id == 0 {
		newErrorResponse(w, r, errors.New("bad body json"), http.StatusBadRequest)
		return
	}

	err = h.usecase.DeleteEvent(idInJson.Id)

	if err != nil {
		newErrorResponse(w, r, err, http.StatusInternalServerError)
		return
	}
	okReposponse(w, r, http.StatusOK)
}

func (h *Handler) EventsForDay(w http.ResponseWriter, r *http.Request) {
	urlArgs := ParseQueryString(r.URL)
	if urlArgs.date == NullTime || urlArgs.user_id == 0 {
		newErrorResponse(w, r, errors.New("bad formed url query"), http.StatusBadRequest)
		return
	}
	events, err := h.usecase.EventsForDay(urlArgs.date, urlArgs.user_id)
	if err != nil {
		newErrorResponse(w, r, err, http.StatusInternalServerError)
		return
	}
	sliceResponse(w, r, http.StatusOK, events)
}
func (h *Handler) EventsForWeek(w http.ResponseWriter, r *http.Request) {
	urlArgs := ParseQueryString(r.URL)
	if urlArgs.date == NullTime || urlArgs.user_id == 0 {
		newErrorResponse(w, r, errors.New("bad formed url query"), http.StatusBadRequest)
		return
	}
	events, err := h.usecase.EventsForWeek(urlArgs.date, urlArgs.user_id)
	if err != nil {
		newErrorResponse(w, r, err, http.StatusInternalServerError)
		return
	}
	sliceResponse(w, r, http.StatusOK, events)
}

func (h *Handler) EventsForMonth(w http.ResponseWriter, r *http.Request) {
	urlArgs := ParseQueryString(r.URL)
	if urlArgs.date == NullTime || urlArgs.user_id == 0 {
		newErrorResponse(w, r, errors.New("bad formed url query"), http.StatusBadRequest)
		return
	}
	events, err := h.usecase.EventsForMonth(urlArgs.date, urlArgs.user_id)

	if err != nil {
		newErrorResponse(w, r, err, http.StatusInternalServerError)
		return
	}
	sliceResponse(w, r, http.StatusOK, events)
}

var NullTime = time.Time{}

func ParseQueryString(r *url.URL) urlArgsStruct {
	var urlArgsStruct urlArgsStruct
	urlValues := r.Query()

	value, err := strconv.Atoi(urlValues.Get("user_id"))
	if err != nil {
		urlArgsStruct.user_id = 0
	} else {
		urlArgsStruct.user_id = value
	}

	time, err := time.Parse("2006-01-02T15:04:05Z07:00", urlValues.Get("date"))
	if err != nil {
		urlArgsStruct.date = NullTime
	} else {
		urlArgsStruct.date = time
	}
	return urlArgsStruct

}

func ParseBodyToEvent(Body io.ReadCloser) (domain.Event, error) {
	var event domain.Event
	decoder := json.NewDecoder(Body)
	err := decoder.Decode(&event)
	if err != nil {
		return event, err
	}
	return event, nil
}

type urlArgsStruct struct {
	user_id int
	date    time.Time
}
