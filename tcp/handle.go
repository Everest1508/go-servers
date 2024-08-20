package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

func Handle(conn net.Conn) {
	err := conn.SetDeadline(time.Now().Add(10 * time.Second))
	if err != nil {
		log.Println("~~~connection timeout~~~")
	}
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		fmt.Fprintf(conn, ">>%s\n", ln)
	}

	defer func() {
		err := conn.Close()
		if err != nil {
			log.Panic(err)
		}
	}()

	fmt.Println("End")
}
