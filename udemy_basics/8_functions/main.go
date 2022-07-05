package main

import (
	"fmt"
)



func main(){
	fmt.Printf("result is: %v \n", multiplyBy2(2));

	ok, len1, len2 := compareLengths("Dave", "Matt")
	fmt.Printf("%v %v %v \n", ok, len1, len2)
	
	veriadicParams(1,2,3,4,5,6)

	dancingPeople := []string{"Mike", "Dave", "Jane", "Dustin"}
	// just like in js we can deconstruct a slice
	fmt.Println(formatNames(dancingPeople...)); 
}

// parameters are when function is declered, arguments are when function is called
// function syntax: func (r reciver) identifier(parameters) (returns(s)) {...}
// EVERYTHING in go is passed by VALUE
func multiplyBy2(value int) int {
	var result = value * 2
	return result; // if we have a return type we have to specify its type
}

// unlike js we can return multiple values from a function
func compareLengths(value1 string, value2 string) (bool, int, int){
	length1 := len(value1)
	length2 := len(value2)
	isSameLength := length1 == length2

	// NO problem with multiple values!
	return isSameLength, length1, length2
}

// just like in js we can use veriadic parameters. We just have to declare hose with ... operator
func veriadicParams(x ...int){
	fmt.Printf("x is value %v and its type is %T \n", x, x)
	for i, v := range x {
		fmt.Printf("veriadic param: %v at index %v \n", v, i)
	}
}

func formatNames(names ...string) string {
	var result string 
	for _, v := range names {
		result += v + "\\(^.-)/";
	} 

	return result
}