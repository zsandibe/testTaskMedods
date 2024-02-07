package service

import (
	"testTaskMedods/config"
	"testTaskMedods/internal/domain"
	"testTaskMedods/internal/repository"
)

type Service interface {
	Create(session domain.Session) error
	Update(session domain.Session) error
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
