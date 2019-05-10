package exsoll6

import "fmt"

type person8 struct {
	fname string
	lname string
	age   int
}

func (p1 person8) speaker() {
	fmt.Print("I'm ", p1.fname, " ", p1.lname, " and i'm ", p1.age, " year old \n")
}

// Ex8 -Exported function to main - Defer exercise
func Ex8() {

	p := person8{fname: "devd", lname: "singh", age: 30}
	p.speaker()

}
