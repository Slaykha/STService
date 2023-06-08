package models

import (
	"time"
)

type User struct {
	ID         string    `json:"id" bson:"id"`
	Name       string    `json:"name" bson:"name"`
	Email      string    `json:"email" bson:"email"`
	Password   []byte    `json:"password" bson:"password"`
	Currency   string    `json:"currency" bson:"currency"`
	DailyLimit float64   `json:"dailyLimit" bson:"dailyLimit"`
	CreatedAt  time.Time `json:"createdAt" bson:"createdAt"`
}

type UserRegisterDTO struct {
	Name     string `json:"name" bson:"name"`
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
	Currency string `json:"currency" bson:"currency"`
}

type UserAuth struct {
	ID         string    `json:"id" bson:"id"`
	Name       string    `json:"name" bson:"name"`
	Email      string    `json:"email" bson:"email"`
	Currency   string    `json:"currency" bson:"currency"`
	DailyLimit float64   `json:"dailyLimit" bson:"dailyLimit"`
	CreatedAt  time.Time `json:"createdAt" bson:"createdAt"`
}

type UserLoginDTO struct {
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}

type UserDTO struct {
	Name       string  `json:"name" bson:"name"`
	Email      string  `json:"email" bson:"email"`
	DailyLimit float64 `json:"dailyLimit" bson:"dailyLimit"`
}
type UserDailySpendingDTO struct {
	DailyLimit float64 `json:"dailyLimit" bson:"dailyLimit"`
}

type UserPasswordDTO struct {
	CurrentPassword string `json:"currentPassword" bson:"currentPassword"`
	NewPassword     string `json:"newPassword" bson:"newPassword"`
}

type Spending struct {
	ID           string    `json:"id" bson:"id"`
	UserID       string    `json:"userId" bson:"userId"`
	Money        float64   `json:"money" bson:"money"`
	SpendingType string    `json:"spendingType" bson:"spendingType"`
	SpendingDate time.Time `json:"spendingDate" bson:"spendingDate"`
}
type SpendingDTO struct {
	UserID       string    `json:"userId" bson:"userId"`
	Money        float64   `json:"money" bson:"money"`
	SpendingType string    `json:"spendingType" bson:"spendingType"`
	SpendingDate time.Time `json:"spendingDate" bson:"spendingDate"`
}
