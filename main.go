package main

import "fmt"

func main() {
	fmt.Println("Hello world!")

	startServer()
	// Wait until it starts -> add just like a 5 second delay

	connectToServer()
}
