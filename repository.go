package main

import (
	"context"
	"log"
	"time"

	"github.com/Slaykha/STService/models"
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository struct {
	client *mongo.Client
}

func NewRepository(dbReplicaSetUrl string) *Repository {
	uri := dbReplicaSetUrl
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}

	return &Repository{client}
}

func (r *Repository) CreateUser(user models.User) error {
	collection := r.client.Database("spending").Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, user)

	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) FindUser(email string) (*models.User, error) {
	collection := r.client.Database("spending").Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"email": email}

	result := collection.FindOne(ctx, filter)

	user := models.User{}
	err := result.Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, err
}

func (r *Repository) GetUser(claims *jwt.StandardClaims) *models.UserAuth {
	collection := r.client.Database("spending").Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"_id": claims.Issuer}

	result := collection.FindOne(ctx, filter)

	user := models.UserAuth{}
	result.Decode(&user)

	return &user
}

func (r *Repository) CreateSpending(spending models.Spending) error {
	collection := r.client.Database("spending").Collection("spendings")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, spending)

	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) GetSpendings(userID string) ([]models.Spending, error) {
	collection := r.client.Database("spending").Collection("spendings")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var spendings []models.Spending

	filter := bson.M{"userId": userID}

	result, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	for result.Next(ctx) {
		spending := models.Spending{}
		err := result.Decode(&spending)
		if err != nil {
			return nil, err
		}

		spendings = append(spendings, spending)
	}

	if err != nil {
		return nil, err
	}

	return spendings, nil
}
