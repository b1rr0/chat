package services

import (
	"chat/domain"
	"chat/repository"
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Message interface {
	Create(ctx context.Context, message domain.Message) (primitive.ObjectID, error)
	FindByTo(ctx context.Context, id primitive.ObjectID) ([]domain.Message, error)
	FindMessagesBetweenUsers(ctx context.Context, fromID, toID primitive.ObjectID) ([]domain.Message, error)
	FindAll(ctx context.Context) ([]domain.Message, error)
}
type Services struct {
	Message Message
}

type Deps struct {
	Repos *repository.Repositories
}

func NewServices(deps Deps) *Services {
	messageService := NewMessageService(deps.Repos.Message)

	return &Services{Message: messageService}
}
