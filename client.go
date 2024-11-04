package main

import (
	"fmt"
	"net"
	"time"
)

var buf []byte
var stack []TimeStack

const (
	PING_RATE = 2
)

func connectToServer() {
	conn, err := net.Dial("tcp", serverAddress)
	if err != nil {
		fmt.Println("Client: Error dialing server")
	}

	defer conn.Close()

	buf = make([]byte, 1024)

	// Pull timestamp
	// Is the timestamp of the sending included in the packet or is it locally checked?
	// Packet does not include local time, we just hold on to that as we send em

	// listen and send at the same time

	go sendPackets(conn)
	go receiveACK(conn)
}

// Add hanging packets to a slice for a Stack
// Do I need a stack for this?
func sendPackets(conn net.Conn) {
	stack := TimeStack{}

	for {
		currentTime := time.Now()
		conn.Write([]byte(currentTime.String()))

		stack.Push(currentTime)

		// Wait until duration elapses and send another
		targetTime := time.Now().Add(time.Duration(PING_RATE) * time.Second) // Wait until 10 seconds from now
		<-time.After(time.Until(targetTime))

	}

}

// Have a buffer for incoming packets
func receiveACK(conn net.Conn) {
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Client: Error reading from server:", err)
		} else {
			// Received the packet we were worried about
			fmt.Println(string(buf[:n]))
		}
	}
}

// Need to reason about how to handle packets delivered out of order
