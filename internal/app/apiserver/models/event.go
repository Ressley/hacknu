package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Event struct {
	ID           primitive.ObjectID `bson:"_id, omitempty"`
	Type         []string           `bson:"type, omitempty"`
	Participants *string            `bson:"participants, omitempty"`
	Admin        *string            `bson:"admin, omitempty"`
}
