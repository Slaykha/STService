package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Slaykha/STService/models"
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

func (r *Repository) GetSpendings() ([]models.Spending, error) {
	collection := r.client.Database("spending").Collection("spendings")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var spendings []models.Spending

	emptyFilter := bson.M{}

	result, err := collection.Find(ctx, emptyFilter)
	if err != nil {
		fmt.Println("2")
		return nil, err
	}

	for result.Next(ctx) {
		spending := models.Spending{}
		err := result.Decode(&spending)
		if err != nil {
			fmt.Println("3")
			return nil, err
		}

		spendings = append(spendings, spending)
	}

	if err != nil {
		fmt.Println("4")
		return nil, err
	}

	return spendings, nil
}
