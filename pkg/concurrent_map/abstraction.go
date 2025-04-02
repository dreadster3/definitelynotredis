package concurrentmap

type IConcurrentMap[K comparable, V any] interface {
	Set(key K, value V)
	Get(key K) (V, bool)
	Delete(key K)
}
