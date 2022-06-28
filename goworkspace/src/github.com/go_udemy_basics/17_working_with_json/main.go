package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int
}

// godoc.org/encoding/json
func main() {
	// composite literal
	p1 := person{
		First: "Dave",
		Last:  "Galeziewski",
		Age:   31,
	}

	p2 := person{
		First: "Mike",
		Last: "Black",
		Age: 50,
	}

	people := []person{p1 , p2}
	fmt.Printf("People slice is: %v \n", people)

	marskaledData, err := json.Marshal(people)

	if err != nil{ // basic error handling using error returned
		fmt.Println(err)
	}

	fmt.Printf("marshaled data is: %v \n", string(marskaledData) ) // we can convert it to slice of byte
}