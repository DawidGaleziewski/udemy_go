package main

import "fmt"

// channels are advance form of using concurrency and controlling the flow
func main(){
	sucessfulBufferOne()
	unsucesfullBuffer()
}

// this will be blocked and will cause error
func pseudocodeBasicIDeaOfChannels(){
	c := make(chan int) // make a channel that we can put integers in
	c <- 42 // put number 42 into c
	fmt.Println(<-c) // take value OFF the c
	// as send and recive needs to happen at the same time when using a channel. It will block the execution of the code. If it cant happen the send will block the execution of the code untill recive is ready
	// IMPORTANT: channels BLOCK. If we start a chanel in a normal synchronus code block it will block our main thread // kill the program
}


func sucessfulBufferOne(){
	channelOne := make(chan int, 1) // second param is a buffer channel. We specify how many values can sit here

	go func(){ //gorutine starts
		channelOne <- 42  // this will block the gorutine running concurrently. Still the main function will continue
	}()

	fmt.Println(<-channelOne) // this will wait untill it can take the value off
}

func unsucesfullBuffer(){
	channelOne := make(chan int, 1)
	channelOne <- 42
	channelOne <- 64 // we cannot put more then one values in the buffer

	fmt.Println(<-channelOne)
}