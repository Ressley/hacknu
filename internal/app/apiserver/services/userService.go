package services

import (
	"context"
	"time"

	"github.com/Ressley/hacknu/internal/app/apiserver/helpers"
	"github.com/Ressley/hacknu/internal/app/apiserver/models"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var client, err = helpers.GetMongoClient()
var userCollection *mongo.Collection = client.Database(helpers.DB).Collection(helpers.USERS)

func GetMongoClient() (*mongo.Client, error) {
	return client, err
}

func CreateUser(account *models.Account) error {
	var ctx, _ = context.WithTimeout(context.TODO(), 100*time.Second)
	_, err = GetUserOneByNumber(account.Login)
	if err == nil {
		return errors.New("account allready exist")
	}
	user := models.User{
		First_name: account.First_name,
		Last_name:  account.Last_name,
		Number:     account.Login,
		User_id:    account.User_id,
	}
	user.ID = primitive.NewObjectID()
	_, err = userCollection.InsertOne(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

func GetUserOneByID(id *primitive.ObjectID) (models.User, error) {
	var ctx, _ = context.WithTimeout(context.TODO(), 100*time.Second)
	result := models.User{}
	filter := bson.D{{Key: "_id", Value: id}}
	err := userCollection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return result, err
	}
	return result, nil
}

func GetUserOneByUserID(uid *string) (models.User, error) {
	var ctx, _ = context.WithTimeout(context.TODO(), 100*time.Second)
	result := models.User{}
	filter := bson.D{{Key: "user_id", Value: uid}}
	err := userCollection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return result, err
	}
	return result, nil
}

func GetUserOneByNumber(number *string) (models.User, error) {
	var ctx, _ = context.WithTimeout(context.TODO(), 100*time.Second)
	result := models.User{}
	filter := bson.D{{Key: "number", Value: number}}

	err = userCollection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return result, err
	}
	return result, nil

}

func UpdateUserOne(user *models.User) error {
	var ctx, _ = context.WithTimeout(context.TODO(), 100*time.Second)

	filter := bson.D{{Key: "user_id", Value: user.User_id}}
	updater := bson.D{primitive.E{Key: "$set", Value: bson.D{
		primitive.E{Key: "first_name", Value: user.First_name},
		primitive.E{Key: "last_name", Value: user.Last_name},
		primitive.E{Key: "number", Value: user.Number},
		primitive.E{Key: "photo", Value: user.Photo},
	}}}

	_, err = userCollection.UpdateOne(ctx, filter, updater)
	if err != nil {
		return err
	}
	return nil
}

func DeleteUserOne(id *primitive.ObjectID) error {
	var ctx, _ = context.WithTimeout(context.TODO(), 100*time.Second)

	filter := bson.D{{Key: "_id", Value: id}}
	_, err := userCollection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	return nil

}
