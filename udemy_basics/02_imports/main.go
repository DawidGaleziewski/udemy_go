package main
// imported some random package we got from https://pkg.go.dev/rsc.io/quote/v4
import "fmt"
import "rsc.io/quote"

func main(){
	fmt.Println(quote.Hello())
}