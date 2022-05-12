package repository

import (
	"L0/internal/domain"
	"sync"
)

//Cache --is it ok to store pointers to structs(for example, from GC point of view)?
type Cache struct {
	mu      sync.RWMutex
	Storage map[string]*domain.Order
}

//NewCacheRepo is redundant
func NewCacheRepo() *Cache {
	storage := make(map[string]*domain.Order)
	return &Cache{Storage: storage}
}

func (c *Cache) Save(order *domain.Order) {
	c.mu.Lock()
	c.Storage[order.Id] = order
	c.mu.Unlock()
}

func (c *Cache) Get(id string) (*domain.Order, bool) {
	order, ok := c.Storage[id]
	return order, ok
}
