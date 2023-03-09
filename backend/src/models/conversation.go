package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Conversation struct {
	ID        primitive.ObjectID `bson:"_id"`
	Users     []string           `bson:"users"`
	Messages  []Message          `bson:"messages"`
	Timestamp time.Time          `bson:"timestamp"`
}

