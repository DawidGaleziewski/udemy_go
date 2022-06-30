package main

import (
	"fmt"
	"io"
	"os"
	"sort"
)

// !to watch again to get it
// type Writer is a interface. Therefore ANY other type with method "Write(p []byte)(n int, err error)" will be also of type writer
func main(){
	usingBuildInFunctionsThatAcceptWiter()
	usingSortLib()
}

func usingBuildInFunctionsThatAcceptWiter(){
	fmt.Fprintln(os.Stdout, "Hello there")
	io.WriteString( os.Stdout,"hello")
}

func usingSortLib(){
	nums := []int{5,3,4,6,10,200,0}
	sort.Ints(nums) // will mutate the variable and sort it

	fmt.Printf("sorted ints asre %v \n", nums)
}