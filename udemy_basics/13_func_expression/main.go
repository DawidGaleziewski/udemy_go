package main

import ("fmt")


func main(){
	// function expression can be declared just like in js
	// in go functions are used as first class citizens
	hello := func(){
		fmt.Println("Hello there")
	}


	hello()
	repeat5Times(func(index int){
		fmt.Printf("My index is %v \n", index * index)
	})

	returnedFunction := nestedReturns()
	fmt.Println(returnedFunction())

	evenValues := getEven(1,2,3,4,5,6,7,8);
	fmt.Printf("even values are %v: , add them together and you get: %v \n", evenValues , reduce(evenValues, func(acc int, value int) int{
		return acc + value
	}))
}




// or return annonymus functions
func nestedReturns() func()int{
	return func() int {
		return 451
	}
}

// we can also return functions from a function (HOCs?)

// callbacks in golang
func repeat5Times(action func(int)){
	for i := 0; i < 5; i++ {
		action(i)
	}
}

func add(x int, y int) int{
	return x + y
}


func getEven(xi ...int) []int{
	var result []int
	for _, v := range xi {
		if(v%2 == 0){
			result = append(result, v)
		}
	}

	return result
}

func reduce(arr []int, callback func(acc int, value int)int)  int{
	var accumulator int

	for _ , value := range arr {
		accumulator = callback(accumulator, value)
	}

	return accumulator
}