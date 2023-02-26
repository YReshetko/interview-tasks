package binary_search

type Comparable interface {
	int | int32 | int64 | float32 | float64 | string
}

func Search[T Comparable](a []T, value T) (int, bool) {
	if len(a) == 0 {
		return -1, false
	}
	lo := 0
	hi := len(a)

	for lo < hi {
		mid := (hi + lo) / 2
		if value < a[mid] {
			hi = mid - 1
		} else if value > a[mid] {
			lo = mid + 1
		} else {
			return mid, true
		}
	}
	return -1, false
}
