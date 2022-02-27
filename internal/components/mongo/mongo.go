package mongo

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ClientI interface {
	GetCollection(collection string) mongo.Collection
}

type Client struct {
	ctx context.Context
	db  *mongo.Database
}

type Credentials struct {
	ConnectionTimeout time.Duration
	DBName            string
	URI               string
}

func NewClient(ctx context.Context, credentials Credentials) (*Client, error) {
	client := &Client{
		ctx: ctx,
	}
	ctx, cancel := context.WithTimeout(ctx, credentials.ConnectionTimeout)
	defer cancel()
	mClient, err := mongo.NewClient(options.Client().ApplyURI(credentials.URI))

	if err != nil {
		return nil, err
	}

	err = mClient.Connect(ctx)
	if err != nil {
		return nil, err
	}

	client.db = mClient.Database(credentials.DBName)

	return client, nil
}

func (c Client) GetCollection(collection string) mongo.Collection {
	return *c.db.Collection(collection)
}
