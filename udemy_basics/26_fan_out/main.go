package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	c1 := make(chan int)
	c2 := make(chan int)

	go populate(c1) // put some values on the chan
	go fanInOut(c1, c2)

	for v := range c2 {
		fmt.Println(v)
	}
	fmt.Println("about to exit")
}

func populate(c chan int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c)
}

func fanInOut(c1, c2 chan int) {
	var wg sync.WaitGroup

	for v := range c1 {
		wg.Add(1)

		go func(v2 int) { // each value from our channel 1 will start a go ruttine
			c2 <- timeConsumingWork(v2) // after finishing a time consuming work result will be passed back to another channel
			wg.Done()
		}(v) // way of passing args into to anonymus function
	}

	wg.Wait()
	close(c2)
}

func timeConsumingWork(n int) int {
	time.Sleep(time.Microsecond * time.Duration(rand.Intn(500)))
	return n + rand.Intn(100)
}
