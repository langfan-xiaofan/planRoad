package db

import (
	"backend/internal/config"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var MongoDatabase *mongo.Database

//var MongoConfig = config.Conf.MongoDb

func MongoDbInit() {
	clientOptions := options.Client().ApplyURI(config.Conf.MongoDb.Url)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(clientOptions)
	if err != nil {
		panic(err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		panic(err)
	}
	MongoDatabase = client.Database(config.Conf.MongoDb.Database)
}
