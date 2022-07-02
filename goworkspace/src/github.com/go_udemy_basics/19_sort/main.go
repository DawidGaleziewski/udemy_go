package main

import (
	"fmt"
	"sort"
)

func main() {
 	defaultSortingLib()
	attachingSortToStructs()
}

func defaultSortingLib(){
	nums := []int{5, 3, 4, 6, 10, 200, 0}
	sort.Ints(nums) // will mutate the variable and sort it

	fmt.Printf("sorted ints asre %v \n", nums)
}



// create a custome type
type person struct{
	name string
	age int
}

// implement new type ByAge that will implement sort.Interface for []person based on the age field
type ByAge []person

// Use reciver to attach a method to this new type

// accordint to documentation in order to use sorts package type Interface, we need to implement 3 methods to that type. After that the SORT interface will be implemented implicitly
func (a ByAge) Len() int {return len(a)}
func (a ByAge) Less(i, j int) bool {return a[i].age < a[j].age}
func (a ByAge) Swap(i, j int){a[i], a[j] = a[j], a[i]}


// implementation to sort by name
type ByName []person
func (a ByName) Len() int {return len(a)}
func (a ByName) Less(i, j int) bool {return a[i].name < a[j].name}
func (a ByName) Swap(i, j int){a[i], a[j] = a[j], a[i]}

func attachingSortToStructs(){


	// declare a variable using that type
	people := []person{
		{"Bob", 31},
		{"John", 14},
		{"Agnes", 9},
	}

	fmt.Printf("people are: %v \n", people)

	sort.Sort(ByAge(people))
	fmt.Printf("people are after sort: %v \n", people)

	sort.Sort(ByName(people))
	fmt.Printf("people are after sort by name: %v \n", people)
}

// Recap how to sort custome values/ types/ datastructures
// 1. Declare a new type, with the underlying slice of type we want to sort: type ByName []person

// 2. declare 3 methods accepting newly created type as a reciver
// a) func (a ByName) Len() int {return len(a)}
// b) func (a ByName) Less(i, j int) bool {return a[i].name < a[j].name}
// c) func (a ByName) Swap(i, j int){a[i], a[j] = a[j], a[i]}
// In this moment, type ByName will become magically, by implicit type Interface (which is also a interface) a type ... Interface.

// 3 conver type []people into type ByName
// 4 Pass the value to sort: sort.Sort(ByName(people))

// people.sort((a,b) => a - b)
