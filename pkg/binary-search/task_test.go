package binary_search_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/YReshetko/interview-tasks/pkg/binary-search"
)

func TestFindInt(t *testing.T) {
	tests := []struct {
		arr   []int
		value int
		ind   int
		find  bool
	}{
		{arr: []int{}, value: 10, ind: -1, find: false},
		{arr: []int{1}, value: 10, ind: -1, find: false},
		{arr: []int{1, 3, 7, 10, 15, 17, 23}, value: 10, ind: 3, find: true},
		{arr: []int{1, 3, 7, 8, 9, 10}, value: 10, ind: 5, find: true},
		{arr: []int{10}, value: 10, ind: 0, find: true},
		{arr: []int{10, 11, 13, 17, 19}, value: 10, ind: 0, find: true},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("Search %d", i), func(t *testing.T) {
			ind, ok := binary_search.Search(test.arr, test.value)
			assert.Equal(t, test.ind, ind)
			assert.Equal(t, test.find, ok)
		})
	}
}
