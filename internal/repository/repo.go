package repository

type Repository struct {
	PgPoolRepo   *SQLRepository
	NatsConnRepo *NatsRepository
	CacheRepo    *Cache
}

func New() *Repository {
	pgPool := NewSQLRepo()
	NatsConn := NewNatsRepo()
	Cache := NewCacheRepo()

	return &Repository{pgPool, NatsConn, Cache}
}
