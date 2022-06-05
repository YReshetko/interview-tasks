package find_sum

import "sort"

func TwoIndex(arr []int, sum int) (int, int) {
	sort.Ints(arr)
	l, r := 0, len(arr)-1
	var a1, a2 int
	for {
		a1, a2 = arr[l], arr[r]
		res := a1 + a2
		if res == sum {
			break
		}
		if res > sum {
			r--
		}
		if res < sum {
			l++
		}
	}
	return a1, a2
}
