package repository

import (
	"chat/domain"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MessageRepo struct {
	db *mongo.Collection
}

func NewMessageRepo(db *mongo.Database) *MessageRepo {
	return &MessageRepo{
		db: db.Collection("message"),
	}
}

func (r *MessageRepo) Create(ctx context.Context, message domain.Message) (primitive.ObjectID, error) {
	message.ID = primitive.NewObjectID()
	res, err := r.db.InsertOne(ctx, message)
	if err != nil {
		return primitive.ObjectID{}, fmt.Errorf("Some Error: %v", err)
	}
	return res.InsertedID.(primitive.ObjectID), nil
}

func (r *MessageRepo) FindByTo(ctx context.Context, toID primitive.ObjectID) ([]domain.Message, error) {
	filter := bson.M{"to": toID}

	cursor, err := r.db.Find(ctx, filter)

	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var result []domain.Message
	for cursor.Next(ctx) {
		var message domain.Message
		if err := cursor.Decode(&message); err != nil {
			return nil, err
		}
		result = append(result, message)
	}
	return result, nil
}

func (r *MessageRepo) FindAll(ctx context.Context) ([]domain.Message, error) {
	filter := bson.M{}

	cursor, err := r.db.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var result []domain.Message
	for cursor.Next(context.TODO()) {
		var message domain.Message
		if err := cursor.Decode(&message); err != nil {
			return nil, err
		}
		result = append(result, message)
	}

	return result, nil
}

func (r *MessageRepo) FindMessagesBetweenUsers(ctx context.Context, fromID, toID primitive.ObjectID) ([]domain.Message, error) {
	filter := bson.M{
		"$or": []bson.M{
			{"from": fromID, "to": toID},
			{"from": toID, "to": fromID},
		},
	}

	cursor, err := r.db.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var result []domain.Message
	for cursor.Next(ctx) {
		var message domain.Message
		if err := cursor.Decode(&message); err != nil {
			return nil, err
		}
		result = append(result, message)
	}

	return result, nil
}
