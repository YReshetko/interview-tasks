package heap_pattern

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHeap_Insert(t *testing.T) {
	values := []int{35, 12, 78, 24, 93, 12, 4, 97, 89, 53, 24, 7, 67}
	h := NewHeap()
	for _, value := range values {
		h.Insert(value)
	}

	expected := []int{97, 93, 89, 78, 67, 53, 35, 24, 24, 12, 12, 7, 4}
	actual := []int{}
	for !h.IsEmpty() {
		actual = append(actual, h.Extract())
	}
	assert.Equal(t, expected, actual)
}
