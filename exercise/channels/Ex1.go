package channels

import "fmt"


/* Channel
Channel in Go provides a connection between two goroutines, allowing them to communicate and synchronize the exchange 
of any resource that is passed through it. It is the channel’s ability to control the goroutines interaction that 
creates the synchronization mechanism. When a channel is created with no capacity, it is called an unbuffered channel. 
In turn, a channel created with capacity is called a buffered channel.  
*/

/* Unbuffered channels
Unbuffered channels have no capacity and therefore require both goroutines to be ready to make any exchange. 
When a goroutine attempts to send a resource to an unbuffered channel and there is no goroutine waiting to receive 
the resource, the channel will lock the sending goroutine and make it wait. When a goroutine attempts to receive from 
an unbuffered channel, and there is no goroutine waiting to send a resource, the channel will lock the receiving 
goroutine and make it wait. Synchronization is inherent in the interaction between the send and the receive. 
One can not happen without the other.
*/

// Ex1 - Channels
func Ex1() {
	c := make(chan int)

	// c <- 40               // This part of code alone will not work as channel blocks without synchronization 

	go func() {              // This is sender, sync is done using goroutine
		c <- 40
	}()


	fmt.Print(<-c)					// Receiver - This is main go routine
}
