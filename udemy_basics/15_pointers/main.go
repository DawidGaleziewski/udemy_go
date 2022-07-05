package main

import (
	"fmt"
)


func main(){
	basicPoinetOperation()
	usePointerToChangeValuesOutOfScope()

}

// & gives yoi the address
// * gives you the value stored at the address
func basicPoinetOperation(){
	// pointer is pointing to a place in memory where the value is stored. Each place in memory has address. 
	age := 42

	fmt.Printf("this is a value of the variable age: %v, this is its address in memory (or pointer) %v \n", age, &age)

    // in order to return a address in memory we put & in front of the variable name
	agePtr := &age

	// a pointer is symbolised by *, we we have a *int (pointer pointing to a aint), *string (pointer pointing to a string)
	fmt.Printf("ptr value is %v and its type %T \n", agePtr, agePtr)

	// in order to get the value from a pointer we can use * symbol in front of a pointer
	fmt.Printf("pointer agePrt points to value: %v \n", *agePtr)

	// IMPORTANT *int is a type, while *agePtr is a operator.
	// Operation of getting the value of the pointer is called DEREFFERENCING the address

	// we can overide the value in memory by DEREFERENCING it
	*agePtr = 500

	fmt.Printf("age is now: %v \n", age)
}

func usePointerToChangeValuesOutOfScope(){
	fmt.Println()
	// when we pass value to a function, and we modify it it wont change outside of its scope
	value := "#global scope"
	fmt.Printf("Value in global scope is: %v \n", value)
	passByValue(value)
	fmt.Println()
	fmt.Printf("After it was modified in function Value in global scope is: %v \n", value) // wont be affected

	fmt.Println()
	fmt.Println()
	passByAddress(&value) // passing value by address
	fmt.Printf("After it was modified in function by ADDRESS Value in global scope is: %v \n", value) // wont be affected

}

func passByValue(value string){
	value ="#passByvalue"
	fmt.Printf("#value inside passByValue %v \n", value)
}

func passByAddress(address *string){
	// this is how we can mutate a value
	*address = "###passByAddress"
}