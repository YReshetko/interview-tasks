package main

import "fmt"

func main() {
	ng := newGraph([][]int{
		{0},
		{7, 0},
		{9, 10, 0},
		{0, 15, 11, 0},
		{0, 0, 0, 6, 0},
		{14, 0, 2, 0, 9, 0},
	})
	/*ng.print()
	ng.printEdges()
	*/
	fmt.Println(ng.path(1, 2))

}
