package exsoll6

import "fmt"

// Ex1 -Exported function to main - functions Variadic and unfurling parameters
func Ex1(t ...int) { // expecting unlimited number of int's and packing it into a slice
	fmt.Print(t)
	fmt.Printf("Type %T \n", t)
	for _, val := range t {
		fmt.Printf("Value %v \n", val)
	}
}
