package service

import (
	"testTaskMedods/config"
	"testTaskMedods/internal/domain"
	"testTaskMedods/internal/repository"

	"github.com/google/uuid"
)

type Service interface {
	Create(guid uuid.UUID) (domain.TokenPair, error)
	Update(sessionId uuid.UUID, refreshToken []byte) (domain.TokenPair, error)
}

type service struct {
	conf config.Config
	repo repository.Repository
}

func NewService(repo repository.Repository, conf config.Config) *service {
	return &service{
		repo: repo,
		conf: conf,
	}
}
