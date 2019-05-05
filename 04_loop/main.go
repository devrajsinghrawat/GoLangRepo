package main

import "fmt"

func main() {
	var i int
	for ; i < 1127; i++ {
		fmt.Printf("%d \t %b \t %x \t %q \n", i, i, i, i)
	}
}
