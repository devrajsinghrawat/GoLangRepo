package pointers

import "fmt"

type person struct {
	fname   string
	lname   string
	flavors []string
}

// Ex3 -Exported function to main - Struct, map and slice
func Ex3() {

	p1 := person{fname: "jhon", lname: "snow", flavors: []string{"vanila", "choco"}}
	// var p2 = person{fname: "bran", lname: "stark", flavors: []string{"strawberry", "tuti-fruti"}}

	fmt.Printf("Person Struct before - %s \n", p1)
	changeMe(&p1)
	fmt.Printf("Person Struct after - %s \n", p1)

}

func changeMe(p *person) {
	(*p).flavors[0] = "tuti-fruti"
	// p.flavors[0] = "tuti-fruti"  // or
}
