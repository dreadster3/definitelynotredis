package engine

import (
	"log/slog"
	"time"

	concurrentmap "github.com/dreadster3/definitelynotredis/pkg/concurrent_map"
)

type cacheEntry struct {
	value     any
	expiresAt time.Time
}

func (entry cacheEntry) isExpired() bool {
	return !entry.expiresAt.IsZero() && time.Now().After(entry.expiresAt)
}

type Cache struct {
	data   concurrentmap.IConcurrentMap[string, cacheEntry]
	ttl    time.Duration
	Logger *slog.Logger
}

func NewCache() *Cache {
	return &Cache{
		data: concurrentmap.NewShardedConcurrentMap[string, cacheEntry](concurrentmap.DefaultStringShardFunc),
		ttl:  time.Minute,
	}
}

func (c *Cache) Get(key string) (any, bool) {
	entry, exists := c.data.Get(key)

	if !exists {
		return nil, false
	}

	if entry.isExpired() {
		return nil, false
	}

	return entry.value, true
}

func (c *Cache) Set(key string, value any) {
	c.SetWithTTL(key, value, c.ttl)
}

func (c *Cache) SetWithTTL(key string, value any, ttl time.Duration) {
	expires := time.Time{}

	if ttl > 0 {
		expires = time.Now().Add(ttl)
	}

	c.data.Set(key, cacheEntry{value: value, expiresAt: expires})
}

func isExpiredFilterFunc(_ string, entry cacheEntry) bool {
	return entry.isExpired()
}

func (c *Cache) removeExpired() {
	for key := range concurrentmap.Filter(c.data.NextSnapshot, isExpiredFilterFunc) {
		c.Logger.Debug("Deleting expired entry", "key", key)
		c.data.Delete(key)
	}
}
