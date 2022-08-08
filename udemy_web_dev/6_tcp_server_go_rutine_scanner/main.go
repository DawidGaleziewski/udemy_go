package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func main(){
	li, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Panicln(err)
	}

	// defer closing of the listener
	defer li.Close()

	// loop forever
	for {
		connection, err := li.Accept() // we accept each connection
		if err != nil {
			log.Panicln(err)
		}

		go handle(connection) // each connection will be launched and handlet in its own go ruoutine
	}
}

func handle(conn net.Conn) {
	// we can set the deadline on the connection after which it will be closed
	err := conn.SetDeadline(time.Now().Add(10 * time.Second))
	if err != nil {
		log.Println("conncection timed out")
	}

	// responding to a opened connection
	io.WriteString(conn, "Responding from tcp server!")

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {// advances to next token. In our case this is line of text. Thos will print all the headers/data regarding the request that is incoming. This will be a eternall loop as Scan will just go on forewer listening for next line
		ln := scanner.Text()
		// Using telnet on windows: $ winpty telnet localhost 8080 we can connect
		// each thing we writte during the telnet connection will be printed here
		fmt.Println(ln)
	}
	defer conn.Close()
	fmt.Println("Code got here")
}