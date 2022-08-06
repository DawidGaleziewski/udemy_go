package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

// we create first a TCP server
// TCP is a protocol on top of which HTTP protocol is build
// protocols are rules of communication
// if a request commin to a server is formatted according to a HTTP protocol standard, we can parse it and use it


func main(){
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panicln(err)
	}

	// defer closing of the listener
	defer li.Close()

	// loop forever
	for {
		connection, err := li.Accept()
		if err != nil {
			log.Panicln(err)
		}

		// as conncection is both reader and a writter we can pass it as a argument to those methods.
		// we can test this bu using telnet
		// Using telnet on windows: $ winpty telnet localhost 8080
		io.WriteString(connection, "\nHello from TCP server \n")
		fmt.Fprintln(connection, "Doing somerhing")
		fmt.Fprintf(connection, "%v", "Well i Hope")

		connection.Close()
	}
}