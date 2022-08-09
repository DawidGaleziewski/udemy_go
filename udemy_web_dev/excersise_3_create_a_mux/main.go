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

var tpl *template.Template

func init() {
	var err error
	tpl, err = tpl.ParseGlob("templates/*.gohtml")
	if err != nil {
		log.Panicln(err)
	}
}

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

		go func(conn net.Conn) {
			method, url := getHeaders(conn)
			router(conn, method, url)
		}(conn)
	}
}

func getHeaders(conn net.Conn) (string, string) {
	scanner := bufio.NewScanner(conn)
	method := ""
	url := ""

	i := 0
	for scanner.Scan() {
		ln := scanner.Text()
		if i == 0 {
			fmt.Println(ln)
			fields := strings.Fields(ln)
			method = fields[0]
			url = fields[1]

		}

		if ln == "" {
			break
		}
		i++
	}

	return method, url
}

func router(conn net.Conn, method string, url string) {


	switch method {
	case "GET":
		defer conn.Close()
		buffer := new(bytes.Buffer)
		templateName, exists := routesToTemplate[url]
		if exists {
			err := tpl.ExecuteTemplate(buffer, templateName, url)
			if err != nil {
				log.Println(err)
			}
		} else {
			tpl.ExecuteTemplate(buffer, "404.gohtml", url)
		}

		http := HTTP{
			url:    url,
			method: method,
			status: "200",
			body:   buffer.String(),
		}
		http.respond(conn)
	case "POST":
		defer conn.Close()
		buffer := new(bytes.Buffer)
		err := tpl.ExecuteTemplate(buffer, "thanks.gohtml", url)
		if err != nil {
			log.Println(err)
		}
		http := HTTP{
			url:    url,
			method: method,
			status: "200",
			body:   buffer.String(),
		}
		http.respond(conn)

	default:
		defer conn.Close()
		fmt.Fprint(conn, "unknown method")
	}
}

type HTTP struct {
	url    string
	method string
	status string
	body   string
}

func (H HTTP) startLine() string {
	return fmt.Sprintf("HTTP/1.1 %v %v\r\n", H.status, HTTPStatusDesc[H.status])
}
func (H HTTP) contentLength() string {
	return fmt.Sprintf("Content-Length: %d\r\n", len(H.body))
}
func (H HTTP) contentType() string {
	return fmt.Sprintf("Content-Type: text/html\r\n")
}
func (H HTTP) lineBreak() string {
	return fmt.Sprintf("\r\n")
}

func (H HTTP) respond(conn net.Conn) {
	fmt.Println(H)
	fmt.Fprint(conn, H.startLine())
	fmt.Fprint(conn, H.contentLength())
	fmt.Fprint(conn, H.contentType())
	fmt.Fprint(conn, H.lineBreak())
	fmt.Fprint(conn, H.body)
}

var HTTPStatusDesc = map[string]string{
	"200": "OK",
}

var routesToTemplate = map[string]string{
	"/": "index.gohtml",
	"/about": "about.gohtml",
	"/contact": "contact.gohtml",
}
