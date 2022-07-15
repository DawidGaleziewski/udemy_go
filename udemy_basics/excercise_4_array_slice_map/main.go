package main

import "fmt"

func main() {
	practiceSlices()
	prcaticeMap()
}

func practiceSlices() {
	numbersArray := [5]int{1, 2, 3, 4, 5}

	for _, item := range numbersArray {
		fmt.Println(item)
	}

	fmt.Printf("array type is %T and its value %v \n", numbersArray, numbersArray)

	firstHalf := numbersArray[:2]
	secondHalf := numbersArray[2:]
	fmt.Println(firstHalf, secondHalf)

	numbersSlice := []int{55, 66, 1, 22, 42, 52, 64}

	numbersSlice = append(numbersSlice[0:2], numbersSlice[5:]...)
	fmt.Println(numbersSlice)

	// creating underlying array only once. We create a slice with certain length and capacity
	//!IMPORTANT: we have to set length to 0. Otherwise append will append new values AFTER index 5, overflowing the cap and forcing creation of 10/10 array
	states := make([]string, 0, 5)
	states = append(states, "Nevada", "Maine", "Iowa", "Hawaii", "Georgia") // we have to use append. If we re-assing array here, slice we created with make goes thrown away

	fmt.Println("len and cap:", len(states), cap(states))

	// If we wanted for length and cap to stay the same. We would have to iterate on created array and assing each value to a specific index

	multiDimensionalArray := [][2]int{{1, 2}, {52, 32}}
	fmt.Println(multiDimensionalArray)
}

func prcaticeMap() {
	favThings := map[string][]string{"bond": {"Martini", "Football"}, "germoney": {"money"}}

	favThings["germoney"] = []string{"test"}
	delete(favThings, "bond")

	favThings["test1"] = []string{"test", "test"}

	for i, item := range favThings {
		fmt.Println("map item index and value:", i, item)
	}
}
