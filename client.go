package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	var msg string
	var uname string
	print("Enter your name: ")
	fmt.Scanln(&uname)
	print("Enter your message: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	msg = scanner.Text()
	conn, err := net.Dial("tcp", "localhost:5050")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	message := []byte(msg)
	_, err = conn.Write(message)
	if err != nil {
		log.Fatal(err)
	}
	rb := make([]byte, 1024)
	resp, err := conn.Read(rb)
	if err != nil {
		log.Fatal(err)
	}
	println(uname, ":", string(rb[:resp]))
}
