package main

import ("fmt")
type animal struct {
	name string;
	age int;
	weight int;
}

type wolf struct {
	animal
	furColor string
	partOfPack bool
}

// func (r reciver) identifier(parametes) (returns) {....}
// the way we attach a method to a struct is via reciver. We provide the struct same as we would with argument but we put it in front of the name of the function
func (s wolf) howl(){
	fmt.Printf("AWOOO! I am %v \n", s.name)
}

func (s wolf) eat(name string){
	fmt.Printf("*Nom Nom* %v eats %v \n", s.name, name)
}


func main(){
	henryTheWOlf := wolf{
		animal: animal{
			name: "henry",
			age: 2,
			weight: 23,
		},
		furColor: "Blue",
		partOfPack: false,

	}

	henryTheWOlf.howl()
	henryTheWOlf.eat("Murry The rabbit")
}