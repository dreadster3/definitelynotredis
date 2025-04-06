package concurrentmap

type IConcurrentMap[K comparable, V any] interface {
	Set(key K, value V)
	Get(key K) (V, bool)
	Delete(key K)

	// IterSnapshot Iterates over a snapshot of the current map, more memory usage but no locks
	Iter() func(func(K, V) bool)

	// NextSnapshot Iterates over a snapshot of the current map, more memory usage but no locks
	Next(func(K, V) bool)
}
