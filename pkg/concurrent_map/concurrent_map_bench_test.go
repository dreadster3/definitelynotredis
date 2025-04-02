package concurrentmap

import "testing"

func BenchmarkConcurrentMapSetGetDelete(b *testing.B) {
	concurrentMap := NewConcurrentMap[string, int]()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			key := "key"
			value := 1
			concurrentMap.Set(key, value)
			_ = concurrentMap.Get(key)
			concurrentMap.Delete(key)
		}
	})
}
