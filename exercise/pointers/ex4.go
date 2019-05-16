package pointers

import (
	"encoding/json"
	"fmt"
	"sort"
)

type person4 struct {
	Fname string
	Lname string
	Age   int
}

// type person4 struct {
// 	fname string
// 	lname string
// 	age   int
// }

// Ex4 -Exported function to main - json Marshal
func Ex4() {

	p1 := person4{Fname: "jhon", Lname: "snow", Age: 30}
	p2 := person4{Fname: "bran", Lname: "stark", Age: 23}

	person := []person4{p1, p2}

	so := []string{"a", "c", "A", "Hello", "1a"}
	sort.Strings(so)

	fmt.Printf("Struct before marshal is - %v \n", so)

	// fmt.Printf("Struct before marshal is - %v \n", person)

	bs, err := json.Marshal(person)

	if err != nil {
		fmt.Printf("Error is %s \n", err)
	}

	fmt.Printf("Value of string is %s \n", bs)

	// personNew := []person4{}
	people := []person4{}

	s := `[{"Fname":"jhony","Lname":"neon","Age":3},{"Fname":"bran","Lname":"adams","Age":32}]`

	bs1 := []byte(s)

	fmt.Printf("Type of bs1 - %T \n", bs1)
	uerr := json.Unmarshal(bs1, &people)
	if uerr != nil {
		fmt.Print(uerr)
	}

	fmt.Printf("Value of unmarshal is %+v \n", people)

}
