package exsoll1

import "fmt"

const typed int = 42
const untyped = 42.1

// Ex4 -Exported function to main
func Ex4() {
	x := 16
	fmt.Printf("Decimal %d \n", x)
	fmt.Printf("Binary %b \n", x)
	fmt.Printf("Hex %#x \n", x)

	fmt.Printf("typed - %T %v \n", typed, typed)
	fmt.Printf("untyped -%T %v\n", untyped, untyped)

	x = x << 1
	fmt.Printf("Decimal %d \n", x)
	fmt.Printf("Binary %b \n", x)
	fmt.Printf("Hex %#x \n", x)
}
