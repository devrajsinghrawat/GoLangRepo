package exsoll6

import "fmt"

// Ex2 -Exported function to main - functions defer
func Ex2() {
	for i := 0; i < 4; i++ {
		defer fmt.Printf("%v \n", i)
	}
}
