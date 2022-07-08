// test should:
// be in same package as tested function
// end with _test.go
// be in a function with a signature
// names as func TestXxx where x is the name of the function we test
// running tests is done with go test
package main

import (
	"fmt"
	"testing"
)

// pointer to package pointers type T
func TestAdd(t *testing.T) { // we need pointer to type T from

	if add(2, 3) != 5 {
		t.Error("Wrong result of addition")
	}

	// runing table testes (a test suite)
	type test struct {
		data   []int
		answer int
	}

	tests := []test{
		test{[]int{2, 3}, 5},
		test{[]int{5, 5}, 10},
	}

	for _, test := range tests {
		result := add(test.data...)
		if result != test.answer {
			t.Error("Expected:", test.answer, ", got:", result)
		}
	}
}

// examples are great way of documenting (it will appear in docs as a example) and testing the code
// we do it by comment // Output:. This will be passed by test runner and checked agains the output.
// needs to be printed out, the printout will be compared to output
func ExampleAdd() {
	fmt.Println(add(2, 2, 2, 2))
	// Output:
	// 200
}
