package main

import (
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:5050")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	message := []byte("Bruuuuuuuuh")
	_, err = conn.Write(message)
	if err != nil {
		log.Fatal(err)
	}
	rb := make([]byte, 1024)
	resp, err := conn.Read(rb)
	if err != nil {
		log.Fatal(err)
	}
	println("Message received:", string(rb[:resp]))
}
