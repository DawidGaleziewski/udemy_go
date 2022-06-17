package main

import ("fmt"
"runtime")

// decaring our own type
type hotdog int

var a int
var b hotdog

func main(){
	fmt.Printf("#var b is of type %T \n", b)

	// we cant assing type hotdog to type int. Despite the fact underling type for hotdog is int. But we can conver it:
	a = int(b)
	fmt.Printf("#var a is of type %T \n", a)

	amIAwesome()
	numericTypes()
	architecture()
}


// Booleans
var isAwesome bool

func amIAwesome(){
	fmt.Printf("he is awesome: %v \n", isAwesome)
}

// Main numaric uses in go. If we do not care about space so much. int will store whole numbers. Float will have decimal
var numberOfCars int // cool thing about int is it will determine and optimise if the value should be 32 or 64 after compilation
var litersOfGas float64

// U int will go from 0 to 255
var spareTickets uint8

//ints will go from -128 to 127
var cardBalance int8

var numberofPixels byte // byte is same as uint8

func numericTypes(){
	numberOfCars = 2
	litersOfGas = 5.99

	fmt.Printf("%v cars will require %v liters of gasoline \n", numberOfCars, litersOfGas)
}

// checking architecture you are running
func architecture(){
	fmt.Printf("os is: %v, and architecture is: %v \n", runtime.GOOS, runtime.GOARCH)
}