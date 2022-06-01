package usecase

import (
	"time"
	"whyslove/wbl2/dev11/core/domain"
	"whyslove/wbl2/dev11/core/repository"
)

type EventUseCase struct {
	repos *repository.Repository
}

func NewEventUseCase(repos *repository.Repository) *EventUseCase {
	return &EventUseCase{repos: repos}
}

func (evu *EventUseCase) CreateEvent(ev domain.Event) error {
	return evu.repos.CreateEvent(ev)
}

func (evu *EventUseCase) UpdateEvent(ev domain.Event) error {
	return evu.repos.UpdateEvent(ev)
}

func (evu *EventUseCase) DeleteEvent(id int) error {
	return evu.repos.DeleteEvent(id)
}

func (evu *EventUseCase) EventsForDay(start time.Time, user_id int) ([]domain.Event, error) {
	return evu.repos.EventsForInterval(start, start.AddDate(0, 0, 1), user_id)
}

func (evu *EventUseCase) EventsForWeek(start time.Time, user_id int) ([]domain.Event, error) {
	return evu.repos.EventsForInterval(start, start.AddDate(0, 0, 7), user_id)
}

func (evu *EventUseCase) EventsForMonth(start time.Time, user_id int) ([]domain.Event, error) {
	return evu.repos.EventsForInterval(start, start.AddDate(0, 1, 0), user_id)
}
