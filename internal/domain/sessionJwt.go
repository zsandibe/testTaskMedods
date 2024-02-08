package domain

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type Session struct {
	Id                 uuid.UUID `bson:"id" json:"id"`
	Guid               uuid.UUID `bson:"guid" json:"guid"`
	HashedRefreshToken []byte    `bson:"hashed_refresh_token" json:"hashed_refresh_token"`
	CreatedAt          time.Time `bson:"created_at" json:"created_time"`
	UpdatedAt          time.Time `bson:"updated_at" json:"updated_time"`
}

type TokenPair struct {
	AccessToken  string `json:"access_token"`
	RefreshToken []byte `json:"refresh_token"`
}

type AccessTokenClaims struct {
	jwt.RegisteredClaims
	SessionId uuid.UUID `json:"session_id"`
	Guid      uuid.UUID `json:"guid"`
}
