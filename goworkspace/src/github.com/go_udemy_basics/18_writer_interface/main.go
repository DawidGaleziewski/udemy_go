package main

import (
	"fmt"
	"io"
	"os"
)

// !to watch again to get it
// type Writer is a interface. Therefore ANY other type with method "Write(p []byte)(n int, err error)" will be also of type writer
func main(){
	usingBuildInFunctionsThatAcceptWiter()
}

func usingBuildInFunctionsThatAcceptWiter(){
	fmt.Fprintln(os.Stdout, "Hello there")
	io.WriteString( os.Stdout,"hello")
}

