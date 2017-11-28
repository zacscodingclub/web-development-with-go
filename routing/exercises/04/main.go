package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	connType := "tcp"
	port := ":3000"
	s, err := net.Listen(connType, port)
	if err != nil {
		log.Fatalln(err)
	}
	defer s.Close()

	for {
		conn, err := s.Accept()
		if err != nil {
			log.Println(err)
		}
		defer conn.Close()

		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			ln := scanner.Text()
			fmt.Println(ln)
		}

		fmt.Println("Code got here")
		io.WriteString(conn, "I see you connected.")
	}
}
