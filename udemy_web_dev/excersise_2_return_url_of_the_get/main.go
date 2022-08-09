package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":9090");
	if err != nil {
		log.Panicln(err)
	}

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Panicln(err)
		}

		go func(conn net.Conn){
			index := 0
			scanner := bufio.NewScanner(conn)
			for scanner.Scan() {
				if index == 0 {
					ln := scanner.Text()
					fields := strings.Fields(ln)
					method := fields[0]
					url := fields[1]

					if method == "GET"{
						fmt.Println(method, url)
					}
				}
				index++
			}
		}(conn)
	}
}

