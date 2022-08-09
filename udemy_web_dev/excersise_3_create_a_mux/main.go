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
			router(conn)
		}(conn)
	}
}

func router(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	method := ""
	url := ""
	i := 0;

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

	req := HTTP{
		conn,
		method,
		url,
	}


	route(req, "GET", "/", func(res HTTP){
		page := getTemplate("index.gohtml", "test")
		res.send(page)
		conn.Close()
	})

	route(req, "GET", "/about", func(res HTTP){
		page := getTemplate("about.gohtml", "test")
		res.send(page)
		conn.Close()
	})

	
	route(req, "GET", "/contact", func(res HTTP){
		page := getTemplate("contact.gohtml", "test")
		res.send(page)
		conn.Close()
	})

	route(req, "POST", "/contact", func(res HTTP) {
		page := getTemplate("thanks.gohtml", "test")
		res.send(page)
		conn.Close()
	})


	route(req, "*", "*", func(res HTTP) {
		page := getTemplate("404.gohtml", "test")
		res.send(page)
		conn.Close()
	})
}

func route(req HTTP, method string, url string, callback func(HTTP)){	
	if req.method == method && req.url == url {
		callback(req)
	}
}

func getTemplate(name string, variable string) string{
	buffer := new(bytes.Buffer)
	err := tpl.ExecuteTemplate(buffer, name, variable)
	if err != nil {
		log.Panicln(err)
	}

	return buffer.String()
}

type HTTP struct {
	conn net.Conn;
	method string;
	url string;
}

func (H HTTP) startLine(status string) string {
	return fmt.Sprintf("HTTP/1.1 %v %v\r\n", status, HTTPStatusDesc[status])
}
func (H HTTP) contentLength(body string) string {
	return fmt.Sprintf("Content-Length: %d\r\n", len(body))
}
func (H HTTP) contentType() string {
	return fmt.Sprintf("Content-Type: text/html\r\n")
}
func (H HTTP) lineBreak() string {
	return fmt.Sprintf("\r\n")
}

func (H HTTP) send(body string) {
	fmt.Println(H)
	fmt.Fprint(H.conn, H.startLine("200"))
	fmt.Fprint(H.conn, H.contentLength(body))
	fmt.Fprint(H.conn, H.contentType())
	fmt.Fprint(H.conn, H.lineBreak())
	fmt.Fprint(H.conn, body)
}

var HTTPStatusDesc = map[string]string{
	"200": "OK",
}