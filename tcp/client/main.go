package main

import (
	"fmt"
	"log"
	"net"
)

func main() {

	conn, err := net.Dial("tcp", "localhost:8000")

	if err != nil {
		log.Panic(err)
	}
	defer func() {
		err := conn.Close()
		if err != nil {
			log.Panic(err)
		}
	}()
	fmt.Fprintf(conn, "I called You")
}
