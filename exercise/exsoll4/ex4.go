package exsoll4

import "fmt"

type person struct {
	string  // Anonymous field
	address string
}

type student struct {
	person  // Anonymous field
	id, age int
}

// Ex4 -Exported function to main - Struct
func Ex4() {

	s1 := student{
		person: person{string: "DevD", address: "Sre"},
		id:     1,
		age:    26,
	}

	fmt.Printf("%v \n", s1)

}
