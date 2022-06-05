package linked_list_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/YReshetko/interview-tasks/pkg/linked-list"
)

func TestMiddle(t *testing.T) {
	test := []struct {
		arr   []int
		value int
	}{
		{arr: []int{1, 3, 4, 5, 7, 8, 10, 12, 13, 15}, value: 8},
		{arr: []int{1, 3, 4, 5, 7, 8, 10, 12, 13, 15, 16, 17, 19}, value: 10},
	}
	for _, s := range test {
		node := linked_list.NewList(s.arr)
		v := linked_list.Middle(node)
		assert.Equal(t, s.value, v)
	}
}

func TestFindFromEnd(t *testing.T) {
	test := []struct {
		arr   []int
		index int
		value int
	}{
		{arr: []int{1, 3, 4, 5, 7, 8, 10, 12, 13, 15}, value: 15, index: 1},
		{arr: []int{1, 3, 4, 5, 7, 8, 10, 12, 13, 15, 16, 17, 19}, value: 15, index: 4},
	}
	for _, s := range test {
		node := linked_list.NewList(s.arr)
		v := linked_list.FindFromEnd(node, s.index)
		assert.Equal(t, s.value, v)
	}
}

func TestReverse(t *testing.T) {
	test := []struct {
		arr      []int
		expected []int
	}{
		{
			arr:      []int{1, 3, 4, 5, 7, 8, 10, 12, 13, 15},
			expected: []int{15, 13, 12, 10, 8, 7, 5, 4, 3, 1},
		},
		{
			arr:      []int{1, 3, 4, 5, 7, 8, 10, 12, 13, 15, 16, 17, 19},
			expected: []int{19, 17, 16, 15, 13, 12, 10, 8, 7, 5, 4, 3, 1},
		},
	}
	for _, s := range test {
		node := linked_list.NewList(s.arr)
		v := linked_list.Reverse(node)
		assert.Equal(t, s.expected, v.ToArray())
	}
}
