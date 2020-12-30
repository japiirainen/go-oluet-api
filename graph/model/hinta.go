package model

import (
	"time"
)

//Hinta is the type of one hinta
type Hinta struct {
	ID        int       `json:"id"`
	Date      time.Time `json:"date"`
	ProductID string    `json:"productId"`
	Hinta     float64   `json:"hinta"`
}
