package service

import "testTaskMedods/internal/repository"

type Service interface {
}

type service struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) *service {
	return &service{
		repo: repo,
	}
}
