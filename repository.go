package main

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/Slaykha/STService/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (r *Repository) GetUserInfo(userId string) (*models.UserAuth, error) {
	collection := r.client.Database("spending").Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"id": userId}

	result := collection.FindOne(ctx, filter)

	user := models.UserAuth{}
	err := result.Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *Repository) GetUser(userId string) (*models.User, error) {
	collection := r.client.Database("spending").Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"id": userId}

	result := collection.FindOne(ctx, filter)

	user := models.User{}
	err := result.Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *Repository) UpdateUser(userModel models.User) (*models.UserAuth, error) {
	collection := r.client.Database("spending").Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"id": userModel.ID}

	collection.FindOneAndReplace(ctx, filter, userModel)

	updatedUser, err := r.GetUserInfo(userModel.ID)
	if err != nil {
		return nil, err
	}

	return updatedUser, nil
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

func (r *Repository) GetSpendings(userID, spendingType, moneySort, dateSort string, date time.Time) ([]models.Spending, error) {
	collection := r.client.Database("spending").Collection("spendings")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var spendings []models.Spending

	filter := bson.M{
		"$and": []bson.M{
			{"userId": userID},
			{"spendingDate": bson.M{
				"$gte": primitive.NewDateTimeFromTime(time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, time.Now().Location())),
			}},
			{"spendingType": bson.M{"$regex": primitive.Regex{Pattern: spendingType, Options: "i"}}},
		},
	}

	options := options.Find()
	if dateSort != "" {
		options.SetSort(bson.D{{"spendingDate", 1}})
	} else {
		options.SetSort(bson.D{{"spendingDate", -1}})
	}

	if moneySort != "" {
		sort, _ := strconv.Atoi(moneySort)

		options.SetSort(bson.D{{"money", sort}})
	}

	result, err := collection.Find(ctx, filter, options)
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

func (r *Repository) DeleteSpending(spendingID string) error {
	collection := r.client.Database("spending").Collection("spendings")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"id": spendingID}

	result := collection.FindOneAndDelete(ctx, filter)
	if result != nil {
		return result.Err()
	}

	return nil
}

func (r *Repository) GetSpending(spendingId string) (*models.Spending, error) {
	collection := r.client.Database("spending").Collection("spendings")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"id": spendingId}

	result := collection.FindOne(ctx, filter)

	spending := models.Spending{}
	err := result.Decode(&spending)
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	return &spending, nil
}

func (r *Repository) UpdateSpending(spendingModel models.Spending) (*models.Spending, error) {
	collection := r.client.Database("spending").Collection("spendings")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"id": spendingModel.ID}

	collection.FindOneAndReplace(ctx, filter, spendingModel)

	updatedUser, err := r.GetSpending(spendingModel.ID)
	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}
