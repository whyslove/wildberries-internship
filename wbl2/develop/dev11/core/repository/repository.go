package repository

import (
	"time"
	"whyslove/wbl2/dev11/core/domain"

	"github.com/jmoiron/sqlx"
)

// type Purchase interface {
// 	AddNewPurchase(models.PurchaseDTO) error
// 	ReturnAllPurchasesDTO() ([]models.PurchaseDTO, error)
// }

// type Cache interface {
// 	RecoverAllPurchasesCache(purchases []models.PurchaseDTO)
// 	GetPurchaseByUidCache(uid string) (models.PurchaseDTO, error)
// 	AddNewPurchaseCache(purchase models.PurchaseDTO) error
// }

type Event interface {
	CreateEvent(ev domain.Event) error
	UpdateEvent(ev domain.Event) error
	DeleteEvent(id int) error
	EventsForInterval(start_date time.Time, end_date time.Time, user_id int) ([]domain.Event, error)
}

type Repository struct {
	Event
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{Event: NewEventPostgres(db)}
}
