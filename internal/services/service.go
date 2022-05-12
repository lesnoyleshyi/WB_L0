package services

import "L0/internal/repository"

type Service struct {
	SQLService   *SQLService
	NatsService  *NatsService
	CacheService *CacheService
}

func New(repo *repository.Repository) *Service {
	SQLSrv := NewSQLService(repo.PgPoolRepo)
	NatsSrv := NewNatsService(repo.NatsConnRepo)
	CacheSrv := NewCacheService(repo.CacheRepo)

	return &Service{
		SQLService:   SQLSrv,
		NatsService:  NatsSrv,
		CacheService: CacheSrv,
	}
}
