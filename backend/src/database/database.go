package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Client is a MongoDB client
type Client struct {
	*mongo.Client
}

// NewClient creates a new MongoDB client
func NewClient() (*Client, error) {
	// set client options
	clientOptions := options.Client().ApplyURI("mongodb://root:example@mongo:27017")

	// connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}

	// check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}

	log.Println("Connected to MongoDB!")
	return &Client{client}, nil
}

// Disconnect closes the MongoDB connection
func (c *Client) Disconnect(ctx context.Context) error {
	return c.Client.Disconnect(ctx)
}

