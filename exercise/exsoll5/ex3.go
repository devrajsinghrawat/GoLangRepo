package exsoll5

import "fmt"

// Ex3 -Exported function to main - Struct, map and slice
func Ex3() {

	t := struct {
		tokens []string
		bal    map[string]uint
	}{
		tokens: []string{
			"EOS",
			"CATS",
			"DEVS",
		},
		bal: map[string]uint{
			"0xA": 100,
			"0xB": 200,
			"0xC": 300,
		},
	}

	fmt.Printf("%v \n", t)

}
