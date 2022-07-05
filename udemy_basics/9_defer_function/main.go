package main

import (
	"fmt"
)

func main(){
	surroundingFunction()
}

// defer execution of function untill surrounding function executed a return statment or corresponding gorutine is panicking
func surroundingFunction(){
	defer deferedFunc() // defered function will run at the end
	plainFunction()
	
}

func deferedFunc(){
	fmt.Println("1. from deffered function")
}


func plainFunction(){
	fmt.Println("2. plain function")
}

