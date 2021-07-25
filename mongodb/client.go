package mongodb

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Config struct {
	ConnString   string
	DatabaseName string
}

type Client struct {
	session *mongo.Client
	db      *mongo.Database
}

func NewClient(conf *Config) *Client {
	clientOptions := options.Client().ApplyURI(conf.ConnString)
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	session, err := mongo.NewClient(clientOptions)

	if err != nil {
		log.Fatal("Failed to create MongoDB client")
		panic(err)
	}

	err = session.Connect(ctx)
	if err != nil {
		log.Fatal("Failed to establish connection to MongoDB server")
		panic(err)
	}

	err = session.Ping(ctx, nil)
	if err != nil {
		log.Error("Failed to ping MongoDB")
		panic(err)
	}

	db := session.Database(conf.DatabaseName)
	log.Infof("Successfully connected to database: %v", conf.DatabaseName)

	return &Client{
		session: session,
		db:      db,
	}
}

func (c *Client) Disconnect() {

}

func (c *Client) InsertOrUpdateHousingForSale(doc HousingForSaleDoc) error {
	coll := c.db.Collection("housingforsale")

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	opts := options.Update().SetUpsert(true)
	filter := bson.D{primitive.E{
		Key:   "propId",
		Value: doc.PropId,
	}}
	update := bson.M{
		"$set": doc,
	}
	// Fire and forget until further
	_, err := coll.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		log.Warnf("Failed to update/insert housing 'for sale' document for: %v", doc.PropId)
		return err
	}

	return nil
}

func (c *Client) InsertOrUpdateHousingSold(doc HousingSoldDoc) error {
	coll := c.db.Collection("housingsold")

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	opts := options.Update().SetUpsert(true)
	filter := bson.D{primitive.E{
		Key:   "propId",
		Value: doc.PropId,
	}}
	update := bson.M{
		"$set": doc,
	}
	// Fire and forget until further
	_, err := coll.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		log.Warnf("Failed to update/insert housing 'sold' document for: %v", doc.PropId)
		return err
	}

	return nil
}
