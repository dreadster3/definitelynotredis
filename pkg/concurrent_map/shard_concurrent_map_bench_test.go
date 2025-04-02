package concurrentmap

import (
	"math/rand/v2"
	"strings"
	"testing"
)

const charset = "abcdefghijklmnopqrstuvwxyz"

func randomString(n int) string {
	sb := strings.Builder{}
	sb.Grow(n)
	for range n {
		sb.WriteByte(charset[rand.IntN(len(charset))])
	}
	return sb.String()
}

func RandomShardFunc(key string) uint32 {
	return rand.Uint32()
}

func BenchmarkShardConcurrentMapSetGetDelete(b *testing.B) {
	concurrentMap := newShardedConcurrentMap[string, int](DefaultStringShardFunc)

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

func BenchmarkShardConcurrentMapDifferentKeySet(b *testing.B) {
	concurrentMap := newShardedConcurrentMap[string, int](RandomShardFunc)

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			concurrentMap.Set("key", 1)
		}
	})
}

func BenchmarkShardConcurrentMapDifferentKeyGet(b *testing.B) {
	concurrentMap := newShardedConcurrentMap[string, int](RandomShardFunc)

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			concurrentMap.Get("key")
		}
	})
}
