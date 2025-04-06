package concurrentmap

import (
	"fmt"
	"maps"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	data := newConcurrentMap[string, string]()

	key := "key"
	expected := "value"

	data.data[key] = expected
	actual, exists := data.Get("key")

	assert.True(t, exists)
	assert.Equal(t, expected, actual)
}

func TestSet(t *testing.T) {
	data := newConcurrentMap[string, string]()

	key := "key"
	expected := "value"

	data.Set("key", "value")

	assert.Equal(t, expected, data.data[key])
}

func TestDelete(t *testing.T) {
	data := newConcurrentMap[string, string]()

	data.Set("key", "value")
	data.Delete("key")

	assert.NotContains(t, data.data, "key")
}

func TestConcurrentSetGetDelete(t *testing.T) {
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

func TestIter(t *testing.T) {
	data := newConcurrentMap[string, int]()
	data.Set("1", 1)
	data.Set("2", 2)
	data.Set("3", 3)
	data.Set("4", 4)
	data.Set("5", 5)
	data.Set("6", 6)
	data.Set("7", 7)

	expected := map[string]int{
		"1": 1,
		"2": 2,
		"3": 3,
		"4": 4,
		"5": 5,
		"6": 6,
		"7": 7,
	}
	actual := maps.Collect(data.Next)

	assert.Equal(t, expected, actual)
}
