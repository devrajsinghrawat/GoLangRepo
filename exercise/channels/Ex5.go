package channels

import (
	"fmt"
)

// var wg sync.WaitGroup

// Ex5 - directional channel
func Ex5() {
	c := make(chan int) // Bi Directional Channel
	// cs := make(chan<- int, 2) // send channel
	// cr := make(<-chan int, 2) // receive channel

	//	wg.Add(2)
	go send(c)
	receive(c)

	//	wg.Wait()
	fmt.Printf("About to exit \n")
}

func send(cs chan<- int) {
	cs <- 22
	//	wg.Done()
}

func receive(cr <-chan int) {
	fmt.Printf("Receive Value %v \n", <-cr)
	//	wg.Done()
}
