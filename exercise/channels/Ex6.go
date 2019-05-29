package channels

import (
	"fmt"
)

// Ex6 - directional channel
func Ex6() {
	c := make(chan int) // Bi Directional Channel

	// go sendCH(c)

	go func() { // Send
		defer close(c)
		for index := 0; index < 5; index++ {
			c <- index
		}
		// close(c)
	}()

	// Receive
	for r := range c {
		fmt.Printf("Receive Value : %v \n", r)
	}

	fmt.Printf("About to exit \n")
}

// func sendCH(cs chan<- int) {
// 	for index := 0; index < 100; index++ {
// 		cs <- index
// 	}

// 	close(cs)
// }

// func receive(cr <-chan int) {
// 	fmt.Printf("Receive Value %v \n", <-cr)
// 	//	wg.Done()
// }
