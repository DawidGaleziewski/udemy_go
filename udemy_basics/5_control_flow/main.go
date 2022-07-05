package main

import ("fmt")

func main(){
	loopsInGo()
	consitionsInGo()
	switchesInGo()
}

func loopsInGo(){
	// init; condition; post loop
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	y := 0
	// when we look in language specification and for statment EBNF we see that all the forms are optional (ForStmt = "for" [ Condition | ForClause | RangeClause ] Block .). Therefore we can even do a wild loop like this:
	for {
		if(y > 5){
			break // nice way of using wild loops is to put them in go rutines and listen for a event. We can then break them anytime
		}

		fmt.Printf("Listener loop: %v \n", y)
		y++
	}

	// In simplest form we can use loops with only condition
	x := 0
	for x < 5 {
		fmt.Println(x)
		x++
	}

	// continue allows us to jump to next iteration
	for i := 0; i < 5; i++ {

		if(i%2 == 0){
			continue
		}

		fmt.Printf("only odd numbers %v \n", i)
	}
}

func consitionsInGo(){
	// if else works as in any other programming lang
	if true {
		fmt.Println("Hello there!")
	}

	// in golang we can put two statments in the same lane if we use semicolons
	fmt.Println("first stament"); fmt.Println("second statment")

	// just like with loops we can use a initialisatio statment that will be block scoped. it will be initialised before evaluating statment executes
	if x := 42; x == 42 {
		fmt.Printf("x value is: %v",x)
	} else if (x == 40){
		fmt.Printf("x value is: %v",x)
	} else {
		fmt.Printf("x value is: %v",x)
	}
}

func switchesInGo(){
	fmt.Println("")
	x := "mark"

	// we can omit switch expression. It will evaluate by default true and search for cases that are also true!
	switch {
		case ("dave" == x):
			fmt.Println("Hello Dave!")
		case ("mark" == x):
			fmt.Println("Prints as its true and continues to evaluate next cases!")
			fallthrough // keyword that can be used to do multiple things in one switch. Not really recomanded to be used as it may create funky logic
		case ("mark" == x):
			fmt.Println("Prints as its true")
			fallthrough
		case false:
			fmt.Println("Prints despite beeing false and goes to evaluate next case!")
			fallthrough
		case false:
			fmt.Println("Does not print as its not true")
		default:
			fmt.Println("default case, fired only if nothing else was fired")

	}

	// we can also evaluate all cases agains the value we provide to switch
	switch "dave" {
	case "mark":
		fmt.Println("# hi mark")
	case "dave":
		fmt.Println("# hi dave")
	// complier wont allow us to use bool when we provided string!
	// case true: 
	// 	fmt.Println("# hi true")
	}

}