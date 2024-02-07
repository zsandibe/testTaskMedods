package service

import (
	"testTaskMedods/internal/domain"
	"testTaskMedods/internal/repository"
)

type Service interface {
	Create(session domain.Session) error
	Update(session domain.Session) error
}

type service struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) *service {
	return &service{
		repo: repo,
	}
}
