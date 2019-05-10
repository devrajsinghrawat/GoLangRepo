package exsoll6

import "fmt"

// Ex5 -Exported function to main - Struct, map and slice
func Ex5() {
	var count, result int
	count = 4
	result = 1
	for index := count; index > 1; index-- {
		result *= index
		// fmt.Print(result, count)
	}

	fmt.Print(result)
}
