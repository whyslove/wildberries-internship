package main

import (
	"flag"
	"time"

	"github.com/nats-io/stan.go"
	"github.com/sirupsen/logrus"
)

var (
	clusterID   = "test-cluster"
	clientID    = "my-sender-id"
	channelName = "purchases"
)

var message = [][]byte{
	[]byte(`{
    "order_uid": "b563feb7b2b84b6test",
    "track_number": "WBILMTESTTRACK",
    "entry": "WBIL",
    "delivery": {
      "name": "Test Testov",
      "phone": "+9720000000",
      "zip": "2639809",
      "city": "Kiryat Mozkin",
      "address": "Ploshad Mira 15",
      "region": "Kraiot",
      "email": "test@gmail.com"
    },
    "payment": {
      "transaction": "b563feb7b2b84b6test",
      "request_id": "",
      "currency": "USD",
      "provider": "wbpay",
      "amount": 1817,
      "payment_dt": 1637907727,
      "bank": "alpha",
      "delivery_cost": 1500,
      "goods_total": 317,
      "custom_fee": 0
    },
    "items": [
      {
        "chrt_id": 9934930,
        "track_number": "WBILMTESTTRACK",
        "price": 453,
        "rid": "ab4219087a764ae0btest",
        "name": "Mascaras",
        "sale": 30,
        "size": "0",
        "total_price": 317,
        "nm_id": 2389212,
        "brand": "Vivienne Sabo",
        "status": 202
      }
    ],
    "locale": "en",
    "internal_signature": "",
    "customer_id": "test",
    "delivery_service": "meest",
    "shardkey": "9",
    "sm_id": 99,
    "date_created": "2021-11-26T06:22:19Z",
    "oof_shard": "1"
  }`),
	[]byte(`{
    "order_uid": "b563feb7b2b84b6test1",
    "track_number": "WBILMTESTTRACK1",
    "entry": "WBIL",
    "delivery": {
      "name": "Test Testov",
      "phone": "+9720000000",
      "zip": "2639809",
      "city": "Kiryat Mozkin",
      "address": "Ploshad Mira 15",
      "region": "Kraiot",
      "email": "test@gmail.com"
    },
    "payment": {
      "transaction": "b563feb7b2b84b6test1",
      "request_id": "",
      "currency": "USD",
      "provider": "wbpay",
      "amount": 1817,
      "payment_dt": 1637907727,
      "bank": "alpha",
      "delivery_cost": 1500,
      "goods_total": 317,
      "custom_fee": 0
    },
    "items": [
      {
        "chrt_id": 9934931,
        "track_number": "WBILMTESTTRACK1",
        "price": 453,
        "rid": "ab4219087a764ae0btest",
        "name": "Mascaras",
        "sale": 30,
        "size": "0",
        "total_price": 317,
        "nm_id": 2389212,
        "brand": "Vivienne Sabo",
        "status": 202
      },
      {
        "chrt_id": 9934931,
        "track_number": "WBILMTESTTRACK1",
        "price": 453,
        "rid": "ab4219087a764ae0btest",
        "name": "Mascaras",
        "sale": 30,
        "size": "0",
        "total_price": 317,
        "nm_id": 2389212,
        "brand": "Vivienne Sabo",
        "status": 202
      }
    ],
    "locale": "en",
    "internal_signature": "",
    "customer_id": "test",
    "delivery_service": "meest",
    "shardkey": "9",
    "sm_id": 99,
    "date_created": "2021-11-26T06:22:19Z",
    "oof_shard": "1"
  }`), //next will be bad
	[]byte(`{ 
    "order_uid": "b563feb7b2b84b6test2",
    "track_number": "WBILMTESTTRACK2",
    "entry": "WBIL",
    "delivery": {
      "name": "Test Testov",
      "phone": "+9720000000",
      "zip": "2639809",
      "city": "Kiryat Mozkin",
      "address": "Ploshad Mira 15",
      "region": "Kraiot",
      "email": "test@gmail.com"
    },
    "payment": {
      "transaction": "b563feb7b2b84b6test1",
      "request_id": "",
      "currency": "USD",
      "provider": "wbpay",
      "amount": 1817,
      "payment_dt": 1637907727,
      "bank": "alpha",
      "delivery_cost": 1500,
      "goods_total": 317,
      "custom_fee": 0
    },
    "items": [
      {
        "chrt_id": 9934931,
        "track_number": "WBILMTESTTRACK2",
        "price": 453,
        "rid": "ab4219087a764ae0btest",
        "name": "Mascaras",
        "sale": 30,
        "size": "0",
        "total_price": 317,
        "nm_id": 2389212,
        "brand": "Vivienne Sabo",
        "status": 202
      },
      {
        "chrt_id": 9934932,
        "track_number": "WBILMTESTTRACK2",
        "price": 453,
        "rid": "ab4219087a764ae0btest",
        "name": "Mascaras",
        "sale": 30,
        "size": "0",
        "total_price": 317,
        "nm_id": 2389212,
        "brand": "Vivienne Sabo",
        "status": 202
      }
    ],
    "locale": "en",
    "internal_signature": "",
    "customer_id": "test",
    "delivery_service": "meest",
    "shardkey": "9",
    "sm_id": 99,
    "date_created": "2021-11-26T06:22:19Z",
    "oof_shard": "1"
  }`),
	[]byte(`{
    "order_uid": "b563feb7b2b84b6test2",
    "track_number": "WBILMTESTTRACK3",
    "entry": "WBIL",
    "delivery": {
      "name": "Test Testov",
      "phone": "+9720000000",
      "zip": "2639809",
      "city": "Kiryat Mozkin",
      "address": "Ploshad Mira 15",
      "region": "Kraiot",
      "email": "test@gmail.com"
    },
    "payment": {
      "transaction": "b563feb7b2b84b6test2",
      "request_id": "",
      "currency": "USD",
      "provider": "wbpay",
      "amount": 1817,
      "payment_dt": 1637907727,
      "bank": "alpha",
      "delivery_cost": 1500,
      "goods_total": 317,
      "custom_fee": 0
    },
    "items": [
      {
        "chrt_id": 9934931,
        "track_number": "WBILMTESTTRACK3",
        "price": 453,
        "rid": "ab4219087a764ae0btest",
        "name": "Mascaras",
        "sale": 30,
        "size": "0",
        "total_price": 317,
        "nm_id": 2389212,
        "brand": "Vivienne Sabo",
        "status": 202
      },
      {
        "chrt_id": 9934932,
        "track_number": "WBILMTESTTRACK3",
        "price": 453,
        "rid": "ab4219087a764ae0btest",
        "name": "Mascaras",
        "sale": 30,
        "size": "0",
        "total_price": 317,
        "nm_id": 2389212,
        "brand": "Vivienne Sabo",
        "status": 202
      }
    ],
    "locale": "en",
    "internal_signature": "",
    "customer_id": "test",
    "delivery_service": "meest",
    "shardkey": "9",
    "sm_id": 99,
    "date_created": "2021-11-26T06:22:19Z",
    "oof_shard": "1"
  }`),
}

func main() {
	sc, err := stan.Connect(clusterID, clientID)
	if err != nil {
		logrus.Error(err)
		return
	}

	useInit := flag.Bool("init", true, "what part of message send")
	flag.Parse()

	if *useInit {
		err = sc.Publish(channelName, message[0])
		if err != nil {
			logrus.Error(err)
			return
		}
		err = sc.Publish(channelName, message[1])
		if err != nil {
			logrus.Error(err)
			return
		}
		err = sc.Publish(channelName, message[2])
		if err != nil {
			logrus.Error(err)
			return
		}
	} else {
		err = sc.Publish(channelName, message[3])
		if err != nil {
			logrus.Error(err)
			return
		}
	}

	time.Sleep(time.Second)
}
