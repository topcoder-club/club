package cache

import (
	"github.com/dgraph-io/ristretto"
	"gohub/pkg/config"
	"time"
)

// RistrettoStore 实现 cache.Store interface
type RistrettoStore struct {
	Cache     *ristretto.Cache
	KeyPrefix string
}

func NewRistrettoStore() *RistrettoStore {
	rs := &RistrettoStore{}
	cache, err := ristretto.NewCache(&ristretto.Config{
		NumCounters: 1e7,     // number of keys to track frequency of (10M).
		MaxCost:     1 << 30, // maximum cost of cache (1GB).
		BufferItems: 64,      // number of keys per Get buffer.
	})
	if err != nil {
		panic(err)
	}
	rs.Cache = cache
	rs.KeyPrefix = config.GetString("app.name") + ":cache:"
	return rs
}

func (s *RistrettoStore) Set(key string, value string, expireTime time.Duration) {
	s.Cache.SetWithTTL(s.KeyPrefix+key, value, 0, expireTime)
}

func (s *RistrettoStore) Get(key string) string {
	if val, ok := s.Cache.Get(s.KeyPrefix + key); ok {
		return val.(string)
	}
	return ""
}

func (s *RistrettoStore) Has(key string) bool {
	_, ok := s.Cache.Get(s.KeyPrefix + key)
	return ok
}

func (s *RistrettoStore) Forget(key string) {
	s.Cache.Del(s.KeyPrefix + key)
}

func (s *RistrettoStore) Forever(key string, value string) {
	s.Cache.Set(s.KeyPrefix+key, value, 0)
}

func (s *RistrettoStore) Flush() {
	s.Cache.Wait()
}

func (s *RistrettoStore) Increment(parameters ...interface{}) {
	panic("todo")
}

func (s *RistrettoStore) Decrement(parameters ...interface{}) {
	panic("todo")
}

func (s *RistrettoStore) IsAlive() error {
	return nil
}
