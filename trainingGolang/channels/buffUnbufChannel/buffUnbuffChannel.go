package main

import (
	"fmt"
	"time"
)

func main() {
	mainUntukUnbufferedChannel()
	mainUntukBufferedChannel()
}

// Contoh Unbufferd Channel

func mainUntukUnbufferedChannel() {
	c1 := make(chan int)
	go func(c chan int) {
		fmt.Println("func goroutine starts sending data into the channel")
		c <- 10
		fmt.Println("func goroutine after sending data into the channel")
	}(c1)

	fmt.Println("main goroutine sleeps for 2 seconds")
	time.Sleep(time.Second * 2)

	fmt.Println("main goroutine starts receiving data")
	d := <-c1
	fmt.Println("main goroutine received data:", d)

	close(c1)
	time.Sleep(time.Second)
}

// Contoh buffered Channel

func mainUntukBufferedChannel() {
	c1 := make(chan int, 3)

	go func(c chan int) {
		for i := 1; i <= 5; i++ {
			fmt.Printf("func goroutine #%d starts sending data into the channel\n", i)
			c <- i
			fmt.Printf("func goroutine #%d after sending data into the channel\n", i)
		}

		close(c)

	}(c1)

	fmt.Println("main goroutine sleeps 2 seconds")
	time.Sleep(time.Second * 2)
	for v := range c1 { // v = ‹- c1
		fmt.Println("main goroutine received value from channel:", v)
	}
}
