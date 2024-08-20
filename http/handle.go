package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func Handle(conn net.Conn) {
	defer conn.Close()

	// handle request from client ( read )
	method := request(conn)
	if method == "GET" {
		fmt.Print(method)
	}
	// deliver response to client (write)
	response(conn)
}

func request(conn net.Conn) string {
	method := "GET"
	i := 0
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {

		ln := scanner.Text()
		fmt.Println(ln)

		if i == 0 {
			fmt.Printf("***Method %s", strings.Fields(ln)[0])
			method = strings.Fields(ln)[0]
			fmt.Printf("***URL %s", strings.Fields(ln)[1])
		}
		if ln == "" {
			break
		}
		i++
	}

	return method
}

func response(conn net.Conn) {
	body := `<!DOCTYPE html><html>
	<head>
	<title>Server Response</title>
	</head>
	<body>
	<h1>
	Response
	</h1>	
	</body>
	</html>`

	fmt.Fprint(conn, "HTTP/1.1 200 OK \r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html \r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}
