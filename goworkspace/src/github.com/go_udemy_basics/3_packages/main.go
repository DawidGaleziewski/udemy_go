package main

import ("fmt")

func main() {
	fmtPackagePlayground()
	printAndCatchError()
	printAndDoNothingWithError()
}

func printAndCatchError(){
	value, err := fmt.Println("Hello") // most packages and their method return tuple of value error. Error is nil here. We can catch errors by assigning them to variables
	fmt.Println(value, err)
}

func printAndDoNothingWithError(){
	value, _ := fmt.Println("hello 2") // go will force our hand to do something with variables. But we can assign the value to _ if we do not want to use it
	fmt.Println(value) 

}

func fmtPackagePlayground(){
	numVal, strVal, boolVal := 5, "test", true

	fmt.Println("## printing types of arguments provided")
	fmt.Printf("#type is %T for value: %v", numVal, numVal)
	fmt.Printf("#type is %T for value:", strVal)
	fmt.Printf("#type is %T for value:", boolVal)

	fmt.Println("## we have also \"verbs\" for diffrent formats like binary")

	luckyNum := 10
	fmt.Printf("Binary for value %v in binary is %b", luckyNum, luckyNum)
}
