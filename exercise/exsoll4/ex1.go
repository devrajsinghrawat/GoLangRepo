package exsoll4

import (
	"fmt"
	"strconv"
)

var bal = []int{10, 20, 30, 40, 50, 60, 70, 80, 90}

// Ex1 -Exported function to main
func Ex1() {
	balanceArray := [5]int{1, 2, 3, 4, 5} // array

	// bal = balanceArray[:3] // Slice is initlized using a array
	// for i := 0; i < len(bal); i++ {
	// 	bal[i] = bal[i] * 10
	// }

	balslice := bal[:4]
	fmt.Printf("1 - %v\n", balslice)

	balslice = bal[4:]
	fmt.Printf("2 - %v\n", balslice)

	balslice = bal[2:6]
	fmt.Printf("3 - %v\n", balslice)

	fmt.Printf("type of array %T \n", balanceArray)
	fmt.Printf("type of slice %T \n", bal)
	fmt.Printf("capacity of slice %v \n", cap(bal))

	for j, v := range bal {
		fmt.Printf("%v %v\n", j, v)
	}

	con, _ := strconv.ParseInt("100", 10, 64) // converting string to int
	bal = append(bal, int(con))
	fmt.Printf("A - %v\n", bal)
	fmt.Printf("updated capacity of slice %v \n", cap(bal))

	bal = append(bal, 110, 120, 130)
	fmt.Printf("B - %v\n", bal)
	fmt.Printf("updated capacity of slice %v \n", cap(bal))

	bal = append(bal, balslice...)
	fmt.Printf("C - %v\n", bal)
	fmt.Printf("updated capacity of slice %v \n", cap(bal))

	bal = append(bal[2:5], bal[8:]...)
	fmt.Printf("D - %v\n", bal)
	fmt.Printf("updated capacity of slice %v \n", cap(bal))
}
