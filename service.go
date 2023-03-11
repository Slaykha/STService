package main

import (
	"time"

	"github.com/Slaykha/STService/errors"
	"github.com/Slaykha/STService/helpers"
	"github.com/Slaykha/STService/models"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	repository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) CreateUser(userDTO models.UserRegisterDTO) (*models.User, error) {

	password, _ := bcrypt.GenerateFromPassword([]byte(userDTO.Password), 8)
	user := models.User{
		ID:        helpers.CreateID(),
		Name:      userDTO.Name,
		Email:     userDTO.Email,
		Password:  password,
		CreatedAt: time.Now().UTC(),
	}

	err := s.repository.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return &user, err
}
func (s *Service) UserLogin(userDTO models.UserLoginDTO) (*models.User, error) {
	user, err := s.repository.FindUser(userDTO.Email)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userDTO.Password))
	if err != nil {
		return nil, errors.LoginCredentialsWrong
	}

	return user, nil
}

func (s *Service) GetUser(claims *jwt.StandardClaims) models.UserAuth {

	user := s.repository.GetUser(claims)

	return *user
}

func (s *Service) CreateSpending(spendingDTO models.SpendingDTO) (*models.Spending, error) {

	spending := models.Spending{
		ID:           helpers.CreateID(),
		UserID:       spendingDTO.UserID,
		Money:        spendingDTO.Money,
		Currency:     spendingDTO.Currency,
		SpendingType: spendingDTO.SpendingType,
		SpendingDate: time.Now().UTC(),
	}

	err := s.repository.CreateSpending(spending)
	if err != nil {
		return nil, err
	}

	return &spending, err
}

func (s *Service) GetSpendings(userID string) ([]models.Spending, error) {
	spendings, err := s.repository.GetSpendings(userID)
	if err != nil {
		return nil, err
	}

	return spendings, err

}
