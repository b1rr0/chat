package services

import (
	"chat/domain"
	"chat/repository"
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MessageService struct {
	repo repository.Message
}

func NewMessageService(repo repository.Message) *MessageService {
	return &MessageService{repo: repo}
}

func (s *MessageService) Create(ctx context.Context, message domain.Message) (primitive.ObjectID, error) {
	s.repo.Create(ctx, message)
	return primitive.NewObjectID(), nil
}
func (s *MessageService) FindAll(ctx context.Context) ([]domain.Message, error) {
	return s.repo.FindAll(ctx)
}
func (s *MessageService) FindByTo(ctx context.Context, id primitive.ObjectID) ([]domain.Message, error) {
	return s.repo.FindByTo(ctx, id)
}
func (s *MessageService) FindMessagesBetweenUsers(ctx context.Context, fromID, toID primitive.ObjectID) ([]domain.Message, error) {
	return s.repo.FindMessagesBetweenUsers(ctx, fromID, toID)
}
