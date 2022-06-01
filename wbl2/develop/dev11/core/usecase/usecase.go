package usecase

import (
	"time"
	"whyslove/wbl2/dev11/core/domain"
	"whyslove/wbl2/dev11/core/repository"
)

type UseCase struct {
	Event
}

type Event interface {
	CreateEvent(domain.Event) error
	UpdateEvent(domain.Event) error
	DeleteEvent(id int) error
	EventsForDay(time.Time, int) ([]domain.Event, error)
	EventsForWeek(time.Time, int) ([]domain.Event, error)
	EventsForMonth(time.Time, int) ([]domain.Event, error)
}

func NewUseCase(repos *repository.Repository) *UseCase {
	return &UseCase{Event: NewEventUseCase(repos)}
}
