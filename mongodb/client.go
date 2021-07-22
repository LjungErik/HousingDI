package mongodb

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Config struct {
	ConnString string
}

type Client struct {
	connOptions *options.ClientOptions
	ctx         context.Context
	cancel      context.CancelFunc
}

func NewClient(conf *Config) *Client {
	clientOptions := options.Client().ApplyURI(conf.ConnString)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		log.Fatal("Failed to acquire connection to MongoDB server")
		panic(err)
	}

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Error("Failed to ping MongoDB")
		panic(err)
	}

	return &Client{
		connOptions: clientOptions,
		ctx:         ctx,
		cancel:      cancel,
	}
}

func (c *Client) InsertOrUpdateHousingForSale(doc HousingForSaleDoc) error {
	return nil
}

func (c *Client) InsertOrUpdateHousingSold(doc HousingSoldDoc) error {
	return nil
}
