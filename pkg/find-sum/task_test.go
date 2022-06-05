package find_sum_test

import (
	"fmt"
	find_sum "github.com/YReshetko/interview-tasks/pkg/find-sum"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTwoIndex(t *testing.T) {
	tests := []struct {
		arr    []int
		sum    int
		a1, a2 int
	}{
		{
			arr: []int{1, 5, 8, 10, 13, 18, 24},
			sum: 28,
			a1:  10,
			a2:  18,
		}, {
			arr: []int{13, 24},
			sum: 37,
			a1:  13,
			a2:  24,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("Sum %d", i), func(t *testing.T) {
			r1, r2 := find_sum.TwoIndex(test.arr, test.sum)
			assert.Equal(t, r1, test.a1)
			assert.Equal(t, r2, test.a2)
		})
	}
}
