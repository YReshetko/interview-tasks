package coordinates

import (
	"math"
	"sort"
)

type coordinates [][2]int

func (a coordinates) Len() int           { return len(a) }
func (a coordinates) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a coordinates) Less(i, j int) bool { return distance(a[i]) < distance(a[j]) }

func distance(p [2]int) float64 {
	x := float64(p[0])
	y := float64(p[1])
	return math.Sqrt(x*x + y*y)
}

func FindClosest(points [][2]int, size int) [][2]int {
	sort.Sort(coordinates(points))
	return points[:size]
}
