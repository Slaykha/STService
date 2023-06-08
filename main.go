package main

import (
	"fmt"

	"github.com/gofiber/cors"
	"github.com/gofiber/fiber"
)

type Config struct {
	AppPort         int
	Host            string
	DBReplicaSetUrl string
}

var config Config

func main() {
	fmt.Println("Service Starting...")
	setConfig()

	repository := NewRepository(config.DBReplicaSetUrl)
	service := NewService(repository)
	API := NewAPI(service)

	app := SetupApp(API)

	fmt.Println("Spending Tracker service started at ", config.AppPort, "  ...")

	app.Get("/status", func(c *fiber.Ctx) {
		c.Status(fiber.StatusOK)
	})

	//User
	app.Post("/user/register", API.HandleUserCreate)
	app.Post("/user/login", API.HandleUserLogin)
	app.Post("/user/logout", API.HandleUserLogout)
	app.Get("/user/token", API.HandleGetUser)
	app.Put("/user/:id", API.HandleUpdateUser)
	app.Put("/user/:id/dailyLimit", API.HandleUpdateUserDailySpending)
	app.Put("/user/:id/password", API.HandleUpdateUserPassword)

	//Spending
	app.Post("/spending/:userId", API.HandleCreateSpending)
	app.Get("/spendings/:userId", API.HandleGetSpendings)
	app.Delete("/spending/:spendingId", API.HandleDeleteSpending)
	app.Get("/spendings/:userId/today", API.HandleGetTodaysTotal)

	app.Listen(config.AppPort)
}

func SetupApp(API *Api) *fiber.App {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowHeaders:     []string{"Origin, Content-Type, Accept"},
		AllowCredentials: true,
	}))

	return app
}

func setConfig() {
	config = Config{
		AppPort:         12345,
		Host:            "http://localhost:12345",
		DBReplicaSetUrl: "mongodb+srv://admin:HkJpLyv1MclTvMIc@spendingtraacker.ybzvy6n.mongodb.net/?retryWrites=true&w=majority",
	}
}
