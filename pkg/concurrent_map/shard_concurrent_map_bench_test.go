package concurrentmap

import (
	"math/rand/v2"
	"testing"
)

func RandomShardFunc(key string) uint32 {
	return rand.Uint32()
}

func BenchmarkParallelShardSameDelete(b *testing.B) {
	concurrentMap := newShardedConcurrentMap[string, int](DefaultStringShardFunc)

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			key := "key"
			value := 1
			concurrentMap.Set(key, value)
			concurrentMap.Delete(key)
		}
	})
}

func BenchmarkParallelShardRandomDelete(b *testing.B) {
	concurrentMap := newShardedConcurrentMap[string, int](RandomShardFunc)

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			key := "key"
			value := 1
			concurrentMap.Set(key, value)
			concurrentMap.Delete(key)
		}
	})
}

func BenchmarkParallelShardSameSet(b *testing.B) {
	concurrentMap := newShardedConcurrentMap[string, int](DefaultStringShardFunc)

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			concurrentMap.Set("key", 1)
		}
	})
}

func BenchmarkParallelShardRandomSet(b *testing.B) {
	concurrentMap := newShardedConcurrentMap[string, int](RandomShardFunc)

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			concurrentMap.Set("key", 1)
		}
	})
}

func BenchmarkParallelShardSameGet(b *testing.B) {
	concurrentMap := newShardedConcurrentMap[string, int](DefaultStringShardFunc)

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			concurrentMap.Get("key")
		}
	})
}

func BenchmarkParallelShardRandomGet(b *testing.B) {
	concurrentMap := newShardedConcurrentMap[string, int](RandomShardFunc)

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			concurrentMap.Get("key")
		}
	})
}

func BenchmarkSyncShardSet(b *testing.B) {
	concurrentMap := newShardedConcurrentMap[string, int](DefaultStringShardFunc)

	for b.Loop() {
		concurrentMap.Set("key", 1)
	}
}

func BenchmarkSyncShardRandomSet(b *testing.B) {
	concurrentMap := newShardedConcurrentMap[string, int](RandomShardFunc)

	for b.Loop() {
		concurrentMap.Set("key", 1)
	}
}

func BenchmarkSyncShardGet(b *testing.B) {
	concurrentMap := newShardedConcurrentMap[string, int](DefaultStringShardFunc)
	concurrentMap.Set("key", 1)

	for b.Loop() {
		concurrentMap.Get("key")
	}
}

func BenchmarkSyncShardRandomGet(b *testing.B) {
	concurrentMap := newShardedConcurrentMap[string, int](RandomShardFunc)
	concurrentMap.Set("key", 1)

	for b.Loop() {
		concurrentMap.Get("key")
	}
}
