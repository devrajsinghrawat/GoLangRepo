package exsoll1

import "fmt"

// Ex8 -Exported function to main
func Ex8() {
	a := []int{1, 2, 3}
	b := []int{4, 5, 6}
	c := [][]int{a, b}

	fmt.Println(c)

}
