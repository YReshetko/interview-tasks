package coordinates_test

import (
	"github.com/YReshetko/interview-tasks/pkg/coordinates"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindClosest(t *testing.T) {
	tests := []struct {
		points   [][2]int
		number   int
		expected [][2]int
	}{
		{
			points:   [][2]int{{0, 1}, {2, 2}, {5, 5}},
			number:   2,
			expected: [][2]int{{0, 1}, {2, 2}},
		},
		{
			points:   [][2]int{{5, 6}, {0, 1}, {1, 3}, {1, 0}, {2, 0}, {0, 2}, {2, 2}, {5, 5}},
			number:   4,
			expected: [][2]int{{0, 1}, {1, 0}, {2, 0}, {0, 2}},
		},
	}

	for _, test := range tests {
		res := coordinates.FindClosest(test.points, test.number)
		assert.Equal(t, test.expected, res)
	}
}
