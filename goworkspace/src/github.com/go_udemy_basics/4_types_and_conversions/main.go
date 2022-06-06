package main

import ("fmt")

// decaring our own type
type hotdog int

var a int
var b hotdog

func main(){
	fmt.Printf("#var b is of type %T", b)

	// we cant assing type hotdog to type int. Despite the fact underling type for hotdog is int. But we can conver it:
	a = int(b)
	fmt.Printf("#var a is of type %T", a)
}