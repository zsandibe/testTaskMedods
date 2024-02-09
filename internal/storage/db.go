package storage

import (
	"context"
	"fmt"
	"testTaskMedods/config"
	"testTaskMedods/pkg"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoDb(config config.Config) (*mongo.Database, error) {
	url := config.Database.MongoUrl
	fmt.Println(url)
	client, err := mongo.NewClient(options.Client().ApplyURI(url))
	if err != nil {
		pkg.ErrorLog.Printf("Error connecting to MongoDB: %v", err)
		return nil, err
	}

	if err = client.Connect(context.TODO()); err != nil {
		pkg.ErrorLog.Printf("client MongoDB: %v\n", err)
		return nil, err
	}

	if err = client.Ping(context.TODO(), nil); err != nil {
		pkg.ErrorLog.Printf("client connection to MongoDB: %v\n", err)
		return nil, err
	}
	pkg.InfoLog.Println("Connected to MongoDB")
	return client.Database(config.Database.NameDb), nil
}
