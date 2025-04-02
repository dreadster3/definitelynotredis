package concurrentmap

import (
	"fmt"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConcurrentMapGet(t *testing.T) {
	data := NewConcurrentMap[string, string]()

	data.data["key"] = "value"

	assert.Equal(t, "value", data.Get("key"))
}

func TestConcurrentMapSet(t *testing.T) {
	data := NewConcurrentMap[string, string]()

	data.Set("key", "value")

	assert.Equal(t, "value", data.data["key"])
}

func TestConcurrentMapDelete(t *testing.T) {
	data := NewConcurrentMap[string, string]()

	data.Set("key", "value")
	data.Delete("key")

	assert.NotContains(t, data.data, "key")
}

func TestConcurrentMapConcurrentSetGetDelete(t *testing.T) {
	data := NewConcurrentMap[string, int]()

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
				value := data.Get(key)
				assert.Equalf(t, j, value, "Go routine %d: expected %d, got %d", id, j, value)

				data.Delete(key)
			}
		}(i)
	}

	wg.Wait()
}
