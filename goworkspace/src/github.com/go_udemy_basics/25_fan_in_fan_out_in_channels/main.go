package main

func main() {
	fanInPattern()
}

// fan out go rutines to not declared number of gorutines
// fan in is a design pattern we can use to join multiple channels values into one
func fanInPattern() {
	faninChannel := make(chan string)

	go send(namesChannels, surnamesChannel)
}

type person struct {
	name          string
	surname       string
	likesChannels bool
}

const people = []person{}

func send(names, surnames <-chan string) {

}
