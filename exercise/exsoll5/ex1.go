package exsoll5

import "fmt"

type person struct {
	fname   string
	lname   string
	flavors []string
}

// Ex1 -Exported function to main - Struct, map and slice
func Ex1() {

	p1 := person{fname: "jhon", lname: "snow", flavors: []string{"vanila", "choco"}}
	var p2 = person{fname: "bran", lname: "stark", flavors: []string{"strawberry", "tuti-fruti"}}

	s := make([]person, 3, 5)

	s = []person{p1, p2}

	fmt.Printf("%v \n", s)

	m := map[string]person{p1.lname: p1, p2.lname: p2}
	fmt.Printf("KEY \t \t VALUE \n")
	for key, val := range m {
		fmt.Printf("%s \t \t %s\n", key, val.fname)
		fmt.Printf(" \t \t \t --------- \n")
		fmt.Printf(" \t \t \t Flavors \n")
		fmt.Printf(" \t \t \t --------- \n")
		for _, flav := range val.flavors {
			fmt.Printf("\t \t \t %s\n", flav)
		}

	}

}
