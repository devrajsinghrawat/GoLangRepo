package channels

import "fmt"

// Ex7 - Select channel
func Ex7() {
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan bool)

	// Send
	go sendChannel(even, odd, quit)
	receiveChannel(even, odd, quit)

}

func sendChannel(e, o chan<- int, q chan<- bool) {
	// defer close(e)
	// defer close(o)
	// q <- 9999
	for index := 0; index < 10; index++ {
		if index%2 == 0 {
			e <- index
		} else {
			o <- index
		}
	}
	close(q)
}

func receiveChannel(e, o <-chan int, q <-chan bool) {
	for {
		select {
		case v := <-e:
			fmt.Printf("Receive Even: %v \n", v)
			//continue
		case v := <-o:
			fmt.Printf("Receive Odd : %v \n", v)
			//continue
		case i, ok := <-q:
			if !ok {
				fmt.Printf("Receive Quit: %v %v \n", ok, i)
				return
			} else {
				fmt.Printf("Receive Quit: %v %v \n", ok, i)
			}

		}
	}
}
