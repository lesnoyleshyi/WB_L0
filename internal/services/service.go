package services

import "L0/internal/repository"

type Service struct {
	SQLService   *SQLService
	CacheService *CacheService
}

func New(repo *repository.Repository) *Service {
	SQLSrv := NewSQLService(repo.PgPoolRepo)
	CacheSrv := NewCacheService(repo.CacheRepo)

	return &Service{
		SQLService:   SQLSrv,
		CacheService: CacheSrv,
	}
}
