package main

import ("fmt")

func main(){
	exercise1()
	exercise2()
	exercise3()
	exercise4()
	exercise5()
}


func exercise1(){
	x,y,z := 42, "James Bond", true
	fmt.Println(x,y,z)
	fmt.Printf("%v is of type %T", y ,y)
}

var x int
var y string
var z bool

func exercise2(){
	fmt.Println("zeroed values: ", x,y,z)
}




func exercise3(){
	x = 42
	y = "Todd Howard"
	z = true


	s := fmt.Sprintf("variables are of types: %T, %T, %T",x,y,z)

	fmt.Println(s)
}

type hop1 int
type hop2 hop1
type hop3 hop2
var xs hop3

func exercise4(){
	fmt.Printf("x is of type: %T \n", xs)
	xs = 42
	fmt.Printf("x is of type: %T \n", xs)
}

type dog string
var myPet dog
var name = "fafik"


func exercise5(){
	myPet = dog(name)
	fmt.Printf("myPet is of tyep %T and its value is %v \n", myPet, myPet)
}