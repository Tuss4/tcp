package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	var uname string
	conn, err := net.Dial("tcp", "localhost:5050")
	if err != nil {
		log.Fatal(err)
	}
	println("Connection established.")
	print("Enter your name: ")
	fmt.Scanln(&uname)
	// Get a loop going
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		switch {
		case scanner.Text() == "quit":
			conn.Close()
			os.Exit(1)
		default:
			message := []byte(scanner.Text())
			_, err = conn.Write(message)
			if err != nil {
				log.Fatal(err)
			}
			rb := make([]byte, 1024)
			resp, err := conn.Read(rb)
			if err != nil {
				log.Fatal(err)
			}
			msg_strt := fmt.Sprintf("%v:", uname)
			println(msg_strt, string(rb[:resp]))
		}
	}
}
