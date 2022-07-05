package main

import ("fmt")

type person struct {
	first string
	last string
	age int
}



func main(){
	basicStrucs()
	embededStrucs()
	anonymusStruct()
}

// Creating composite data types with structs
func basicStrucs(){
	// Struct is a data structure that allows us to compose together values of different types. AKA aggregate data type

	// there are no classes or objects in go. But this could be compared to initialising a instance of a class in javascript
	p1 := person{
		first: "James",
		last: "Bond",
		age: 20,
	}

	p2 := person{
		first: "Merry",
		last: "James",
		age: 50,
	}

	fmt.Printf("Person 1 %v, person 2L %v, it is of type: %T \n", p1, p2, p1)
}

type warlock struct {
	person
	knownSpells []string
}

// aka composition
func embededStrucs(){
	// This works similar to inheritance! However as go has no classes this is called "composition". Values of another struct will be deconstructed on the same level into this struct
	johnnyMagic := warlock{
		person: person{ // we say that INNER type gets promoted to the OUTERTYPE
			first: "Johnny",
			last: "Magic",
		},
		knownSpells: []string{"Fireball"},
	}

	fmt.Printf("beware of the warlock %v \n", johnnyMagic.first)
}

func anonymusStruct(){
	// we can declare the struct and use it in variable declaration
	p1 := struct {
		first string
		last string
		age int
	}{
		first: "Slim",
		last: "Shady",
		age: 32,
	}
	


	fmt.Printf("my name is %v %v \n", p1.first, p1.last)

	
	p2 := struct {
		first string
		last string
		age int
	}{
		"James", // NOTE you can create a struct without the labels. But this is not a good practice as we will have issues reasing this
		"Bond",
		32,
	}

	fmt.Printf("my name is %v %v \n", p2.last, p2.first)
}