package main

import (
	"fmt"
	"flag"
)

func main(){
	const ASCII_STARTING_POINT = 33
	const ASCII_END_POINT = 177
	upToPtr := flag.Int("upTo", 0, "up to what character we want to print")
	flag.Parse()

	upToValue := *upToPtr

	if(upToValue < ASCII_STARTING_POINT){
		fmt.Println("Number must be at least 33")
		return
	}

	for i := ASCII_STARTING_POINT; i <= upToValue; i++ {
		fmt.Printf("iterating on number: %v, which in ASCII/Unicode is: %#U \n", i, i)
	}

}
