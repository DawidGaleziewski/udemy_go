package main

import "fmt"

func main() {
    selectUseCaseMobsterOperation()
	comaOkIdiom()
}

// example of function sending things thru multiple channels
func sender(ordersChannel chan<- string, complainsChannel chan<- string, sanepidCameAlarmChannel chan<- bool) {
	orders := []string{"milk", "beef", "candy"}
	complains := []string{"milk is bad", "shop is dirty"}

	for _, order := range orders {
		ordersChannel <- order
	}

	for index, complain := range complains {
		if(index > 0){
			fmt.Println("closing the channel")
			close(complainsChannel)
			break
		}
		complainsChannel <- complain
	}

	sanepidCameAlarmChannel <- true
}

func reciver(ordersChannel <-chan string, complainsChannel <-chan string, sanepidCameAlarmChannel <-chan bool) []string {
	var allInfo []string
	var orders []string
	var complains []string

	for { // doing wild loop that will make sure we listen for incoming channels
		select { // select is used similar to switch but for gorutines
			case value := <-ordersChannel:
				allInfo = append(allInfo, value)
				orders = append(orders, value)
			case value, ok := <-complainsChannel:					
				if !ok { // we can use second value "ok" to check if channel was not closed yet
					fmt.Println("something is not ok, channel is mute! \n")
				} else {
					allInfo = append(allInfo, value)
					complains = append(complains, value)
				}

			case alarm := <-sanepidCameAlarmChannel:
				fmt.Printf("Copers are here! Drop the cargo and run! %v \n", alarm)

				return orders  // this will break the wild loop
		}

	}

}

// selct allows us to orchiestrate multiple channels using syntax same as switch statment
func selectUseCaseMobsterOperation(){
	ordersChannel := make(chan string)
	complainsChannel := make(chan string)
	sanepidCameAlarmChannel := make(chan bool)

	go sender(ordersChannel, complainsChannel, sanepidCameAlarmChannel)

	cargo := reciver(ordersChannel, complainsChannel, sanepidCameAlarmChannel)

	fmt.Printf("droping all cargo: %v", cargo)
}

func comaOkIdiom(){
	channel := make(chan int)

	go func(){
		channel <- 11
	}()

	value, isChannelOpened := <- channel // we can check if channel is opened by checking second value returned from it when taking value of the channel

	fmt.Println(value, isChannelOpened)
}
