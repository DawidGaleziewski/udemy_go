package main

import (
	"fmt"
	"sync"
)

func main() {
	fanInPattern()
}

// fan out go rutines to not declared number of gorutines
// fan in is a design pattern we can use to join multiple channels values into one
func fanInPattern() {
	faninChannel := make(chan string)

	namesChannel := make(chan string)
	surnamesChannel := make(chan string)

	// Start sending items from our channels
	go send(namesChannel, surnamesChannel)

	// fanin two channels into one
	go recive(namesChannel, surnamesChannel, faninChannel)

	for v := range faninChannel {
		fmt.Println(v)
	}

}

func send(namesChan, surnamesChan chan<- string) {
	listOfNames := []string{"Dawid", "Marek", "Tytus"}
	listOfSurnames := []string{"Galeziewski", "Kowal", "Serafing"}

	for _, v := range listOfNames {
		namesChan <- v
	}
	close(namesChan)

	for _, v := range listOfSurnames {
		surnamesChan <- v
	}
	close(surnamesChan)
}

func recive(names, surnames <-chan string, fanin chan<- string) {
	// use wait group to wait for ending two channels in gorutines
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for v := range names {
			fanin <- v
		}
		wg.Done()
	}()

	go func() {
		for v := range surnames {
			fanin <- v
		}
		wg.Done()
	}()

	wg.Wait()
	close(fanin)
}
