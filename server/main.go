package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"server/commands"
	"server/config"
)

// server constants
const (
	serverHost = "localhost"
	serverPort = "6379"
	connType   = "tcp"
)

func main() {

	// Listen for incoming connections.
	l, err := net.Listen(connType, serverHost+":"+serverPort)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	fmt.Println("Server listening at :", serverHost, serverPort)
	// Close the listener when the application closes.
	defer l.Close()

	for {
		// Listen for an incoming connection.
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		// Handle connections in a new goroutine.
		go handleRequest(conn)
	}

}

// Handles incoming requests.
func handleRequest(conn net.Conn) {
	// Create a new reader
	fmt.Println("Client connected")
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		bytes := scanner.Bytes()
		inputString, err := config.ParseRequest(bytes)
		log.Println(inputString)
		if err == nil {
			result := commands.ExecuteCommand(inputString)
			conn.Write([]byte(result + "\n"))
		}
	}
}
