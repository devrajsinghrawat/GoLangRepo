package exsoll1

import "fmt"

var x = 42
var y = "Dev"
var z = true

// Ex2 -Exported function to main
func Ex2() {
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	s := fmt.Sprintf("%v\t %v\t %v\t ", x, y, z)

	fmt.Println(s)
}
