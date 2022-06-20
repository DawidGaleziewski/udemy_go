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
	stringType()
	convertingCharsToDiffrentCodingSchemes()
	numericSystems()
	constants()
	bitShifting()
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

var name string
func stringType(){
	name = "hello there" // can use doube quotes
	name = `hello 
	   there` // can use backtiques to eqpress spaces ans special chars with template literals

	fmt.Println(name)

	name = "Dave"
	// converting string to slice of bytes
	sliceOfBytes := []byte(name)

	fmt.Printf("slice of bytes value is (this will return values for ascii characters) %v", sliceOfBytes)
}

func convertingCharsToDiffrentCodingSchemes(){
	sentance := "nice day it is!"

	fmt.Println("#Translate chars to utf-8:")
	for i := 0; i < len(sentance); i++{
		var char = sentance[i]
		fmt.Printf("at position %v, char %v in utf-8 is: %#U \n",i, char, char )
	}


	fmt.Println("#Translate chars to hexadecimal:")
	for i := 0; i < len(sentance); i++{
		var char = sentance[i]
		fmt.Printf("at position %v, char %v in hexadecimal is: %#x \n",i, char, char )
	}
}

func numericSystems(){
	age := 32
	fmt.Printf("my age in decimal: %v, binary: %b, hexadecimal: %#X", age, age, age)
}



// Iota is a predeclared identifier used in constant declarations. It will incremant from declaration to declaration
const (
	packag1 = iota 
	package2 = iota
    package3 = iota
)

// Iota increments per one const declaration
const (
	packag4 = iota // we can also only declare iota once on the first const and it will be used on the rest og those
	package5
    package6
)

func constants(){
	// untyped constants - compiler will figure out the type of constant. We leave compiler the room, for interpretation
	const myName = "Dawid"
	const myAge = 42
	const chargePerHour = 42.99

	// Typed constants - when we want to be sure a type we want will be assigned
	const myName2 string = "Jan"
	const myAge2 int = 18

	fmt.Printf("%T\n %T\n %T\n" , myName, myAge, chargePerHour)

	fmt.Printf("Iotas are %v, %v, %v \n", packag1, package2, package3)
}

// Creative way of using iota to increment const declarations
const (
	_ = iota // we throw first iota which is 0 into the void, so that compile shuts the fuck up
	// kb = 1024
	kb = 1 << (iota * 10) // we shift one to the left 10 times * 1 (iota)
	mb = 1 << (iota * 10)
	gb = 1 << (iota * 10)
)

func bitShifting(){
	x := 2
	fmt.Printf("%d\t\t%b\n",x,x)

	// << is a special operator for shifting
	y := x << 1;
	fmt.Printf("%d\t\t%b\n",y,y)
}