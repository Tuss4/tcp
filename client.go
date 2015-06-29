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

func writeConn(c net.Conn, m string) {
	_, err := c.Write([]byte(m))
	handleError(err)
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
	// Loop:
	for scanner.Scan() {
		t := scanner.Text()
		println("the text:", t)
		switch {
		case t == "quit":
			writeConn(conn, fmt.Sprintf("%v left.", uname))
			readConn(conn)
			conn.Close()
			os.Exit(1)
		default:
			message := []byte(t)
			writeConn(conn, fmt.Sprintf("%v: %v", uname, string(message)))
			readConn(conn)
		}
	}
	println(scanner.Err())
}
