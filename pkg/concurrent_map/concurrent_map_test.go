package concurrentmap

import (
	"fmt"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConcurrentMapGet(t *testing.T) {
	data := newConcurrentMap[string, string]()

	key := "key"
	expected := "value"

	data.data[key] = expected
	actual, exists := data.Get("key")

	assert.True(t, exists)
	assert.Equal(t, expected, actual)
}

func TestConcurrentMapSet(t *testing.T) {
	data := newConcurrentMap[string, string]()

	key := "key"
	expected := "value"

	data.Set("key", "value")

	assert.Equal(t, expected, data.data[key])
}

func TestConcurrentMapDelete(t *testing.T) {
	data := newConcurrentMap[string, string]()

	data.Set("key", "value")
	data.Delete("key")

	assert.NotContains(t, data.data, "key")
}

func TestConcurrentMapConcurrentSetGetDelete(t *testing.T) {
	data := newConcurrentMap[string, int]()

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
