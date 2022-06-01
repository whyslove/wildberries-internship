package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"github.com/whyslove/wildberries-l0/core/models"
)

type PurchasePostgres struct {
	db *sqlx.DB
}

func NewPurchasePostgres(db *sqlx.DB) *PurchasePostgres {
	return &PurchasePostgres{db: db}
}

func (u *PurchasePostgres) AddNewPurchase(pdto models.PurchaseDTO) error {
	tx, err := u.db.Begin()
	if err != nil {
		return fmt.Errorf("error in starting transaction, %s", err.Error())
	}
	defer tx.Rollback()

	add_purchase := "CALL add_purchase($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11);"
	_, err = tx.Exec(add_purchase,
		pdto.PurchaseUID, pdto.TrackNumber, pdto.Entry,
		pdto.Locale, pdto.InternalSignature, pdto.CustomerID, pdto.DeliveryService,
		pdto.Shardkey, pdto.SmID, pdto.DateCreated, pdto.OofShard,
	)
	if err != nil {
		return fmt.Errorf("error in adding in table purchase, %s", err.Error())
	}

	add_payment := "call add_payment($1, $2, $3, $4, $5, $6, $7, $8, $9, $10);"
	_, err = tx.Exec(add_payment,
		pdto.Payment_.Transaction, pdto.Payment_.RequestID, pdto.Payment_.Currency,
		pdto.Payment_.Provider, pdto.Payment_.Amount, pdto.Payment_.PaymentDt,
		pdto.Payment_.Bank, pdto.Payment_.DeliveryCost, pdto.Payment_.GoodsTotal,
		pdto.Payment_.CustomFee,
	)
	if err != nil {
		return fmt.Errorf("error in adding in table payment, %s", err.Error())
	}

	add_delivery := "call add_delivery($1, $2, $3, $4, $5, $6, $7, $8);"
	_, err = tx.Exec(add_delivery, pdto.PurchaseUID, pdto.Deliver.Name, pdto.Deliver.Phone,
		pdto.Deliver.Zip, pdto.Deliver.City, pdto.Deliver.Address, pdto.Deliver.Region, pdto.Deliver.Email,
	)
	if err != nil {
		return fmt.Errorf("error in adding in table delivery, %s", err.Error())
	}

	add_item := "call add_item_to_purchase($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12);"
	for _, item := range pdto.Items_ {
		_, err = tx.Exec(add_item, item.ChrtID, item.TrackNumber, item.Price,
			item.Rid, item.Name, item.Sale, item.Size, item.TotalPrice, item.NmID,
			item.Brand, item.Status, pdto.PurchaseUID)
		if err != nil {
			return fmt.Errorf("error in item in item_tables, %s", err.Error())
		}
	}
	tx.Commit()
	return nil
}

func (u *PurchasePostgres) ReturnAllPurchasesDTO() ([]models.PurchaseDTO, error) {
	var allPurchases []models.PurchaseDTO
	var allPurchasesUid []string
	var err error
	tx, err := u.db.Begin()
	if err != nil {
		logrus.Error(err.Error())
		return nil, fmt.Errorf("error at starting transaction in ReturnAllPurchasesDTO()")
	}

	qGetAllUids := fmt.Sprintf("SELECT (purchase_uid) FROM %s;", purchasesTable)
	rows, err := tx.Query(qGetAllUids)
	if err != nil {
		logrus.Error(err.Error())
		return nil, fmt.Errorf("error select all purchases id's in ReturnAllPurchasesDTO()")
	}

	for rows.Next() {
		var row string
		err = rows.Scan(&row)
		if err != nil {
			logrus.Error(err.Error())
			rows.Close()
			return nil, fmt.Errorf("error reading id row ReturnAllPurchasesDTO()")
		}
		allPurchasesUid = append(allPurchasesUid, row)
	}
	err = rows.Err()
	if err != nil {
		logrus.Error(err.Error())
		return nil, fmt.Errorf("error after reading id row ReturnAllPurchasesDTO()")
	}
	rows.Close()

	qGetDelivery, err := tx.Prepare(`SELECT name, phone, zip, city, adress, region, email
	FROM delivery WHERE purchase_uid = $1`)
	if err != nil {
		fmt.Printf("db.Prepare error: %v\n", err)
		return nil, err
	}
	qGetPayment, err := tx.Prepare(`SELECT money_transaction, request_id, currency,
			provider, amount, payment_dt, bank, delivery_cost, goods_total, custom_fee
			FROM payment WHERE money_transaction = $1`)
	if err != nil {
		fmt.Printf("db.Prepare error: %v\n", err)
		return nil, err
	}
	qGetPurchase, err := tx.Prepare(`SELECT purchase_uid, track_number, entry, locale, internal_signature, customer_id,
			delivery_service, shardkey, sm_id, date_created, oof_shard FROM purchase WHERE purchase_uid = $1`)
	if err != nil {
		fmt.Printf("db.Prepare error: %v\n", err)
		return nil, err
	}
	qGetAllPurchaseItemsId, err := tx.Prepare(`select chrt_id, track_number, price, rid, name, sale, 
		size, total_p,nm_id, brand, code_status, amount
		FROM item 
		inner JOIN purchase_item on item.chrt_id = purchase_item.item_chrt_id 
			and chrt_id in 
				(SELECT item_chrt_id FROM purchase_item WHERE purchase_uid = $1);`)
	if err != nil {
		fmt.Printf("db.Prepare error: %v\n", err)
		return nil, err
	}

	for _, uid := range allPurchasesUid {
		var purchase models.PurchaseDTO
		var delivery models.Delivery
		var payment models.Payment
		var items []models.Item

		//  := fmt.Sprintf(, deliveryTable)
		err := qGetDelivery.QueryRow(uid).Scan(&delivery.Name, &delivery.Phone,
			&delivery.Zip, &delivery.City, &delivery.Address,
			&delivery.Region, &delivery.Email)
		if err != nil {
			logrus.Error(err.Error())
			return nil, fmt.Errorf("error in get delivery for %s", uid)
		}

		err = qGetPayment.QueryRow(uid).Scan(&payment.Transaction, &payment.RequestID, &payment.Currency,
			&payment.Provider, &payment.Amount, &payment.PaymentDt, &payment.Bank, &payment.DeliveryCost,
			&payment.GoodsTotal, &payment.CustomFee)
		if err != nil {
			logrus.Error(err.Error())
			return nil, fmt.Errorf("error in get payment for %s", uid)
		}

		err = qGetPurchase.QueryRow(uid).Scan(&purchase.PurchaseUID, &purchase.TrackNumber, &purchase.Entry, &purchase.Locale,
			&purchase.InternalSignature, &purchase.CustomerID, &purchase.DeliveryService, &purchase.Shardkey, &purchase.SmID,
			&purchase.DateCreated, &purchase.OofShard)
		if err != nil {
			logrus.Error(err.Error())
			return nil, fmt.Errorf("error in load purchase to %s", uid)
		}

		rows, err = qGetAllPurchaseItemsId.Query(uid)
		if err != nil {
			logrus.Error(err.Error())
			return nil, fmt.Errorf("error in load many ites to %s", uid)
		}
		for rows.Next() {
			var item models.Item
			var amount int
			err = rows.Scan(&item.ChrtID, &item.TrackNumber, &item.Price, &item.Rid, &item.Name, &item.Sale,
				&item.Size, &item.TotalPrice, &item.NmID, &item.Brand, &item.Status, &amount)
			if err != nil {
				logrus.Error(err.Error())
				rows.Close()
				return nil, fmt.Errorf("error in load row from items to %s", uid)
			}
			for i := 0; i < amount; i++ {
				items = append(items, item)
			}

		}
		err = rows.Err()
		if err != nil {
			logrus.Error(err.Error())
			rows.Close()
			return nil, fmt.Errorf("error after reading items rows")
		}
		rows.Close()

		purchase.Items_ = items
		purchase.Deliver = delivery
		purchase.Payment_ = payment
		allPurchases = append(allPurchases, purchase)
	}
	qGetDelivery.Close()
	qGetPayment.Close()
	qGetPurchase.Close()
	qGetAllPurchaseItemsId.Close()

	tx.Commit()
	return allPurchases, nil

}
