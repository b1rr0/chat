package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Message struct {
	ID   primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	FROM primitive.ObjectID `json:"from" bson:"from,omitempty"`
	TO   primitive.ObjectID `json:"TO" bson:"to,omitempty"`
	DATA string             `json:"name" bson:"name"`
}
