package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {

	if len(os.Args) < 3 {
		fmt.Println("Please specify IP and Port of server")
		return
	}

	// connect to the socket
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%s", os.Args[1], os.Args[2]))
	if err != nil {
		fmt.Println("Cannot connect to server, exiting", err)
		return
	}
	for {
		// read in input from stdin
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("mockredis$> ")
		text, _ := reader.ReadString('\n')

		if strings.TrimSpace(text) == "QUIT" {
			fmt.Println("TCP client exiting...")
			return
		}
		// send to socket
		_, err = conn.Write([]byte(text))
		// listen for reply
		reply := make([]byte, 1024)
		_, err = conn.Read(reply)
		if err != nil {
			println("Server unavailable, exiting", err.Error())
			return
		}
		fmt.Print(string(reply))
	}
}
