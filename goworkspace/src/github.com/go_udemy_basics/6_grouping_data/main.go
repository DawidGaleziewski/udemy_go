package main

import ("fmt")

func main(){
	workingWithArray()
	workingWithSlice()
	usingRangeForSliceIterations()
	slickingASliceWithColonOperator()
	appendToSlice()
	optimizingSlicesWithMake()
	multiDemensionalSlices()
	usingMaps()
}

func workingWithArray(){
	// for arrays, we have to specify the size. Which is done inside the brackets. Length is part of arrays type!
	// arrays are primary building blocks for slices and we do not use them all the time
	var x [5]int
	fmt.Println(x)

	x[3] = 42 // override at index
	fmt.Println(x)
}

func workingWithSlice(){
	// using composite literal to declare the type of a variable AND assign a value to it
	// composite literals are expressions that create a new instance each time they are evaluated
	//x := type{values}
	x := []int{4,5,6} //we specify a type and at the same time we specify THE VALUES in the curly braces that will be assigned to the variable
	// SLICE allows us to group together VALUES of the same type

	// can be used with arrays
	y := [3]int{3,2,1}

	// can also use diffrent types of values. We can use this syntax as long as those are "composite types". We cannot use it with primitives
	z := [3]string{"Mike", "Siri", "Trevor"}

	fmt.Println(x, y, z)
}

func usingRangeForSliceIterations(){
	names := []string{"Dave", "Merry", "Susan"}

	// keyword range can be used in for loop on arrays/slices to iterate on elements of the array and their indexes
	for i, v:= range names {
		fmt.Printf("# i is: %v. v is %v \n", i, v)
	}
}

func slickingASliceWithColonOperator(){
	ingridients := []string{"tomato", "potato", "sugar", "salt", "coal", "blood", "x factor"}

	// colon allows us to slice part of a array in index range
	middle := ingridients[2:3]

	// similar to python slice from begging until index or from index until end if we wont specify it
	eatable := ingridients[:3]
	nonEatable := ingridients[4:]

	//joined := middle + eatable + nonEatable

	fmt.Println(middle, eatable, nonEatable)
}

// We use append, which is special build in function in go
func appendToSlice(){
	var partyPeople []string
	fmt.Printf("party people at 8:00: %v \n", partyPeople)
	partyPeople = append(partyPeople, "Dave", "Marie")
	fmt.Printf("party people at 16:00: %v \n", partyPeople)

	// we can also append slice to slice, go has similar operator that js has ...
	drunkPeople := []string{"Sam", "Carrol"}

	partyPeople = append(partyPeople, drunkPeople...)
	fmt.Printf("party people at 18:00: %v \n", partyPeople) // returns another slice of the same type

	// we can also use append to delete values. We create new slice that includes values up to index 1 and from idex 2
	partyPeople = append(partyPeople[:1], partyPeople[2:]...)
	fmt.Printf("party people at 20:00: %v \n", partyPeople) 
}
func optimizingSlicesWithMake(){
	// type, length, capacity
	x := make([]int, 10, 12)
	fmt.Printf("x is %v, its length is: %v, and capacity is: %v \n", x, len(x), cap(x))

	// x[70] = 15-- we cant do this as we go over the length of the slice
	// we can however append it. Then its length will grow. If we go over cap it will work but it will double the cap of the array. This will take some processing power however as it will have to throw away the old array, create new one and move all the values
	x = append(x, 100)
	fmt.Printf("x is %v, its length is: %v, and capacity is: %v \n", x, len(x), cap(x))
}

func multiDemensionalSlices(){
	geojson := [][2]int{{25, 26}, {15, 90}, {20, 70}}
	fmt.Printf("geojson data: %v \n", geojson)
}

// same as objects in  javascript. Unordered lists with quick lookup
func usingMaps(){

	// inside brackets we put the key type (here string), after that we declare the values type. Rest is declaring the values inside {}
	personelAges := map[string]int{
		"James": 32,
		"Merry": 18,
		"Json" : 200,
	}
	fmt.Printf("map person is %v \n", personelAges)
	fmt.Printf("map person by key is %v \n", personelAges["James"])
	fmt.Printf("map person by key is %v \n", personelAges["Duffy"]) // important thing is that if we do not find value by key, zeroed value will be returned

	valueNotExisting, ok1 := personelAges["Duffy"] // there is a way however to check if the value exists. Into ok it will return a check
	valueExisting, ok2 := personelAges["James"] 

	fmt.Printf("non existing value and not ok: %v , %v \n existing value and ok: %v %v \n", valueNotExisting, ok1, valueExisting, ok2)

	// often used idiomatic construction to check if value exists. First we initialise value and ok, after that if consition is ok we do something
	if v, ok := personelAges["does not exist"]; ok {
		fmt.Printf("Do something with %v", v)
	} else {
		fmt.Println("Value does not exist")
	}

	// We can add values to map simply by using equals
	personelAges["Lucy"] = 99

	fmt.Printf("newly added value: %v \n", personelAges)

	// We cen iterate over map by using range keyword

	const MAX_AGE = 30;
	for key, value := range personelAges {

		if(value > MAX_AGE){
			fmt.Printf("%v is already %v years old. Lets fire him and hire new one \n", key, value)
		} else {
			fmt.Printf("%v is only %v years old. Lets still keep him for the next %v years \n", key, value, MAX_AGE - value)
		}
	}

	// we use delete build in method for deleting keys from map
	delete(personelAges, "Lucy")
	fmt.Printf("after deleting Lucy %v", personelAges)
}