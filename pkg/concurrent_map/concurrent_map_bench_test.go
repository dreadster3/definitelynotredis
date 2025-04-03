package concurrentmap

import "testing"

func BenchmarkParallelDelete(b *testing.B) {
	concurrentMap := newConcurrentMap[string, int]()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			key := "key"
			value := 1
			concurrentMap.Set(key, value)
			concurrentMap.Delete(key)
		}
	})
}

func BenchmarkParallelSet(b *testing.B) {
	concurrentMap := newConcurrentMap[string, int]()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			concurrentMap.Set("key", 1)
		}
	})
}

func BenchmarkParallelGet(b *testing.B) {
	concurrentMap := newConcurrentMap[string, int]()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			concurrentMap.Get("key")
		}
	})
}

func BenchmarkSyncSet(b *testing.B) {
	concurrentMap := newConcurrentMap[string, int]()

	for b.Loop() {
		concurrentMap.Set("key", 1)
	}
}

func BenchmarkSyncGet(b *testing.B) {
	concurrentMap := newConcurrentMap[string, int]()
	concurrentMap.Set("key", 1)

	for b.Loop() {
		concurrentMap.Get("key")
	}
}
