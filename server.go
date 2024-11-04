package main

import (
	"fmt"
	"net"
)

const (
	serverAddress = "localhost:8000"
)

func startServer() {
	listener, err := net.Listen("tcp", serverAddress)
	if err != nil {
		fmt.Println("Error listening on ", serverAddress)
	}

	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	// read the packet we get

	// send ack packet

}
