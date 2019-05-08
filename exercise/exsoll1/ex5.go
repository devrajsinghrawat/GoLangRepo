package exsoll1

import "fmt"

const (
	a = 2014 + iota
	b = 2014 + iota
	c = 2014 + iota
)

// Ex5 -Exported function to main
func Ex5() {
	fmt.Print("\n", a, b, c)
}
