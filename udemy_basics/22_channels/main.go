package main

import "fmt"

// channels are advance form of using concurrency and controlling the flow
func main(){
	//sucessfulBufferOne()
	//unsucesfullBuffer()
	//directionalChannels()
	channelsWhenCommunicatingBetweenFunctions()
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
	channelOne := make(chan int, 1) // second param is a buffer channel. We specify how many values can sit here.
	// PRO tip: mostly stay away from buffer channels

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

// by default chanels are bidirectional. But we can create directional channels that only send or recive
func directionalChannels(){
	// by default the channel will be bidirectional
	bidirectionalChannel := make(chan int)
	fmt.Printf("channel %v and its type: %T \n", bidirectionalChannel, bidirectionalChannel)

	sendOnlyChannel := make(chan <- string) // chan <- we can only send to a channel
	sendOnlyChannel <- "test"
	fmt.Printf("channel %v and its type: %T \n", sendOnlyChannel, sendOnlyChannel)

	reciveOnlyChannel := make(<- chan string) // recive only channel
	fmt.Printf("channel %v and its type: %T \n", reciveOnlyChannel, reciveOnlyChannel)

	// we can convert from bidirectional channel to directional. But not the other way arround (from general to more strict)
}

// as we can go from general to strict type definition. We can reduce the scope of what can be done with the channel, making it only send
func sender(senderChannel chan<- string){
	senderChannel <- "Howdy there from sender"
}

//  we can reduce the scope for channel only to recive the data
func reciver(reciverChannel <-chan string){
	fmt.Printf("recived data inside reciver: %v", <-reciverChannel)
}

// practical work with channels
func channelsWhenCommunicatingBetweenFunctions(){
	bidirectionalChannel := make(chan string)

	go sender(bidirectionalChannel) // spin up goruttine and pass a channel to it. It is biddirectional but inside of the function it can be only use to send data

	reciver(bidirectionalChannel) // pass in the same channel for communication. important we dont want to use goruttine here. Otherwise gorutine will start and the main will finish/close/we wont see anything. However if we pass a notmal sync function. It will wait for the code to finish as the channel WILL BLOCK THE EXECUTION.
	
	// we could also take out the value from channel like this:
	//fmt.Printf("recived data inside reciver: %v", <-bidirectionalChannel)
}