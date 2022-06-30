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
	goToJSON()
	jsonTOGO()
}

// GO to json
func goToJSON(){
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

	// idiomatic code
	if err != nil{ // basic error handling using error returned
		fmt.Println(err)
	}

	fmt.Printf("go object transformed to json: %v \n", string(marskaledData) ) // we can convert it to slice of byte
}

const wildJSON = `{"_id": "62bde8888f44b29f7e1f118e",
    "index": 0,
    "guid": "71476ee4-5fdc-49a2-9799-c68945b5bcbf",
    "isActive": true}`

// we use string literals to map our json object to go struct
type order struct{
	Id string `json:"_id"`
	Index int `json:"index"`
	Guid string `json:"guid"`
	IsActive bool `json:"isActive`
}

// transform json to struct
func jsonTOGO(){
	// first we need to transform string to slice opf byte or uint8. This is often a step between tranforming data from/to strings
	byteSlice := []byte(wildJSON) 
	var incomingOrder order
	err := json.Unmarshal(byteSlice, &incomingOrder) // unmarshall will acept two arguments. slice of byte of the string and a address of the variable we want to store it
	// we always should print error right after the place it could happen for easier debbuging
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("json was parsed into %v,Guid is %v \n", incomingOrder, incomingOrder.Guid);
}