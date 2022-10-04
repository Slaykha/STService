package main

import (
	"fmt"

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

	fmt.Println("Order List service started at ", config.AppPort, "  ...")

	app.Listen(config.AppPort)
}

func SetupApp(API *Api) *fiber.App {
	app := fiber.New()

	app.Get("/status", func(c *fiber.Ctx) {
		c.Status(fiber.StatusOK)
	})

	//User
	app.Post("/user/register", API.HandleUserCreate)
	app.Post("/user/login", API.HandleUserLogin)

	//Spending
	app.Post("/spending/:userID", API.HandleCreateSpending)
	app.Get("/spendings/:userID", API.HandleGetSpendings)

	return app
}

func setConfig() {
	config = Config{
		AppPort:         12345,
		Host:            "http://localhost:12345",
		DBReplicaSetUrl: "mongodb+srv://admin:HkJpLyv1MclTvMIc@spendingtraacker.ybzvy6n.mongodb.net/?retryWrites=true&w=majority",
	}
}

//Also can do like this
/* func createSpendingEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	var spending models.Spending
	json.NewDecoder(request.Body).Decode(&spending)
	collection := client.Database("thespendingtrackerdeneme").Collection("spendings")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := collection.InsertOne(ctx, spending)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}

	json.NewEncoder(response).Encode(result)

}

func getSpendingsEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	var spendings []models.Spending
	collection := client.Database("thespendingtrackerdeneme").Collection("spendings")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var spending models.Spending
		cursor.Decode(&spending)
		spendings = append(spendings, spending)
	}

	if err := cursor.Err(); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}
	json.NewEncoder(response).Encode(spendings)
} */
