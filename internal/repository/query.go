package repository

import (
	"context"
	"errors"
	"testTaskMedods/internal/domain"
	"testTaskMedods/pkg"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

func (r *repositoryMongo) Create(session domain.Session) error {
	_, err := r.collection.InsertOne(context.Background(), session)
	if err != nil {
		pkg.ErrorLog.Printf("Error in creating: %v", err)
		return err
	}
	pkg.InfoLog.Println("Successfully created token pairs")
	return nil
}

func (r *repositoryMongo) GetSessionById(sessionId uuid.UUID) (domain.Session, error) {
	var session domain.Session

	filter := bson.D{{"id", sessionId}}

	if err := r.collection.FindOne(context.Background(), filter).Decode(&session); err != nil {
		pkg.ErrorLog.Printf("Error in getting session document: %v", err)
		return session, errors.New("Error no documents")
	}
	return session, nil
}

func (r *repositoryMongo) DeleteSessionById(sessionId uuid.UUID) error {
	filter := bson.D{{"id", sessionId}}

	if _, err := r.collection.DeleteOne(context.Background(), filter); err != nil {
		pkg.ErrorLog.Printf("Error in deleting session document: %v", err)
		return errors.New("Can`t delete session document")
	}
	return nil
}

func (r *repositoryMongo) Update(session domain.Session) error {
	filter := bson.D{{"id", session.Id}}
	update := bson.D{{"$set", bson.D{
		{"hashed_refresh_token", session.HashedRefreshToken},
		{"updated_at", session.UpdatedAt},
	}}}

	if _, err := r.collection.UpdateOne(context.Background(), filter, update); err != nil {
		pkg.ErrorLog.Printf("Error in updating session document: %v", err)
		return errors.New("Can`t update session document")
	}
	return nil
}
