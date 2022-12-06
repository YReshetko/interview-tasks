package bridge_family

import (
	"fmt"
	"sort"
)

func Run() string {
	n := node{
		state: state{
			leftSide:  []person{father, mother, grandmother, kid},
			rightSide: []person{},
		},
		transmissions: map[move]node{},
	}

	n.moveRight()

	res := n.results()
	sort.Slice(res, func(i, j int) bool {
		return res[i].time < res[j].time
	})

	for _, r := range res {
		fmt.Println(r.string())
	}
	return res[0].string()
}
