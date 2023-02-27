package main

import (
	"time"

	"github.com/Slaykha/STService/errors"
	"github.com/Slaykha/STService/helpers"
	"github.com/Slaykha/STService/models"
	"github.com/gofiber/fiber"
)

type Api struct {
	service *Service
}

func NewAPI(service *Service) *Api {
	return &Api{
		service: service,
	}
}

func (a *Api) HandleUserCreate(c *fiber.Ctx) {
	var userDTO models.UserRegisterDTO

	err := c.BodyParser(&userDTO)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
	}

	user, err := a.service.CreateUser(userDTO)
	switch err {
	case nil:
		c.JSON(user)
		c.Status(fiber.StatusCreated)
	default:
		c.Status(fiber.StatusInternalServerError)
	}
}

func (a *Api) HandleUserLogin(c *fiber.Ctx) {
	var userDTO models.UserLoginDTO

	err := c.BodyParser(&userDTO)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
	}

	user, err := a.service.UserLogin(userDTO)

	switch err {
	case nil:
		cookie := fiber.Cookie{
			Name:     "user_token",
			Value:    helpers.CreateUserToken(user.ID),
			Expires:  time.Now().Add(time.Hour * 24),
			HTTPOnly: true,
		}

		c.Cookie(&cookie)
		c.JSON(fiber.Map{"message": "success"})
		c.Status(fiber.StatusOK)
	case errors.LoginCredentialsWrong:
		c.JSON(err.Error())
		c.Status(fiber.StatusNotFound)
	default:
		c.Status(fiber.StatusInternalServerError)
	}
}

func (a *Api) HandleCreateSpending(c *fiber.Ctx) {
	userId := c.Params("userID")

	spendingDTO := models.SpendingDTO{}
	spendingDTO.UserID = userId
	err := c.BodyParser(&spendingDTO)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
	}

	spending, err := a.service.CreateSpending(spendingDTO)
	switch err {
	case nil:
		c.JSON(spending)
		c.Status(fiber.StatusCreated)
	default:
		c.Status(fiber.StatusInternalServerError)
	}
}

func (a *Api) HandleGetSpendings(c *fiber.Ctx) {
	userID := c.Params("userID")

	spendings, err := a.service.GetSpendings(userID)

	switch err {
	case nil:
		c.JSON(spendings)
		c.Status(fiber.StatusOK)
	default:
		c.Status(fiber.StatusInternalServerError)
	}

}
