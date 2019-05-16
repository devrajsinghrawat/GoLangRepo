package pointers

import (
	"fmt"
	"sort"
)

type person5 struct {
	fname string
	lname string
	age   int
}

// ByAge implements sort.Interface for []Person based on
// the Age field.
type ByAge []person5

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].age < a[j].age }

// Ex5 -Exported function to main - json Marshal
func Ex5() {

	p1 := person5{fname: "jhon1", lname: "snow", age: 30}
	p2 := person5{fname: "bran2", lname: "stark", age: 23}
	p3 := person5{fname: "jhon3", lname: "snow", age: 10}
	p4 := person5{fname: "bran4", lname: "stark", age: 40}

	// ByAge = append(ByAge, {p1, p2, p3, p4})

	person := []person5{p1, p2, p3, p4}

	sort.Sort(ByAge(person))
	fmt.Print(person)
}
