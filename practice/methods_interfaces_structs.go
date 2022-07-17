package main

import "fmt"

// read on: https://www.ardanlabs.com/blog/2014/05/methods-interfaces-and-embedded-types.html
type Human struct {
	name    string
	surname string
	age     int
	hobbies []string
}

type Dog struct {
	name string
	age  int
}

func (h Human) Walk(destination string) {
	fmt.Printf("%v is walking slowly towards %v \n", h.name, destination)
}

func (h Superhuman) Walk(destination string){
	fmt.Printf("%v Ziuuum (superhuman speed) %v \n", h.name, destination)
}

func (d Dog) Walk(destination string) {
	fmt.Printf("%v is walking towards %v \n", d.name, destination)
}

type Walker interface { // covention to call interface with -er suffix
	Walk() // in standard library we rarly will see a interface with more then two method so this is ok.
}

// read more on good practice when using recivers: https://github.com/golang/go/wiki/CodeReviewComments#Receiver_Type
// and naming recivers: https://github.com/golang/go/wiki/CodeReviewComments#Receiver_Names
func (h *Human) Wait(years int) { // if we want to modify something, we will have to use a
	// when the receiver is not a pointer, the method is operating against a copy of the receiver value. Therefore we cant mutate it
	h.age = h.age + years
}

// composition by embedding types into structs
type Superhuman struct {
	Human; // all values from human will be promoted to the same level as Superhuman
	superpowers []string;
	age int;
}

func main() {
	fmt.Println("test")

	mark := Human{
		name:    "Mark",
		surname: "Henderson",
		age:     50,
		hobbies: []string{"singing", "walking"},
	}

	mark.Walk("Fridge")

	mark.Wait(50)
	fmt.Println("mark is now", mark.age)

	// playing with composition
	clark := Superhuman{
		Human: Human{
			name: "Clark",
			surname: "Kent",
			age: 500,
			hobbies: []string{"photography"},
		},
		age: 1,
		superpowers: []string{"flying", "strength", "laser sight"},
	}

	fmt.Println(clark.Human.name)
	clark.Human.Walk("fridge") // when we embeed a type that has methods. We can call it from inner type or directly from the outer type (they get promoted). Methods will hovewer, revice as arguments only the inner type
	clark.Walk("to Loris") // it is still able to use methods from human.
	fmt.Println("clarks age", clark.age) // all properties and methods are "hoisted" to the top level
}

// !IMPORTANT 
/**
"When we embed a type, the methods of that type become methods of the outer type, but when they are invoked, the receiver of the method is the inner type, not the outer one." - Effective Go
**/

// If we have two implementation of method that is named the same. Like walk above. The inner implementation wont get promoted. We can access seperate implementations by clark.Walk(), clark.Human.Walk()
