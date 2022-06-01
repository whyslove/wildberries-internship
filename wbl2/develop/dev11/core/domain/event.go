package domain

import (
	"time"
)

//Event domain
type Event struct {
	ID     int       `json:"id"`
	UserID int       `json:"user_id"`
	Date   time.Time `json:"date"`
	Descr  string    `json:"descr"`
}
