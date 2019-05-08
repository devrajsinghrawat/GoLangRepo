package exsoll4

import (
	"fmt"
)

var names = make([]string, 10, 15)
var bals = make([]int, 10, 15)
var bals2 = make([]int, 10, 15)

// Ex2 -Exported function to main
func Ex2() {

	fmt.Printf("A - %v\n", names)
	fmt.Printf("capacity of slice %v \n", cap(names))
	names = []string{"0x1", "0x2", "0x3", "0x4", "0x5", "0x6", "0x7", "0x8", "0x9"}

	bals = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	bals2 = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Printf("B - %v\n", bals)
	fmt.Printf("capacity of slice %v \n", cap(bals))

	combine := [][]int{bals, bals2}
	fmt.Printf("C - %v\n", combine)
	fmt.Printf("capacity of slice %v \n", cap(combine))
}
