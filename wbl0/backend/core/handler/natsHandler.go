package handler

import (
	"encoding/json"

	"github.com/nats-io/stan.go"
	"github.com/sirupsen/logrus"
	"github.com/whyslove/wildberries-l0/core/models"
)

func (h *Handler) ListenChannelPurchases(m *stan.Msg) {
	var purchase models.PurchaseDTO
	if err := json.Unmarshal(m.Data, &purchase); err != nil {
		logrus.Error("cannot parse purchase, drop message")
		return
	}
	if purchase.PurchaseUID != purchase.Payment_.Transaction {
		logrus.Error("purchase.PurchaseUID != purchase.Payment_.Transaction, drop message")
		return
	}
	logrus.Debugf("get new purchase from channel %v", purchase)
	err := h.repository.AddNewPurchase(purchase)
	if err != nil {
		logrus.Errorf("error %s while adding in db, drop message", err.Error())
		return
	}

	err = h.repository.AddNewPurchaseCache(purchase)
	if err != nil {
		logrus.Error(err.Error())
	} else {
		logrus.Infof("added new purchase to cache with uid %s", purchase.PurchaseUID)
	}
}
