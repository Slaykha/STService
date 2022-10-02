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

func (a *Api) HandleCreateSpending(c *fiber.Ctx) {

	spendingDTO := models.SpendingDTO{}
	err := c.BodyParser(&spendingDTO)
	if err != nil{
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
