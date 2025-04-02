package concurrentmap

import "testing"

func BenchmarkConcurrentMapSetGetDelete(b *testing.B) {
	concurrentMap := newConcurrentMap[string, int]()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			key := "key"
			value := 1
			concurrentMap.Set(key, value)
			concurrentMap.Get(key)
			concurrentMap.Delete(key)
		}
	})
}

func BenchmarkConcurrentMapDifferentKeySet(b *testing.B) {
	concurrentMap := newConcurrentMap[string, int]()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			concurrentMap.Set("key", 1)
		}
	})
}

func BenchmarkConcurrentMapDifferentKeyGet(b *testing.B) {
	concurrentMap := newConcurrentMap[string, int]()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			concurrentMap.Get("key")
		}
	})
}
