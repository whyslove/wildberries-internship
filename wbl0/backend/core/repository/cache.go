package repository

import (
	"fmt"
	"sync"

	"github.com/whyslove/wildberries-l0/core/models"
)

type CacheImplementation struct {
	mu    sync.RWMutex
	cache map[string]models.PurchaseDTO
}

func NewCacheImplementation() *CacheImplementation {
	return &CacheImplementation{cache: make(map[string]models.PurchaseDTO)}
}

func (ci *CacheImplementation) AddNewPurchaseCache(purchase models.PurchaseDTO) error {
	ci.mu.Lock()
	defer ci.mu.Unlock()

	if _, ok := ci.cache[purchase.PurchaseUID]; ok {
		return fmt.Errorf(`error in saving object with uid=%s to cache, 
			it is already in it`, purchase.PurchaseUID)
	}
	ci.cache[purchase.PurchaseUID] = purchase
	return nil
}

func (ci *CacheImplementation) GetPurchaseByUidCache(uid string) (models.PurchaseDTO, error) {
	ci.mu.RLock()
	defer ci.mu.RUnlock()

	purchase, ok := ci.cache[uid]
	if !ok {
		return models.PurchaseDTO{}, fmt.Errorf("this uid is not presented in cache")
	}
	return purchase, nil
}

func (ci *CacheImplementation) RecoverAllPurchasesCache(purchases []models.PurchaseDTO) {
	ci.mu.Lock()
	defer ci.mu.Unlock()

	for _, purchase := range purchases {
		ci.cache[purchase.PurchaseUID] = purchase
	}
}
