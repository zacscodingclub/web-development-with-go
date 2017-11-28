package main

import (
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

		io.WriteString(conn, "I see you connected.")

		conn.Close()
	}
}
