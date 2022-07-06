package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

// go has a interface build in for errors. Type error is present in GO!
// therefore to create our own errors, we just need to have types with Error methods, that return string
// type error interface {
// 	Error() string
// }

func main() {
	// checkPrintError()
	// inputWithCheck()
	// createFile()
	waysOfLogginErrors()
}

// checking errors in package
func checkPrintError() {

	// many packages return two values, where second is a error
	value, err := fmt.Println("Hello")

	if err != nil { // we should ALWAYS check the error closes to the place it could happen to help us debug it. This is very common pattern.
		// Live templates in go support completion of this with tab
		fmt.Println(err)
	}

	fmt.Println(value)
}

//multiple error checks
func inputWithCheck() {
	var answer1, answer2 string

	// scan will put answer provided via input into variable thatr address we provide
	_, err := fmt.Scan(&answer1) // it will also return error. First time we assign it
	if err != nil {
		panic(err)
	}

	_, err = fmt.Scan(&answer2) // next error check we can just re-assing the error to the variable
	if err != nil {
		// when we throw panic, we can use recover. there are also log.Panic
		panic(err)
	}

	fmt.Printf("answers: %v, %v", answer1, answer2)
}

// executing function with return
func createFile() {
	file, err := os.Create(("name.txt"))
	if err != nil {
		fmt.Println(err)
		return // just like in js we can create a guard clause to end function early
	}
	defer file.Close() // cool trick. We create this statment here so we wont forget about closing a file. But it will execute at the end of the function run

	read := strings.NewReader("Wassup")
	io.Copy(file, read)
}

// ways of logging a error
func waysOfLogginErrors() {
	// we have panic and fatal errors. panic errors can be recovered from. Fatal errors will shut down our program
	logFile, err := os.Create("log.txt")
	if err != nil { // error log also has a error check just in case :)
		fmt.Println(err)
		log.Println(err) // provides also date and time stamp
	}
	defer logFile.Close()  // we defer closing of the file
	log.SetOutput(logFile) // we set a output for our log file. Any logging done now will be dont to this file (log file is a address in memory)

	_, err = os.Open("no-file.txt") //provoke a error
	if err != nil {
		log.Println("Howdy, logging error", err) // will print it to the direction we defined
	}

	_, err = os.Open("no-file.txt")
	if err != nil {
		log.Panicln("Howdy, logging error: Panic", err) // panic will stop the execution and allow us to recover
	}

	_, err = os.Open("no-file.txt")
	if err != nil {
		log.Fatalln("Howdy, logging error: Fatal 3", err) // fatal will shut the whole program. Defer functions will also not run! It will throw error status 1
	}

}
