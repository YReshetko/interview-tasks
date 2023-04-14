package main

import (
	"fmt"
	"sort"
)

// Edge value from knot i to j is placed to the matrix
type graph [][]int

func newGraph(value [][]int) graph {
	ng := zeroMatrix(len(value))
	for i, arr := range value {
		for j := 0; j < i; j++ {
			ng[i][j] = arr[j]
			ng[j][i] = arr[j]
		}
	}

	return ng
}

func (g graph) print() {
	for _, arr := range g {
		fmt.Println(arr)
	}
}

func (g graph) printEdges() {
	for i, arr := range g {
		for j, v := range arr {
			if i == j {
				break
			}
			if v == 0 {
				continue
			}
			fmt.Printf("%d -> %d = %d\n", i+1, j+1, v)
		}
	}
}

type stack [][2]int

func (s *stack) push(index, sum int) {
	if s == nil {
		return
	}
	sp := *s
	sp = append(sp, [2]int{index, sum})
	sort.Slice(sp, func(i, j int) bool {
		return sp[i][1] > sp[j][1]
	})
	*s = sp
}

func (s *stack) pop() [2]int {
	if s == nil || len(*s) == 0 {
		return [2]int{}
	}
	sp := *s
	if len(sp) == 1 {
		v := sp[0]
		*s = [][2]int{}
		return v
	}
	v, sp := sp[len(sp)-1], sp[:len(sp)-1]
	*s = sp
	return v
}

func (s *stack) isEmpty() bool {
	if s == nil {
		return true
	}
	return len(*s) == 0
}

func (g graph) path(from, to int) int {
	s := &stack{}
	s.push(from-1, 0)
	c := make(map[int]struct{}, len(g))
	paths := make([]int, len(g))
	for i := 0; i < len(paths); i++ {
		paths[i] = 1 << 31
	}

	maxOf := func(jj int) (int, int, bool) {
		max, j := -1, -1
		for i := 0; i < len(g); i++ {
			_, ok := c[i]
			if ok || g[jj][i] == 0 || g[jj][i] == 1<<31 {
				continue
			}
			if g[jj][i] > max {
				j = i
				max = g[jj][i]
			}
		}
		if j != -1 {
			g[jj][j] = -1
		}
		return j, max, j != -1
	}

	for !s.isEmpty() {
		p := s.pop()
		c[p[0]] = struct{}{}
		if p[0] == to-1 {
			return p[1]
		}
		fmt.Println(p)
		for {
			j, max, ok := maxOf(p[0])
			if !ok {
				break
			}
			if _, ok := c[j]; !ok && paths[j] > max+p[1] {
				paths[j] = max + p[1]
				s.push(j, max+p[1])
			}
		}
	}
	return 1
}

func (g graph) size() int {
	return len(g)
}

func zeroMatrix(l int) [][]int {
	out := make([][]int, l)
	for i := 0; i < l; i++ {
		out[i] = make([]int, l)
	}
	return out
}
