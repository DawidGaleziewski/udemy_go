package main

import (
	"fmt"
)

func main(){

	// anonymus function in go. We can run it similar to js IIFE
	func(){
		fmt.Println("Hello from anonymus func")
	}()

	// and passing params
	func(name string){
		fmt.Printf("Hello there %v", name)
	}("General Kennobi")
}