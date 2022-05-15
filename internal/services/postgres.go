package services

import (
	"L0/internal/repository"
)

type SQLService struct {
	*repository.SQLRepository
}

func NewSQLService(SQLrepo *repository.SQLRepository) *SQLService {
	return &SQLService{SQLRepository: SQLrepo}
}

func (s SQLService) Save(id string, body []byte) error {
	return s.SQLRepository.Save(id, body)
}
