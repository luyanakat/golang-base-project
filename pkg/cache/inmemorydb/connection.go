package inmemorydb

import (
	"time"

	"github.com/dgraph-io/ristretto"
)

type InmemmoryDB struct {
	cache *ristretto.Cache
}

func NewInmemoryDB() (*InmemmoryDB, error) {
	cache, err := ristretto.NewCache(&ristretto.Config{
		NumCounters: 1e7,     // number of keys to track frequency of (10M).
		MaxCost:     2 << 30, // maximum cost of cache (1GB).
		BufferItems: 64,      // number of keys per Get buffer.
	})
	if err != nil {
		return nil, err
	}

	return &InmemmoryDB{
		cache: cache,
	}, nil
}

func (m *InmemmoryDB) SetTTL(key string, value any, ttl time.Duration) bool {
	ok := m.cache.SetWithTTL(key, value, 1, ttl)
	m.cache.Wait()
	return ok
}

func (m *InmemmoryDB) Get(key string) (any, bool) {
	value, ok := m.cache.Get(key)
	return value, ok
}

func (m *InmemmoryDB) Delete(key string) {
	m.cache.Del(key)
}

func (m *InmemmoryDB) Close() {
	m.cache.Close()
}
