package main

import (
	"bufio"
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Panicln(err)
	}

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Panicln(err)
		}

		go handle(conn)
	}
}

func handle(conn net.Conn){
	defer conn.Close()
	request(conn)
	response(conn)
}

func request(conn net.Conn){
	var i int
	log.Println("# Incoming request")

	// using scanner we have access to request here. This is not doing anything now just showcase on how to get to the request data
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)

		if i == 0 {
			// on first scan, first word will be the method, if this is a HTTP protocol request
			method := strings.Fields(ln)[0]
			fmt.Println("method: ", method)
		}

		if ln == "" {
			// according to HTTP protocol standard, if there is a empty line we are passed the headers, therefore we can break out of scannel as we do not need to wait for any tokens
			break
		}
		i++
	}
}

func response(conn net.Conn) {
	buffer := new(bytes.Buffer)
	tpl, err := template.ParseGlob("templates/*.gohtml")
	if err != nil {
		log.Println(err)
	}

	err = tpl.ExecuteTemplate(buffer, "index.gohtml", "")
	if err != nil {
		log.Println(err)
	}

	fmt.Println("sending ", buffer.String())

	// we need those headers to be send to the browser
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(buffer.String()))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, buffer.String())
}