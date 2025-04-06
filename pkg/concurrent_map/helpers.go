package concurrentmap

import "iter"

func Filter[K comparable, V any](source iter.Seq2[K, V], filter func(K, V) bool) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for key, value := range source {
			if filter(key, value) {
				if !yield(key, value) {
					return
				}
			}
		}
	}
}
