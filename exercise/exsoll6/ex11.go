package exsoll6

import "fmt"

var y int

// Ex11 -Exported function to main -
func Ex11() {
	c := increment()
	c()
	c()
	c()
	c()
	d := increment()
	d()
	d()
	c()
	d()
}

func increment() func() {
	x := 0
	fmt.Printf("I'm inside function %v %v \n", x, y)

	return func() {
		x++
		y++
		fmt.Printf("increment x and printing %v %v \n", x, y)
	}
}
