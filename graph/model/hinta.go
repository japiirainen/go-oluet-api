package model

import (
	"time"
)

type Hinta struct {
	ID        string    `json:"id"`
	Date      time.Time `json:"date"`
	ProductID string    `json:"productId"`
	Hinta     float64  `json:"hinta"`
}


