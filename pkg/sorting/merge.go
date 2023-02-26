package sorting

type MergeSorter struct{}

func (i MergeSorter) Sort(arr []int) {
	i.mergeSort(arr, 0, len(arr)-1)
}

func (i MergeSorter) mergeSort(arr []int, start, end int) {
	if start < end {
		middle := (start + end) / 2
		i.mergeSort(arr, start, middle)
		i.mergeSort(arr, middle+1, end)
		i.merge(arr, start, middle, end)
	}
}

func (i MergeSorter) merge(arr []int, start, middle, end int) {
	buf := make([]int, end-start+1)
	lIndex, rIndex := start, middle+1
	index := 0
	for ; lIndex <= middle && rIndex <= end; index++ {
		if arr[lIndex] < arr[rIndex] {
			buf[index] = arr[lIndex]
			lIndex++
		} else {
			buf[index] = arr[rIndex]
			rIndex++
		}
	}

	for i := lIndex; i <= middle; i++ {
		buf[index] = arr[i]
		index++
	}
	for i := rIndex; i <= end; i++ {
		buf[index] = arr[i]
		index++
	}

	for i := 0; i < end-start+1; i++ {
		arr[i+start] = buf[i]
	}
}
