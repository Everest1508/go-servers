package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func Handle(conn net.Conn) {
	data := make(map[string]string)

	// err := conn.SetDeadline(time.Now().Add(10 * time.Second))
	// if err != nil {
	// 	log.Println("Error setting connection deadline:", err)
	// 	return
	// }

	_, err := fmt.Fprintf(conn, "IN-MEMORY DB\nCOMMANDS\nSET key value\nGET key\nDEL key\n")
	if err != nil {
		log.Println("Error sending initial message:", err)
		return
	}

	defer func() {
		if err := conn.Close(); err != nil {
			log.Println("Error closing connection:", err)
		}
	}()

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		log.Println(ln)
		cmd := strings.Fields(ln)
		log.Println(cmd)

		if len(cmd) == 0 {
			continue
		}

		switch cmd[0] {
		case "GET":
			if len(cmd) != 2 {
				fmt.Fprintf(conn, "Usage: GET key\n")
				continue
			}
			v, exists := data[cmd[1]]
			if exists {
				fmt.Fprintf(conn, "%s\n", v)
			} else {
				fmt.Fprintf(conn, "Key not found\n")
			}

		case "SET":
			if len(cmd) != 3 {
				fmt.Fprintf(conn, "Usage: SET key value\n")
				continue
			}
			k := cmd[1]
			v := cmd[2]
			data[k] = v

		case "DEL":
			if len(cmd) != 2 {
				fmt.Fprintf(conn, "Usage: DEL key\n")
				continue
			}
			delete(data, cmd[1])

		default:
			fmt.Fprintf(conn, "Unknown command: %s\n", cmd[0])
		}
	}

	if err := scanner.Err(); err != nil {
		log.Println("Error reading from connection:", err)
	}

	fmt.Println("End")
}
