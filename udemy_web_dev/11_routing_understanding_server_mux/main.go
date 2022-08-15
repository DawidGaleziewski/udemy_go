package main

import (
	"io"
	"net/http"
)

type ecommerceApp struct{}

func (ecom ecommerceApp) ServeHTTP(resWriter http.ResponseWriter, req *http.Request){
	io.WriteString(resWriter, "see our new wares. Today all hats are on -20 sale")
}

type adminPanel struct{}
func (ap adminPanel) ServeHTTP(resWriter http.ResponseWriter, req *http.Request){
	io.WriteString(resWriter, "Login to admin: <form>")
}

func main() {
	basicCode()
	refactores()
	refactor2()
}

// using mux
func basicCode(){
	var ecomApp ecommerceApp
	var adminPanelApp adminPanel

	mux := http.NewServeMux()
	// mux server Handle method accepts the path and handler interface
	mux.Handle("/shop", ecomApp)
	// if we have /shop it will not catch and execute the code on sites with subpaths like /shop/test. but if we put path /shop/ it will.
	mux.Handle("/admin/", adminPanelApp)
	http.ListenAndServe(":9090", mux) // due to the fact that mux implements handler interface. We can pass it to ListenAndServe method
}

// using default mux
func refactores(){
	var ecomApp ecommerceApp
	var adminPanelApp adminPanel

	// passing values to default http server mux
	http.Handle("/shop", ecomApp)
	http.Handle("/admin", adminPanelApp)
	// we can use a default serve mux without the need to create it. If we pass nil, the mux that is build in http will be used. Thats why we can just do http.Handle()
	http.ListenAndServe(":9090", nil)
}

// using handle func (callbacks?)
func refactor2(){

	// we can also pass directly the functions to a handler. We therefore do not have to create our own types
	http.HandleFunc("/shop", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "see our new wares. Today all hats are on -20 sale")
	})

	http.HandleFunc("/admin", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "please log in")
	})

	http.ListenAndServe(":9090", nil)
}