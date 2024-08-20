package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Panic(err)
	}

	defer func() {
		err = li.Close()
		if err != nil {
			log.Panic(err)
		}
	}()

	for {
		conn, err := li.Accept()

		if err != nil {
			log.Panic(err)
		}

		io.WriteString(conn, "\n TCP Server \n")
		fmt.Fprintln(conn, "Fprint Statment")
		fmt.Fprintf(conn, "%v", "Fprintf ")
	}
}
