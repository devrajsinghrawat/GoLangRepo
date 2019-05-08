package exsoll4

import "fmt"

var m map[string][]string

// Ex3 -Exported function to main, Map
func Ex3() {
	b := []string{"EOS", "STM", "XLM"}

	m = map[string][]string{"dev": []string{"BTC", "ETH"}, "raj": b}

	// Adding a record into map
	c := []string{"NEO", "STM", "0x", "DAI"}
	m["income"] = c

	fmt.Println(m)
	fmt.Println(m["income"])

	if _, ok := m["invest"]; ok {
		fmt.Printf("Map key invest has a valid value \n")
	} else {
		fmt.Printf("Map key invest has a invalid value \n")
	}

	if v, ok := m["income"]; ok {
		fmt.Printf("Map key income has a valid value %v \n", v)
	} else {
		fmt.Printf("Map key income has a invalid value %v \n", v)
	}
}
