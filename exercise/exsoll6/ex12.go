package exsoll6

import "fmt"

// Ex12 -Exported function to main - callback : pass a func into a func as a argument
func Ex12() {

	foo12(callback, "World ! \n")

}

func foo12(f func(string), s string) {
	s1 := fmt.Sprint("Hello-", s)
	f(s1)
}

func callback(s string) {
	fmt.Printf(s)
}
