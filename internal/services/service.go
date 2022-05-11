package services

import "L0/internal/repository"

type Service struct {
	*repository.Repository
}

func New(repo *repository.Repository) *Service {
	return &Service{repo}
}
