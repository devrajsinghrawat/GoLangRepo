package channels

import "fmt"

/*
Channel
----------
Channel in Go provides a connection between two goroutines, allowing them to communicate and synchronize the exchange 
of any resource that is passed through it. It is the channel’s ability to control the goroutines interaction that 
creates the synchronization mechanism. When a channel is created with no capacity, it is called an unbuffered channel. 
In turn, a channel created with capacity is called a buffered channel.

Buffered channels
-------------------
Buffered channels have a capacity and therefore can behave a bit differently. When a goroutine attempts to send a 
resource to a buffered channel and the channel is full, the channel will lock the goroutine and make it wait until 
a buffer becomes available. If there is room in the channel, the send can take place immediately and the goroutine 
can move on. When a goroutine attempts to receive from a buffered channel and the buffered channel is empty, 
the channel will lock the goroutine and make it wait until a resource has been sent. If the buffer is full or if there 
is nothing to receive, a buffered channel will behave very much like an unbuffered channel.

*/

// Ex2 - Channels buffer
func Ex2() {
	c := make(chan int, 2)

	c <- 40
	c <- 45
	// c <- 48        this will block the channel

	fmt.Print(<-c, <-c)

}
