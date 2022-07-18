package main

import "fmt"

// whole read on https://www.ardanlabs.com/blog/2015/09/composition-with-go.html

type Board struct {
	NailsNeeded int;
	NailsDriven int;
}

/**
"These two interfaces, NailDriver and NailPuller, each implement a single, well defined behavior. This is what we want. Being able to break down behavior into individual, simple acts, lends itself to composability, flexibility and readability."
 **/

// Represents the behaviour to drive nails into the board
type NailDriver interface { // often interface will end with -er
	DriveNail(nailSupply *int, b *Board)
}

type NailPuller interface {
	PullNail(nailSupply *int, b *Board)
}


/**
This interface is composed from both the NailDriver and NailPuller interfaces. This is a very common pattern you will find in Go, taking existing interfaces and grouping them into composed behaviors.
**/
type NailDrivePuller interface {
	NailDriver
	NailPuller
}

// Mallet is a tool that punds in nails
type Mallet struct {} 

// DriveNail punds a nail into the specified board
func (Mallet) DriveNail(nailSupply *int, b *Board){
	*nailSupply--

	b.NailsDriven++

	fmt.Printf("Mallet punds nail int board, nails left %v \n", nailSupply)
}

type Crowbar struct {}

func (Crowbar) PullNail(nailSupply *int, b *Board){
	b.NailsDriven--

	*nailSupply++
	fmt.Println("Crowbar: yanked nail out of the board.")
}

// Contractor carries out the task
type Contractor struct{}

// Fasten drives the nails into the board
//  The method requires the user to pass as the first parameter a value that implements the NailDriver interface. This value represents the tool the contractor will use to execute this behavior. Using an interface type for the this parameter allows the user of the API to later create and use different tools without the need for the API to change.
func (Contractor) Fasten(d NailDriver, nailSupply *int, b *Board){
	for b.NailsDriven < b.NailsNeeded {
		d.DriveNail(nailSupply, b)
	}
}

func (Contractor) Unfasten(p NailPuller, nailSupply *int, b *Board){
	for b.NailsDriven > b.NailsNeeded {
		p.PullNail(nailSupply, b)
	}
}
//The final behavior for a contractor is a method called ProcessBoards which allows the contractor to work on a set of boards at one time:
func (Contractor Contractor) ProcessBoards(dp NailDrivePuller, nailSupply *int, boards []Board){
	for i := range boards {
		board := &boards[i]; // we take address of the board to work at the pointer and mutate those values

		switch {
			case  board.NailsDriven < board.NailsNeeded:
					Contractor.Fasten(dp, nailSupply, board)
			case board.NailsDriven > board.NailsNeeded:
					Contractor.Unfasten(dp, nailSupply, board)		
		}
		
	}
}

// Notice that so far Fasten and Unfasten requires requires value that implements only ONE act of behaviour. But process boards requires a value that implements two acts of behaviour.

// Toolbox can contains any number of tools.
/** When embedding a type inside of another type, it is good to think of the new type as the outer type and the type being embedded as the inner type. This is important because then you can see the relationship that embedding a type creates. **/

type Toolbox struct {
	NailDriver // We have not embedded a struct type into our Toolbox but two interface types. This means any concrete type value that implements the NailDriver interface can be assigned as the inner type value for the NailDriver embedded interface type.
	NailPuller
	nails int
}

func main(){

	// We specify a slice of boards with information on what each of them needs
	boards := []Board{
		// To be removed
		{NailsDriven: 3},
		{NailsDriven: 1},
		{NailsDriven: 5},

		// to be fasten
		{NailsNeeded: 6},
		{NailsNeeded: 10},
		{NailsNeeded: 1},
	}

	tb := Toolbox{
		NailDriver: Mallet{},
		NailPuller: Crowbar{},
		nails: 10,
	}

	var c Contractor
	c.ProcessBoards(&tb )
	fmt.Println("test")
}