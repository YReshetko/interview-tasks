package sorting

type ShakerSorter struct{}

func (i ShakerSorter) Sort(arr []int) {
	start, end := 0, len(arr)-1

	for start < end {
		for i := start; i < end; i++ {
			if arr[i+1] < arr[i] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
			}
		}
		end--
		for i := end; i > start; i-- {
			if arr[i] < arr[i-1] {
				arr[i], arr[i-1] = arr[i-1], arr[i]
			}
		}
		start++
	}
}
