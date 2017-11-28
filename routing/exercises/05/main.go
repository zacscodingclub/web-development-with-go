package main

import (
	"bufio"
	"fmt"
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
		go serve(conn)
	}
}

func serve(c net.Conn) {
	defer c.Close()

	scanner := bufio.NewScanner(c)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if ln == "" {
			fmt.Println("THE END OF THE HEADER.")
		}
	}
}
