package main

import (
	"log"
	"net"
)

func main() {
	li, err := net.Listen("tcp", "0.0.0.0:8000")
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

		go Handle(conn)
	}
}
