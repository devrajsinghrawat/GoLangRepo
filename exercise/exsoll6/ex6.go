package exsoll6

import "fmt"

// Ex6 -Exported function to main - Struct, map and slice
func Ex6() {

	i := []int{1, 2, 3, 4, 5}
	s2 := foo(i...)

	s3 := bar(i)

	fmt.Print(s2)
	fmt.Print(s3)
}

func bar(n []int) int {
	var s1 int
	for _, v := range n {
		s1 = s1 + v
	}

	return s1
}

func foo(n ...int) int { // # 2
	var s1 int
	for _, v := range n {
		s1 = s1 + v
	}

	return s1
}
