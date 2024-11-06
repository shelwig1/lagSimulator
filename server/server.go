package main

import (
	"fmt"
	"net"
	"time"
)

const (
	serverAddress = "localhost:8000"
)

func main() {
	startServer()
}
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
	defer conn.Close()
	fmt.Println("Client connected:", conn.RemoteAddr())

	for {
		buf := make([]byte, 1024)

		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Connection closed:", err)
			return
		}

		if n > 0 {
			raw := buf[:n]
			var newTime time.Time

			err = newTime.UnmarshalBinary(raw)
			if err != nil {
				fmt.Println("Error unmarshaling binary time")
			}

			//fmt.Println("Client: ", string(buf[:n]))
			fmt.Println("Client: ", newTime)

			_, err := conn.Write([]byte(raw))
			if err != nil {
				fmt.Println("Error sending packet: ", err)
			}

		}

	}
}
