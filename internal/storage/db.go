package storage

import (
	"context"
	"testTaskMedods/config"
	"testTaskMedods/pkg"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoDb(config config.Config) (*mongo.Database, error) {
	port := config.MongoUrl

	client, err := mongo.NewClient(options.Client().ApplyURI(port))
	if err != nil {
		pkg.InfoLog.Printf("Error connecting to MongoDB: %v", err)
		return nil, err
	}

	if err = client.Connect(context.TODO()); err != nil {
		pkg.InfoLog.Printf("client MongoDB: %v\n", err)
		return nil, err
	}

	if err = client.Ping(context.TODO(), nil); err != nil {
		pkg.InfoLog.Printf("client connection to MongoDB: %v\n", err)
		return nil, err
	}
	pkg.InfoLog.Println("Connected to MongoDB")
	return client.Database(config.NameDb), nil
}
