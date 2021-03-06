package main

import (
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
	println("Server running.")
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
			//TODO: put this loop in its own routine
			for _, k := range room.connections {
				_, err := k.Write(msg[:read])
				if err != nil {
					log.Fatal(err)
				}
			}
			defer c.Close()
		}(conn)
	}
}
