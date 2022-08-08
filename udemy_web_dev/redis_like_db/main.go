package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

var data = make(map[string]string)

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


func handle(conn net.Conn) {
	//err := conn.SetDeadline(time.Now().Add(time.Second * 5))
	// if err != nil {
	// 	log.Println("CONNECTION TIME OUT")
	// }
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		fmt.Println("connection run")
		ln := scanner.Text()
		fs := strings.Fields(ln)
		//fmt.Fprintln(conn, "incoming request")


		if len(fs) == 0 {
			fmt.Fprint(conn, "NO ARGUMNETS")
			continue
		}
		switch fs[0] {
		case "GET":
				key := fs[1]
				value, isFound := data[key]
				if !isFound {
				  fmt.Fprintf(conn, "value %s not found \n", key)
				  continue
				}
				fmt.Fprintf(conn,"%s\n", value)
		case "SET":
				if len(fs) != 3 {
					fmt.Fprintln(conn, "EXPECTED VALUE!")
					continue
				}
				key := fs[1]
				value := fs[2]
				data[key] = value
				fmt.Fprintf(conn, "field %s set to %s\n", key, value)
			case "DEL":
				if len(fs) != 2 {
					fmt.Fprintln(conn, "PROVIDE THE NAME OF THE FIELD")
					continue
				}
				key := fs[1]
				delete(data, key)
				fmt.Fprintf(conn, "entry %s was deleted \n", key)
	
		default:
			fmt.Fprintln(conn, "INVALID COMMAND")
			continue

		}
	}
	defer conn.Close()
	fmt.Fprintln(conn, "CODE END")
}