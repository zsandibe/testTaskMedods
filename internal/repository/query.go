package repository

import (
	"context"
	"errors"
	"fmt"
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
	fmt.Println(session.HashedRefreshToken, "after")
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

func (r *repositoryMongo) GetAllSessions() ([]domain.Session, error) {
	ctx := context.Background()
	rows, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, errors.New("Can`t get all sessions rows")
	}
	defer rows.Close(ctx)

	var sessions []domain.Session

	for rows.Next(ctx) {
		var session domain.Session
		if err := rows.Decode(&session); err != nil {
			fmt.Println("Error in getting")
			return nil, errors.New("Failed to decode row")
		}
		sessions = append(sessions, session)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	fmt.Println("sessions:", sessions)
	return sessions, nil
}
