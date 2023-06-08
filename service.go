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
		Currency:  userDTO.Currency,
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

func (s *Service) GetUser(claims *jwt.StandardClaims) (*models.UserAuth, error) {

	user, err := s.repository.GetUserInfo(claims.Issuer)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Service) UpdateUser(userId string, user models.UserDTO) (*models.UserAuth, error) {
	userModel, err := s.repository.GetUser(userId)
	if err != nil {
		return nil, err
	}

	userModel.Name = user.Name
	userModel.Email = user.Email
	userModel.DailyLimit = user.DailyLimit

	updatedUser, err := s.repository.UpdateUser(*userModel)
	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}

func (s *Service) UpdateUserDailySpending(userId string, userDailySpending models.UserDailySpendingDTO) (*models.UserAuth, error) {
	userModel, err := s.repository.GetUser(userId)
	if err != nil {
		return nil, err
	}
	userModel.DailyLimit = userDailySpending.DailyLimit

	updatedUser, err := s.repository.UpdateUser(*userModel)
	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}

func (s *Service) UpdateUserPassword(userId string, user models.UserPasswordDTO) (*models.UserAuth, error) {
	userModel, err := s.repository.GetUser(userId)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(userModel.Password), []byte(user.CurrentPassword))
	if err != nil {
		return nil, errors.WrongPassword
	}

	newPassword, _ := bcrypt.GenerateFromPassword([]byte(user.NewPassword), 8)

	userModel.Password = newPassword

	updatedUser, err := s.repository.UpdateUser(*userModel)
	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}
func (s *Service) CreateSpending(spendingDTO models.SpendingDTO) (*models.Spending, error) {

	spending := models.Spending{
		ID:           helpers.CreateID(),
		UserID:       spendingDTO.UserID,
		Money:        spendingDTO.Money,
		SpendingType: spendingDTO.SpendingType,
		SpendingDate: spendingDTO.SpendingDate,
	}

	err := s.repository.CreateSpending(spending)
	if err != nil {
		return nil, err
	}

	return &spending, err
}

func (s *Service) GetSpendings(userID, spendingType string, date time.Time) ([]models.Spending, error) {
	spendings, err := s.repository.GetSpendings(userID, spendingType, date)
	if err != nil {
		return nil, err
	}

	return spendings, err

}

func (s *Service) DeleteSpending(spendingID string) error {
	err := s.repository.DeleteSpending(spendingID)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) GetTodaysTotal(userId string) (float32, error) {
	spendings, err := s.repository.GetSpendings(userId, "", time.Now())
	if err != nil {
		return 0, err
	}

	var total float32

	for _, spending := range spendings {
		total += float32(spending.Money)
	}

	return total, nil
}
