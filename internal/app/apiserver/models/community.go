package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Community struct {
	Id           primitive.ObjectID `bson:"_id, omitempty"`
	Name         *string            `bson:"name, omitempty"`
	City         *string            `bson:"city, omitempty"`
	Participants *string            `bson:"participants, omitempty"`
	Type         *string            `bson:"type, omitempty"`
	Admin        *string            `bson:"admin, omitempty"`
}
