package main

import (
	"fmt"
)

func main(){
	// pointer is pointing to a place in memory where the value is stored. Each place in memory has address. 
	age := 42

	fmt.Printf("this is a value of the variable age: %v, this is its address in memory (or pointer) %v \n", age, &age)

    // in order to return a address in memory we put & in front of the variable name
	agePtr := &age

	// a pointer is symbolised by *, we we have a *int (pointer pointing to a aint), *string (pointer pointing to a string)
	fmt.Printf("ptr value is %v and its type %T \n", agePtr, agePtr)

	// in order to get the value from a pointer we can use * symbol in front of a pointer
	fmt.Printf("pointer agePrt points to value: %v \n", *agePtr)

	// IMPORTANT *int is a type, while *agePtr is a operator
}

