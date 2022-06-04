package main

import ("fmt")

func main() {
	printAndCatchError()
}

func printAndCatchError(){
	value, err := fmt.Println("Hello") // most packages and their method return tuple of value error. Error is nil here. We can catch errors by assigning them to variables
	fmt.Println(value, err)
}

func printAndDoNothingWithError(){
	value, _ := fmt.Println("hello 2") // go will force our hand to do something with variables. But we can assign the value to _ if we do not want to use it
	fmt.Println(value) 
}
