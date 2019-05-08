package exsoll1

import "fmt"

type mytype int

var x1 mytype
var y1 int

// var y1 int

// Ex3 -Exported function to main
func Ex3() {
	fmt.Printf("variable z from ex2 file package level scope but small case - %s \n", y)
	fmt.Printf("Type of custom defined variable - %T \n", x1)
	fmt.Println(x1)
	x1 = 20

	// y1 := x1            // Local Scope like a new variable
	y1 = int(x1)
	fmt.Println(y1)
	fmt.Printf("Type of y1 - %T \n", y1)

}
