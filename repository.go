package main

import (
	"context"
	"log"
	"time"

	"github.com/Slaykha/STService/models"
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
