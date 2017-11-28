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
			break
		}
	}
	io.WriteString(c, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(c, "Content-Length: %d\r\n", 10)
	fmt.Fprint(c, "Content-Type: text/plain\r\n")
	io.WriteString(c, "\r\n")
}
