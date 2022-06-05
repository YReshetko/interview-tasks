package task_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	task "github.com/YReshetko/interview-tasks/pkg/area"
)

func TestRun_InvalidArea(t *testing.T) {
	tests := []struct {
		area [][]int
	}{
		{area: nil},
		{area: [][]int{}},
		{area: [][]int{{}, {}, {}}},
		{area: [][]int{{1}, {1, 1, 1}, {1}}},
	}
	for i, test := range tests {
		t.Run(fmt.Sprintf("Invalid area %d", i), func(t *testing.T) {
			assert.Nil(t, task.Run(test.area))
		})
	}
}

func TestRun(t *testing.T) {
	tests := []struct {
		area     [][]int
		expected []int
	}{
		{
			area:     [][]int{{1, 1, 0, 0, 0, 1, 0, 1, 1, 0, 0}},
			expected: []int{1, 2, 3},
		},
		{
			area:     [][]int{{1}, {1}, {0}, {0}, {0}, {1}, {0}, {1}, {1}, {0}, {0}},
			expected: []int{1, 2, 3},
		},
		{
			area: [][]int{
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0}},
			expected: []int{24},
		},
		{
			area: [][]int{
				{1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1},
				{1, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1},
				{1, 0, 0, 1, 1, 1, 0, 0, 0, 1, 1},
				{0, 1, 1, 0, 1, 1, 0, 0, 1, 1, 1},
				{0, 0, 0, 0, 1, 1, 1, 0, 1, 1, 1},
				{0, 0, 0, 0, 1, 1, 0, 0, 1, 1, 1},
				{0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1},
				{0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0},
				{0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 1},
				{0, 0, 0, 0, 1, 1, 1, 0, 0, 1, 1}},
			expected: []int{5, 8, 9, 26},
		},
	}
	for i, test := range tests {
		t.Run(fmt.Sprintf("Valid area %d", i), func(t *testing.T) {
			actual := task.Run(test.area)
			assert.NotNil(t, actual)
			assert.Equal(t, test.expected, actual)

		})
	}
}
