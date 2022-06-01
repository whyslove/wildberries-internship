package repository

import (
	"time"
	"whyslove/wbl2/dev11/core/domain"

	"github.com/jmoiron/sqlx"
)

type EventPostgres struct {
	db *sqlx.DB
}

func NewEventPostgres(db *sqlx.DB) *EventPostgres {
	return &EventPostgres{db: db}
}

func (evp *EventPostgres) CreateEvent(ev domain.Event) error {
	createEvent := `INSERT INTO event (user_id, date, descr) VALUES ($1, $2, $3)`
	_, err := evp.db.Exec(createEvent, ev.UserID, ev.Date, ev.Descr)
	if err != nil {
		return err
	}
	return nil
}

func (evp *EventPostgres) UpdateEvent(ev domain.Event) error {
	updateEvent := `UPDATE event  SET
         		event.id = $1, event.user_id = $2, event.date = $3, event.descr = $4
				WHERE event.id = id;`

	_, err := evp.db.Exec(updateEvent, ev.ID, ev.UserID, ev.Date, ev.Descr)
	if err != nil {
		return err
	}
	return nil
}

func (evp *EventPostgres) DeleteEvent(id int) error {
	deleteEvent := `DELETE FROM event WHERE
					event.id = $1`
	_, err := evp.db.Exec(deleteEvent, id)
	if err != nil {
		return err
	}
	return nil
}

func (evp *EventPostgres) EventsForInterval(start time.Time, end time.Time, user_id int) ([]domain.Event, error) {
	var events []domain.Event
	getEventsForInterval := `SELECT * FROM event WHERE
	event.user_id=$1 AND event.date >= $2 AND event.date <= $3
	`
	err := evp.db.Select(&events, getEventsForInterval, user_id, start, end)
	if err != nil {
		return nil, err
	}
	return events, nil
}
