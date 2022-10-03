package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/Slaykha/STService/models"
	"github.com/google/uuid"
)

type Service struct {
	repository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) CreateSpending(spendingDTO models.SpendingDTO) (*models.Spending, error) {

	spending := models.Spending{
		ID:           createID(),
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

func (s *Service) GetSpendings() ([]models.Spending, error) {
	spendings, err := s.repository.GetSpendings()
	if err != nil {
		fmt.Println("1")
		return nil, err
	}

	return spendings, err

}

func createID() (id string) {
	id = uuid.New().String()

	id = strings.ReplaceAll(id, "-", "")

	id = id[0:8]

	return
}
