package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Community struct {
	ID           primitive.ObjectID `bson:"_id, omitempty"`
	Photo        *string            `bson:"photo, omitempty"`
	Name         *string            `bson:"name, omitempty"`
	City         *string            `bson:"city, omitempty"`
	Participants []string           `bson:"participants, omitempty"`
	Admin        *string            `bson:"admin, omitempty"`
}
