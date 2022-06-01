package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/whyslove/wildberries-l0/core/models"
)

type Purchase interface {
	AddNewPurchase(models.PurchaseDTO) error
	ReturnAllPurchasesDTO() ([]models.PurchaseDTO, error)
}

type Cache interface {
	RecoverAllPurchasesCache(purchases []models.PurchaseDTO)
	GetPurchaseByUidCache(uid string) (models.PurchaseDTO, error)
	AddNewPurchaseCache(purchase models.PurchaseDTO) error
}

type Repository struct {
	Purchase
	Cache
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{Purchase: NewPurchasePostgres(db), Cache: NewCacheImplementation()}
}
