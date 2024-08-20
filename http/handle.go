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
	info := request(conn)

	// deliver response to client (write)
	response(conn, info)
}

func request(conn net.Conn) []string {
	var method string
	var url string
	i := 0
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {

		ln := scanner.Text()
		fmt.Println(ln)

		if i == 0 {
			fmt.Printf("***Method %s\n", strings.Fields(ln)[0])
			fmt.Printf("***URL %s\n", strings.Fields(ln)[1])
			method = strings.Fields(ln)[0]
			url = strings.Fields(ln)[1]
		}
		if ln == "" {
			break
		}
		i++
	}

	return []string{
		method, url,
	}
}

func response(conn net.Conn, info []string) {
	body := `<!DOCTYPE html><html>
	<head>
	<title>Server Response</title>
	</head>
	<body>
	<h1>
	method : %s
	</h1>
	<h1>
	path : %s
	</h1>	
	</body>
	</html>`

	fmt.Fprint(conn, "HTTP/1.1 200 OK \r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html \r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprintf(conn, body, info[0], info[1])
}
