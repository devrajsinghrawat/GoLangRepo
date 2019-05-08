package exsoll1

import "fmt"

// Ex1 -Exported function to main
func Ex1() {
	x := 42
	y := "james"
	z := true

	fmt.Println(x, y, z)
	fmt.Print(x, y, z)
	fmt.Printf("%b %s %t \n", x, y, z)
}
