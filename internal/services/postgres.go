package services

import "L0/internal/repository"

type SQLService struct {
	*repository.SQLRepository
}

func NewSQLService(SQLrepo *repository.SQLRepository) *SQLService {
	return &SQLService{SQLRepository: SQLrepo}
}
