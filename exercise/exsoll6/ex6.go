package exsoll6

import "fmt"

// Ex6 -Exported function to main - Struct, map and slice
func Ex6() {

	i := []int{1, 2, 3, 4, 5}
	// s2 := foo(i)             #1  passing a slice
	s2 := foo(i...) // #2

	// j, s := bar()

	fmt.Print(s2)
}

// func foo() int {
// 	return 4
// }

// func bar() (int, string) {
// 	return 5, "hello bar"
// }

// func foo(n []int) int {        # 1  receiving a slice
func foo(n ...int) int { // # 2
	var s1 int
	for _, v := range n {
		s1 = s1 + v
	}

	return s1
}
