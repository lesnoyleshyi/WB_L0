package services

import (
	"L0/internal/domain"
	"L0/internal/repository"
)

type CacheService struct {
	*repository.Cache
}

func NewCacheService(repoCache *repository.Cache) *CacheService {
	return &CacheService{Cache: repoCache}
}

func (cs CacheService) Save(order *domain.Order) {
	cs.Cache.Save(order)
}

func (cs CacheService) Get(orderId string) (*domain.Order, bool) {
	return cs.Cache.Get(orderId)
}
