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

type Contractor struct{}

func (Contractor) Fasten(d NailDriver, nailSupply *int, b *Board){
	for b.NailsDriven < b.NailsNeeded {
		d.DriveNail(nailSupply, b)
	}
}

func main(){
	fmt.Println("test")
}