package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func handleError(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func readConn(c net.Conn) {
	r := make([]byte, 1024)
	msg, err := c.Read(r)
	handleError(err)
	println(string(r[:msg]))
}

func main() {
	var uname string
	conn, err := net.Dial("tcp", "localhost:5050")
	handleError(err)
	println("Connection established.")
	print("Enter your name: ")
	fmt.Scanln(&uname)
	// Get a loop going
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		switch {
		case scanner.Text() == "quit":
			_, err := conn.Write([]byte(fmt.Sprintf("%v left.", uname)))
			handleError(err)
			readConn(conn)
			defer conn.Close()
			os.Exit(1)
		default:
			message := []byte(scanner.Text())
			msgToSend := fmt.Sprintf("%v: %v", uname, string(message))
			_, err := conn.Write([]byte(msgToSend))
			readConn(conn)
			handleError(err)
		}
	}
}
