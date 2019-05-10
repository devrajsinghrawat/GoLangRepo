package exsoll6

import (
	"fmt"
)

type person struct {
	fname string
	lname string
}

type student struct {
	person
	rollnumber int
	grade      string
}

type living interface {
	speak()
}

func (p person) speak() {
	fmt.Printf("i'm student's method due to receiver type %s %s - Person Speak \n", p.fname, p.lname)
}

func (s student) speak() {
	fmt.Printf("i'm student's method due to receiver type %s %s - Student Speak \n", s.fname, s.lname)
}

func livfun(l living) {
	fmt.Print("This is a function \n", l)
}

// Ex3 -Exported function to main - Struct, map and slice
func Ex3() {

	p1 := person{fname: "Dev",
		lname: "D"}

	s1 := student{person: person{
		fname: "Jhon",
		lname: "Snow"},
		rollnumber: 1,
		grade:      "A++"}

	s2 := student{person: person{
		fname: "Bran",
		lname: "Stark"},
		rollnumber: 1,
		grade:      "B++"}

	s1.speak()
	s2.speak()
	p1.speak()

	livfun(p1)
	livfun(s1)
	livfun(s2)
}
