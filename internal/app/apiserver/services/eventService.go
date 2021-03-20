package services

import (
	"context"
	"time"

	"github.com/Ressley/hacknu/internal/app/apiserver/helpers"
	"github.com/Ressley/hacknu/internal/app/apiserver/models"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var eventCollection *mongo.Collection = client.Database(helpers.DB).Collection(helpers.EVENT)

func CreateEvent(event *models.Event) error {
	var ctx, _ = context.WithTimeout(context.TODO(), 100*time.Second)

	_, err = GetEventByName(event.Name, event.Admin)
	if err == nil {
		return errors.New("community allready exist")
	}

	_, err = eventCollection.InsertOne(ctx, event)
	if err != nil {
		return err
	}
	return nil
}

func GetEventAll() ([]models.Event, error) {
	var ctx, _ = context.WithTimeout(context.TODO(), 100*time.Second)
	result := []models.Event{}
	filter := bson.D{}
	cursor, err := eventCollection.Find(ctx, filter)

	err = cursor.All(ctx, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}

func GetEventByName(name *string, admin *string) (models.Community, error) {
	var ctx, _ = context.WithTimeout(context.TODO(), 100*time.Second)
	result := models.Community{}
	filter := bson.D{{Key: "name", Value: name}, {Key: "admin", Value: admin}}

	err = eventCollection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return result, err
	}
	return result, nil

}
