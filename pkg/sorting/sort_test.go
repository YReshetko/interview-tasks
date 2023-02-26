package sorting_test

import (
	"github.com/YReshetko/interview-tasks/pkg/sorting"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInsertionSorter(t *testing.T) {
	arr := []int{10, 4, 8, 2, 7, 5, 1, 3, 6, 9}
	sorting.InsertionSorter{}.Sort(arr)

	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, arr)
}

func TestSelectionSorter(t *testing.T) {
	arr := []int{10, 4, 8, 2, 7, 5, 1, 3, 6, 9}
	sorting.SelectionSorter{}.Sort(arr)

	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, arr)
}

func TestQuickSorter(t *testing.T) {
	arr := []int{10, 4, 8, 2, 7, 5, 1, 3, 6, 9}
	sorting.QuickSorter{}.Sort(arr)

	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, arr)
}

func TestMergeSorter(t *testing.T) {
	arr := []int{10, 4, 8, 2, 7, 5, 1, 3, 6, 9}
	sorting.MergeSorter{}.Sort(arr)

	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, arr)
}

func TestShakerSorter(t *testing.T) {
	arr := []int{10, 4, 8, 2, 7, 5, 1, 3, 6, 9}
	sorting.ShakerSorter{}.Sort(arr)

	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, arr)
}
