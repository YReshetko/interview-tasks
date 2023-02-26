package sorting

type QuickSorter struct{}

func (i QuickSorter) Sort(arr []int) {
	i.quicksort(arr, 0, len(arr)-1)
}

func (i QuickSorter) quicksort(arr []int, start, end int) {
	if start >= end {
		return
	}
	pivot := i.partition(arr, start, end)
	i.quicksort(arr, start, pivot-1)
	i.quicksort(arr, pivot+1, end)
}

func (i QuickSorter) partition(arr []int, start, end int) int {
	pivot := arr[end]
	pIndex := start
	for i := start; i < end; i++ {
		if arr[i] < pivot {
			arr[pIndex], arr[i] = arr[i], arr[pIndex]
			pIndex++
		}
	}

	arr[pIndex], arr[end] = arr[end], arr[pIndex]
	return pIndex
}
