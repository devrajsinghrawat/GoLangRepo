package exsoll1

import "fmt"

// Ex9 -Exported function to main
func Ex9() {
	a := map[string]int{"0xa": 10,
		"0xb": 20, "0xc": 30, "0xd": 40}
	// This help us to distingush between missing entry and zero value
	if v, ok := a["0xb"]; ok {
		fmt.Println("This is if check -", v, ok)
	}

	delete(a, "0xb")

	fmt.Println(a)

	v, ok := a["0xb"]
	fmt.Println(v, ok)
	fmt.Print("\nKey \t \t Value \n")
	for k, v := range a {
		fmt.Printf("%s \t \t %d \n", k, v)
	}

}
