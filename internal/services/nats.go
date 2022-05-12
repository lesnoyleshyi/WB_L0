package services

import "L0/internal/repository"

type NatsService struct {
	*repository.NatsRepository
}

func NewNatsService(natsRepo *repository.NatsRepository) *NatsService {
	return &NatsService{NatsRepository: natsRepo}
}
