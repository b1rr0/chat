package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type FromToModel struct {
	From primitive.ObjectID
	To   primitive.ObjectID
}
