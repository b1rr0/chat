package repository

import (
	"chat/domain"
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Message interface {
	Create(ctx context.Context, message domain.Message) (primitive.ObjectID, error)
	FindByTo(ctx context.Context, toID primitive.ObjectID) ([]domain.Message, error)
	FindAll(ctx context.Context) ([]domain.Message, error)
	FindMessagesBetweenUsers(ctx context.Context, fromID, toID primitive.ObjectID) ([]domain.Message, error)
}

type Repositories struct {
	Message Message
}

func NewRepositories(db *mongo.Database) *Repositories {
	return &Repositories{
		Message: NewMessageRepo(db),
	}
}
