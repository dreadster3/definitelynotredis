package concurrentmap

import (
	"fmt"
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

type bigStruct struct {
	test1  int
	test2  int
	test3  int
	test4  int
	test5  int
	test6  int
	test7  int
	test8  int
	test9  int
	test10 int
	test11 int
	test12 int
	test13 int
	test14 int
	test15 int
}

func newBigStruct(i int) bigStruct {
	return bigStruct{i, i + 1, i + 2, i + 4, i + 5, i + 6, i + 7, i + 8, i + 9, i + 10, i + 11, i + 12, i + 13, i + 14, i + 15}
}

func BenchmarkSyncShardIter(b *testing.B) {
	concurrentMap := newShardedConcurrentMap[string, bigStruct](DefaultStringShardFunc)
	for i := range 1000000 {
		concurrentMap.Set(fmt.Sprintf("%d", i), newBigStruct(i))
	}

	b.ResetTimer()
	for b.Loop() {
		for key, value := range concurrentMap.Next {
			_ = key
			_ = value
		}
	}
}

func BenchmarkParallelShardIter(b *testing.B) {
	concurrentMap := newShardedConcurrentMap[string, bigStruct](DefaultStringShardFunc)
	for i := range 1000000 {
		concurrentMap.Set(fmt.Sprintf("%d", i), newBigStruct(i))
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for key, value := range concurrentMap.Next {
				_ = key
				_ = value
			}
		}
	})
}
