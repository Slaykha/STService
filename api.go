package main

import (
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
		c.JSON(user)
		c.Status(fiber.StatusOK)
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
	spendings, err := a.service.GetSpendings()

	switch err {
	case nil:
		c.JSON(spendings)
		c.Status(fiber.StatusOK)
	default:
		c.Status(fiber.StatusInternalServerError)
	}

}
