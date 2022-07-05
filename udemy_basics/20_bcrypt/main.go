package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	passwd := "test1234!"
	bs, err := bcrypt.GenerateFromPassword([]byte(passwd), bcrypt.MinCost)

	if(err != nil){
		fmt.Println(err)
	}

	fmt.Printf("hashed passwrord is: %v \n", string(bs))
}