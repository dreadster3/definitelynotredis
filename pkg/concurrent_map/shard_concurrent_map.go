package concurrentmap

type ShardedConcurrentMapOptions[K comparable, V any] func(*ShardedConcurrentMap[K, V])

// WithShards sets the number of shards. Recommended to keep this value a power of 2 for faster modulo operations
func WithShards[K comparable, V any](shards uint32) ShardedConcurrentMapOptions[K, V] {
	return func(m *ShardedConcurrentMap[K, V]) {
		m.shards = make([]*ConcurrentMap[K, V], shards)
		m.size = shards
	}
}

type ShardFunc[K comparable] func(K) uint32

func DefaultStringShardFunc(key string) uint32 {
	hash := uint32(2166136261)
	for _, c := range key {
		hash ^= uint32(c)
		hash *= 16777619
	}
	return hash
}

func DefaultIntShardFunc(key int) uint32 {
	return uint32(key)
}

type ShardedConcurrentMap[K comparable, V any] struct {
	shards    []*ConcurrentMap[K, V]
	shardFunc ShardFunc[K]
	size      uint32
}

func newShardedConcurrentMap[K comparable, V any](shardFunc ShardFunc[K], opts ...ShardedConcurrentMapOptions[K, V]) *ShardedConcurrentMap[K, V] {
	result := &ShardedConcurrentMap[K, V]{
		shardFunc: shardFunc,
	}

	for _, option := range opts {
		option(result)
	}

	if result.shards == nil {
		result.shards = make([]*ConcurrentMap[K, V], 16)
		result.size = 16
	}

	for idx := range result.shards {
		result.shards[idx] = newConcurrentMap[K, V]()
	}

	return result
}

func NewShardedConcurrentMap[K comparable, V any](shardFunc ShardFunc[K], opts ...ShardedConcurrentMapOptions[K, V]) IConcurrentMap[K, V] {
	return newShardedConcurrentMap(shardFunc, opts...)
}

func (m *ShardedConcurrentMap[K, V]) Get(key K) (V, bool) {
	shard := m.shards[m.shardFunc(key)%m.size]
	return shard.Get(key)
}

func (m *ShardedConcurrentMap[K, V]) Set(key K, value V) {
	shard := m.shards[m.shardFunc(key)%m.size]
	shard.Set(key, value)
}

func (m *ShardedConcurrentMap[K, V]) Delete(key K) {
	shard := m.shards[m.shardFunc(key)%m.size]
	shard.Delete(key)
}
