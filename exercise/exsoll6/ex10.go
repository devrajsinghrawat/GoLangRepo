package exsoll6

import "fmt"

var g func()

// Ex10 -Exported function to main -
func Ex10() {

	c := func() {
		fmt.Printf("I'm function variable \n")
	}

	g = c

	func() {
		fmt.Printf("I'm anonymous function \n")
	}()
}
