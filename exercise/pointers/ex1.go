package pointers

import "fmt"

// Ex1 -Exported function to main - Pointers
func Ex1() {

	a := 32
	fmt.Printf("Value of A %v \n", a)
	fmt.Printf("Address of A %v \n", &a)
	fmt.Printf("Type of A - %T\n", a)
	fmt.Printf("Type of &A - %T\n", &a)

	b := &a
	fmt.Printf("Value of B %v \n", b)
	fmt.Printf("Address of B %v \n", &b)
	fmt.Printf("Type of B - %T\n", b)
	fmt.Printf("Value stored at Pointer of B - %v\n", *b)

	// Update the value of A
	*b = 42
	fmt.Printf("Updated Value of A %v \n", a)

}
