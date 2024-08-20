package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func Handle(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
	}

	defer func() {
		err := conn.Close()
		if err != nil {
			log.Panic(err)
		}
	}()

	fmt.Println("End")
}
