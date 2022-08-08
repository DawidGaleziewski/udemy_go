package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:9090")
	if err != nil {
		log.Panicln(err)
	}
	defer conn.Close()

	// we could also use fmt.Println to write into connection
	num, err := conn.Write([]byte("howdy there from client!"))
	if err != nil {
		log.Println(err)
	}

	log.Println("response from server write action is: ", num)

	bs, err := ioutil.ReadAll(conn)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(string(bs))

	conn.Close()
}