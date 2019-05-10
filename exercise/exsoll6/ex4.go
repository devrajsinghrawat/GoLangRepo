package exsoll6

import "fmt"

// Ex4 -Exported function to main - Struct, map and slice
func Ex4() {

	n := []int{1, 2, 3, 4, 5}

	s := even(sum, n...)

	fmt.Printf("result - %v \n", s)
}

func sum(n1 ...int) int {
	var r int
	for _, s := range n1 {
		r += s
	}
	return r
}

func even(f func(x ...int) int, e ...int) int {
	var r []int
	for _, s := range e {
		if s%2 == 0 {
			r = append(r, s)
		}
	}

	return f(r...)

}
