package main

import (
	"fmt"
)

//POLIMORPHISM: Notice golang allows us to declare methods with SAME NAME but different RECIVER. So those methode can be named the same but OPERATE on different types
func (s man) speak(){
	fmt.Printf("Hello there my name is: %v and i am a: %T \n", s.name, s)
}

func (s homonculus) speak(){
	fmt.Printf("I was created by %v, my name is: %v and i am a: %T \n", s.creator, s.name, s)
}
// interface allows us to define BEHAVIOUR and allows us to use polymorphism
// Interfaces are very useful with packages such as Writer. Any type that has a method Write will also be of type writer. Many packages such as http package accept writer as a argument. Those packages via polymorphism can implement this method diffrently, dpeneding on the the type
type human interface {
	speak() // anybody who has methid speak will ALSO be of type human
}

//Notice: "keyword identifier type" pattern that is used in golang:
// type human interface
// type car struct
// var x int

// lets go back to methods for a moment and notice this. Notice we assigned a method also to woman by reciver in the fucntion declaration
type man struct{
	name string;
	age int;
} 

type homonculus struct{
	name string;
	creator string;
}

// we can use our newly atatched
func whereAreYou(h human){
	fmt.Printf("Hey where are you %v? %T \n",h,h)

	// we can use type assertion in order to figure out what type the passed value is. We use value.(type) syntax for this. We need to go to underlying type of the interface
	switch h.(type) { // this is assert (zapewniać) that type is a X
		case man:
			fmt.Printf("Please come back you have nothing to worry about mr %v! \n", h.(man).age) // as we checked this is a man type in switch, we can assert this type here and use its value
		case homonculus:
			fmt.Printf("Shoot the homonculus!! Go back to your maker %v", h.(homonculus).creator)
	}
}

func main(){
	somebody := man{
		name: "Josh",
		age: 32,
	}

	x001 := homonculus{
		name: "x001",
		creator: "Josh",
	}


	// polymorphism (many and change) - different methods have the same name,
	somebody.speak();
	x001.speak(); // Notice homonculus method speak() uses diffrent values

	fmt.Printf("somebody is of type: %T \n", somebody)


	whereAreYou(somebody)
	whereAreYou(x001)
}