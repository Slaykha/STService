package models

import (
	"time"
)

type Spending struct {
	ID           string    `json:"_id" bson:"_id"`
	Money        float64   `json:"money" bson:"money"`
	Currency     string    `json:"currency" bson:"currency"`
	SpendingType string    `json:"spendingType" bson:"spendingType"`
	SpendingDate time.Time `json:"SpendingDate" bson:"SpendingDate"`
}

type SpendingDTO struct {
	Money        float64 `json:"money" bson:"money"`
	Currency     string  `json:"currency" bson:"currency"`
	SpendingType string  `json:"spendingType" bson:"spendingType"`
}
