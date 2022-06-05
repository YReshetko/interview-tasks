package task

import (
	"sort"
)

type point [2]int

func (p point) add(p2 point) point {
	return [2]int{p[0] + p2[0], p[1] + p2[1]}
}

var steps = [4]point{{0, 1}, {1, 0}, {-1, 0}, {0, -1}}

type stack []point

func (s *stack) push(p point) {
	st := *s
	st = append([]point{p}, st...)
	*s = st
}

func (s *stack) pop() point {
	st := *s
	p, n := st[0], st[1:]
	*s = n
	return p
}

func (s stack) size() int {
	return len(s)
}

type marks [][]bool

func newMarks(n, m int) marks {
	res := make([][]bool, n, n)
	for i := range res {
		res[i] = make([]bool, m, m)
	}
	return res
}

func (m marks) mark(p point) {
	m[p[0]][p[1]] = true
}

func (m marks) isChecked(p point) bool {
	return m[p[0]][p[1]]
}

type area [][]int

func (a area) outOfBound(p point) bool {
	return p[0] < 0 || p[1] < 0 || p[0] >= len(a) || p[1] >= len(a[0])
}

func (a area) isPlain(p point) bool {
	return a[p[0]][p[1]] == 1
}

func (a area) size() (int, int) {
	return len(a), len(a[0])
}

func (a area) valid() bool {
	if len(a) == 0 || len(a[0]) == 0 {
		return false
	}
	s := len(a[0])
	for _, line := range a {
		if len(line) != s {
			return false
		}
	}
	return true
}

func (a area) forEach(fn func(point)) {
	for i, v := range a {
		for j, _ := range v {
			fn([2]int{i, j})
		}
	}
}

func Run(a area) []int {
	if !a.valid() {
		return nil
	}
	m := newMarks(a.size())
	sizes := make([]int, 0)

	a.forEach(func(p point) {
		if a.isPlain(p) || m.isChecked(p) {
			m.mark(p)
			return
		}
		sizes = append(sizes, areaSize(a, m, p))
	})
	sort.Ints(sizes)

	return sizes
}

func areaSize(a area, m marks, p point) int {
	size := 1
	s := new(stack)
	s.push(p)
	m.mark(p)
	for s.size() > 0 {
		p = s.pop()
		for _, step := range steps {
			p2 := p.add(step)
			if a.outOfBound(p2) {
				continue
			}
			if a.isPlain(p2) || m.isChecked(p2) {
				m.mark(p2)
				continue
			}
			size++
			s.push(p2)
			m.mark(p2)
		}
	}
	return size
}
