package exsoll6

import (
	"fmt"
)

// Ex7 -Exported function to main - Defer exercise
func Ex7() {
	defer foo7()
	fmt.Printf("Hello EX7 \n")

}

func foo7() {
	defer func() {
		fmt.Printf("this is deferred anonymous function \n")
	}()
	fmt.Printf("this printing has been deferred \n")
}
