package goroutines

import (
	"fmt"
	"time"
)

func GoRoutines() {

	fmt.Println("A goroutine is a lightweight thread managed by the Go runtime.")
	go say("\tworld")
	say("\thello")

	fmt.Println("Channels are a typed conduit through which you can send and receive values.")
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c) //start goroutine with half of the ints, the sum will be sent via the channel
	go sum(s[len(s)/2:], c) //start goroutine with half of the ints, the sum will be sent via the channel
	x, y := <-c, <-c        // receive from channel

	fmt.Println("\t", x, y, x+y)

	fmt.Println("Channels can be buffered. Provide the buffer length as the second argument.")
	bufCh := make(chan int, 2) // buffered channel, length 2
	bufCh <- 1                 //put 1
	bufCh <- 2                 //put 2

	// bufCh <- 3 this would cause an error
	fmt.Println(<-bufCh) //read
	fmt.Println(<-bufCh) //read

	fmt.Println("A sender can close a channel to indicate that no more values will be sent")
	fmt.Println("Receivers can test whether a channel has been closed by assigning a second parameter to the receive expression.")

	c2 := make(chan int)
	go foo(5, c2)
	for ii := range c2 {
		fmt.Println("\t int sent", ii)
	}

	fmt.Println("The select statement lets a goroutine wait on multiple communication operations.")
	c3 := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			c3 <- i*i + i
		}
		quit <- 9
	}()
	bar(c3, quit)

}

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s { //range over a channel to recieve
		sum += v //sum the ints
	}
	c <- sum // send sum to c
}

func foo(n int, c chan int) {
	for i := 0; i < n; i++ {
		c <- i * i //send an int
	}
	close(c) //close channel after i ints sent
}

func bar(c chan int, quit chan int) {

	for {
		select {
		case x := <-c:
			fmt.Println("\trecieved ", x)
		case <-quit:
			fmt.Println("\tquit")
			return
		default:
			fmt.Println("\tdefault")
		}
	}
	fmt.Println("\tafter for loop ")
}
