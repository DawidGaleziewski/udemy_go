package main

import "fmt"

func main() {
	channel := make(chan int)

	go sender(channel)

	// we range over the values we recive from the channel
	for value := range channel {// range works beatifuly with channel. It will pull values one by one from a channel untill it is closed
		fmt.Println(value)
	}
}

func sender(channel chan<- int) {
	for i := 0; i < 100; i++ {
		someValue := i
		channel <- someValue // we put some values on a looop to the channel
	}

	close(channel) // once we are done with the channel we want to close it. If we wont the range will be waiting
}