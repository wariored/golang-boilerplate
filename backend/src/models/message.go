package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Message struct {
	ID        primitive.ObjectID `bson:"_id"`
	Text      string             `bson:"text"`
	SenderID    string             `bson:"sender"`
	RecipientID    string             `bson:"sender"`
	Timestamp time.Time          `bson:"timestamp"`
}

