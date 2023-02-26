package sorting

type InsertionSorter struct{}

func (i InsertionSorter) Sort(arr []int) {
	for i := 1; i < len(arr); i++ {
		v := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > v {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = v
	}
}
