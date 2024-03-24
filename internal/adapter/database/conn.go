package database

import (
	"context"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var mongoClient *mongo.Client

func GetConnection(ctx context.Context) *mongo.Client {
	if mongoClient != nil {
		return mongoClient
	}
	uri := viper.GetString("MONGO_URI")
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().
		ApplyURI(uri).
		SetServerAPIOptions(serverAPIOptions)
	var err error
	mongoClient, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Panic(err)
	}
	return mongoClient
}
