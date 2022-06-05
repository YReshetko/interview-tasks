package binary_tree_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/YReshetko/interview-tasks/pkg/binary-tree"
)

func TestPrint(t *testing.T) {
	root := tree()
	//binary_tree.Print(root)

	tests := []struct {
		v1, v2   int
		expected int
	}{
		{v1: 35, v2: 34, expected: 24},
		{v1: 40, v2: 34, expected: 19},
		{v1: 44, v2: 34, expected: 10},
		{v1: 40, v2: 44, expected: 10},
		{v1: 20, v2: 35, expected: 10},
		{v1: 25, v2: 35, expected: 19},
		{v1: 29, v2: 50, expected: 20},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%d - %d: %d", test.v1, test.v2, test.expected), func(t *testing.T) {
			actual := binary_tree.FindCommon(root, test.v1, test.v2)
			assert.Equal(t, test.expected, actual)
			actual = binary_tree.FindCommon(root, test.v2, test.v1)
			assert.Equal(t, test.expected, actual)
		})
	}
}

func tree() binary_tree.Leaf {
	l1 := binary_tree.NewLeaf(50, nil, nil)
	l2 := binary_tree.NewLeaf(49, nil, nil)

	l3 := binary_tree.NewLeaf(45, nil, nil)
	l4 := binary_tree.NewLeaf(44, nil, nil)

	l5 := binary_tree.NewLeaf(40, nil, nil)
	l6 := binary_tree.NewLeaf(39, nil, nil)

	l7 := binary_tree.NewLeaf(35, nil, nil)
	l8 := binary_tree.NewLeaf(34, nil, nil)

	l9 := binary_tree.NewLeaf(30, &l1, &l2)
	l10 := binary_tree.NewLeaf(29, &l3, &l4)

	l11 := binary_tree.NewLeaf(25, &l5, &l6)
	l12 := binary_tree.NewLeaf(24, &l7, &l8)

	l13 := binary_tree.NewLeaf(20, &l9, &l10)
	l14 := binary_tree.NewLeaf(19, &l11, &l12)

	return binary_tree.NewLeaf(10, &l13, &l14)
}
