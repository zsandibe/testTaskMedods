package repository

import "go.mongodb.org/mongo-driver/mongo"

type Repository interface {
}

type repositoryMongo struct {
	db *mongo.Database
}

func NewRepository(db *mongo.Database) *repositoryMongo {
	return &repositoryMongo{
		db: db,
	}
}
