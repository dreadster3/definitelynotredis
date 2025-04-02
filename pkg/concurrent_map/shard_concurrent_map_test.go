package concurrentmap

import (
	"fmt"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShardedConcurrentMapGet(t *testing.T) {
	data := newShardedConcurrentMap[string, string](DefaultStringShardFunc)

	key := "key"
	expected := "value"

	data.shards[data.shardFunc(key)%data.size].data[key] = expected
	actual, exists := data.Get(key)

	assert.True(t, exists)
	assert.Equal(t, expected, actual)
}

func TestShardedConcurrentMapSet(t *testing.T) {
	data := newShardedConcurrentMap[string, string](DefaultStringShardFunc)

	key := "key"
	expected := "value"

	data.Set(key, expected)

	assert.Equal(t, expected, data.shards[data.shardFunc(key)%data.size].data[key])
}

func TestShardedConcurrentMapDelete(t *testing.T) {
	data := newShardedConcurrentMap[string, string](DefaultStringShardFunc)

	key := "key"
	expected := "value"

	data.Set(key, expected)
	data.Delete(key)

	assert.NotContains(t, data.shards[data.shardFunc(key)%data.size].data, key)
}

func TestShardedConcurrentMapConcurrentSetGetDelete(t *testing.T) {
	data := newShardedConcurrentMap[string, int](DefaultStringShardFunc)

	var wg sync.WaitGroup
	goRoutines := 10
	numReps := 1000

	for i := range goRoutines {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			for j := range numReps {
				key := fmt.Sprintf("key-%d-%d", id, j)
				data.Set(key, j)
				value, exists := data.Get(key)

				assert.True(t, exists)
				assert.Equalf(t, j, value, "Go routine %d: expected %d, got %d", id, j, value)

				data.Delete(key)

				_, exists = data.Get(key)
				assert.False(t, exists)
			}
		}(i)
	}

	wg.Wait()
}
