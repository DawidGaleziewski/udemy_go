package main

import (
	"fmt"
	"flag"
)

func main(){
	charPtr := flag.String("char", "", "char passed")
	numPtr := flag.Int("num", 0, "number passed")
	shiftPtr := flag.Int("shiftBy", 0, "how many positions we want to shift the numbers")

	flag.Parse()

	shiftBy := *shiftPtr

	char := *charPtr
	if(len(char) > 0){
		charSliceOfByte := []byte(char)[0] << shiftBy
		fmt.Printf("for character: \"%v\", when shifted by: %v, its decimal encoding is: %v,\n in binary: %b,\n hexadecimal: %x, \n", char, shiftBy, charSliceOfByte, charSliceOfByte, charSliceOfByte)
		return
	}

	num := *numPtr << shiftBy
	if(num > 0){
		fmt.Printf("for number passed: %v,\n when shiftedBy %v in binary numeric system it is: %b \n, in hexadecimal: %#x", num ,shiftBy, num, num)
	}
}