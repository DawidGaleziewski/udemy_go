package main

import ("fmt")

// closing variable scope so that it is narrower.
func main(){

	{ // in golang we can create simply a codeblock and scope variables to it
			isInScope := true
			fmt.Printf("is in scope of brackets: %v", isInScope)
	}

	//This will throw error as it is undefined
	//fmt.Printf("is otside of scope: %v", isInScope)

	incrementByTwo := incrementBy(5)
	incrementByTen := incrementBy(10) // functions returned point to diffrent internalState. Variable internlState points to diffrent place in memory. Those functions have diffrewnt scope

	fmt.Printf("increment fn 1 %v \n", incrementByTwo())
	fmt.Printf("increment fn 1 %v \n", incrementByTwo())
	fmt.Printf("increment fn 1 %v \n", incrementByTwo())
	fmt.Printf("increment fn 1 %v \n", incrementByTwo())

	fmt.Printf("increment fn 2 %v \n", incrementByTen())
	fmt.Printf("increment fn 2 %v \n", incrementByTen())
	fmt.Printf("increment fn 2 %v \n", incrementByTen())

}


// using closure to create a function with internal state
func incrementBy(value int) func() int {
	var internalState int // just like in js. DUe to the fact it is used in another functaion that will be returned. This is not garbage collected

	return func() int{
		internalState += value
		return internalState
	}
}
