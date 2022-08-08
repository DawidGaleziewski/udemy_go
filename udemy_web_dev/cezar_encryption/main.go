package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func main() {
	li, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Panicln(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Panicln(err)
		}
		go handle(conn)
	}

}

func handle(conn net.Conn){
	conn.SetDeadline(time.Now().Add(time.Second * 5))
	defer conn.Close()
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := strings.ToLower(scanner.Text())
		encrypt := rot13(ln)
		fmt.Fprintf(conn, "%s - %s", ln, encrypt)
	}
}

// cezars encryption
func rot13(s string) []byte{
	bs := []byte(s)
	var r13 = make([]byte, len(bs))
	// ascii 97 - 122. as each ascii numbner in this range is lower case letter, we simply rotate the char on a table by 13 forward if its in range or backowrds in ortger cases
	for i, v := range bs {
		if v <= 109 {
			r13[i] = v + 13
		} else {
			r13[i] = v - 13
		}
	}

	return r13
}