package channels

import "fmt"

// Ex4 - directional channel
func Ex4() {
	c := make(chan int, 2) // Bi Directional buffered Channel

	cs := make(chan<- int, 2) // send channel
	cr := make(<-chan int, 2) // receive channel

	fmt.Printf("Type Channel:  %T \n", c)
	fmt.Printf("Type Send Channel: %T \n", cs)
	fmt.Printf("Type Receive Channel: %T \n", cr)

	cr = c
	fmt.Printf("Type Receive Channel: %T \n", cr)
}
