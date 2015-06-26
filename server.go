package main

import (
	// "io"
	"log"
	"net"
)

var current_msg string

type ChatRoom struct {
	connections []net.Conn
}

func main() {
	l, err := net.Listen("tcp", ":5050")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()
	room := ChatRoom{}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		room.connections = append(room.connections, conn)
		println("We got a connection, bruh!!")
		go func(c net.Conn) {
			msg := make([]byte, 1024)
			read, err := c.Read(msg)
			if err != nil {
				log.Fatal(err)
			}
			for _, k := range room.connections {
				k.Write(msg[:read])
			}
			// io.Copy(c, c)
			// c.Close()
		}(conn)
	}
}
