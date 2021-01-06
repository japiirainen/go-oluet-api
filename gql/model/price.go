package model

import (
	"time"
)

//Price is the type of one hinta
type Price struct {
	ID        int       `json:"id"`
	Date      time.Time `json:"date"`
	ProductID string    `json:"productId"`
	Hinta     float64   `json:"hinta"`
}
