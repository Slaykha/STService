package models

import (
	"time"
)

type User struct {
	ID        string    `json:"_id" bson:"_id"`
	Name      string    `json:"name" bson:"name"`
	Email     string    `json:"email" bson:"email"`
	Password  []byte    `json:"password" bson:"password"`
	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
}

type UserRegisterDTO struct {
	Name     string `json:"name" bson:"name"`
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}

type UserAuth struct {
	ID        string    `json:"_id" bson:"_id"`
	Name      string    `json:"name" bson:"name"`
	Email     string    `json:"email" bson:"email"`
	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
}

type UserLoginDTO struct {
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}

type Spending struct {
	ID           string    `json:"_id" bson:"_id"`
	UserID       string    `json:"userId" bson:"userId"`
	Money        float64   `json:"money" bson:"money"`
	Currency     string    `json:"currency" bson:"currency"`
	SpendingType string    `json:"spendingType" bson:"spendingType"`
	SpendingDate time.Time `json:"spendingDate" bson:"spendingDate"`
}
type SpendingDTO struct {
	UserID       string  `json:"userId" bson:"userId"`
	Money        float64 `json:"money" bson:"money"`
	Currency     string  `json:"currency" bson:"currency"`
	SpendingType string  `json:"spendingType" bson:"spendingType"`
}
