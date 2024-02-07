package repository

import (
	"testTaskMedods/internal/domain"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	Create(session domain.Session) error
	GetSessionById(sessionId uuid.UUID) (domain.Session, error)
	DeleteSessionById(sessionId uuid.UUID) error
	Update(session domain.Session) error
}

type repositoryMongo struct {
	collection *mongo.Collection
}

func NewRepository(db *mongo.Database) *repositoryMongo {
	return &repositoryMongo{
		collection: db.Collection("session"),
	}
}
