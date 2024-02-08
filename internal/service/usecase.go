package service

import (
	"errors"
	"testTaskMedods/internal/domain"
	"testTaskMedods/pkg"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

func (s *service) Create(guid uuid.UUID) (domain.TokenPair, error) {
	session := domain.Session{
		Id:        uuid.New(),
		Guid:      guid,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	tokenPair, err := s.signTokenPair(session.Id, session.Guid)
	if err != nil {
		return domain.TokenPair{}, err
	}

	session.HashedRefreshToken, err = pkg.GetHashFromToken(tokenPair.RefreshToken)
	if err != nil {
		return domain.TokenPair{}, err
	}

	if err = s.repo.Create(session); err != nil {
		return domain.TokenPair{}, err
	}
	return tokenPair, nil
}

func (s *service) Update(session domain.Session) error {
	return nil
}

func (s *service) signTokenPair(sessionId uuid.UUID, guid uuid.UUID) (domain.TokenPair, error) {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS512, domain.AccessTokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(s.conf.AccessTokenAge)),
		},
		SessionId: sessionId,
		Guid:      guid,
	})

	signedAccessToken, err := accessToken.SignedString(s.conf.AccessKey)
	if err != nil {
		return domain.TokenPair{}, errors.New("Failed to signed")
	}
	refreshToken := uuid.New()

	tokenPair := domain.TokenPair{
		AccessToken:  signedAccessToken,
		RefreshToken: refreshToken[:],
	}
	return tokenPair, nil
}
