package domain

import "github.com/google/uuid"

type AuthRequest struct {
	Guid uuid.UUID `json:"guid"`
}

type RefreshRequest struct {
	SessionID    uuid.UUID `json:"session_id"`
	RefreshToken []byte    `json:"refresh_token"`
}
