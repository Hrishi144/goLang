package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	listner, err := net.Listen("tcp", ":9000")
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}
	defer listner.Close()

	fmt.Println("Server running on port 9000")

	for {
		conn, err := listner.Accept()
		if err != nil {
			continue
		}
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewScanner(conn)
	for reader.Scan() {
		msg := reader.Text()
		fmt.Println("Recieved:", msg)
	}
}
