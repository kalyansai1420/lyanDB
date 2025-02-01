package main

import (
	"fmt"
	"net"
	"strings"
	"github.com/kalyansai1420/lyandDB/resp" // Replace with the actual import path of the resp package
)

// handleConnection handles communication with each client
func handleConnection(conn net.Conn) {
	defer conn.Close()

	fmt.Println("New client connected:", conn.RemoteAddr())

	// Read the client's request
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading from client:", err)
		return
	}

	// Parse the received message and determine the command
	message := string(buf)
	parts := strings.Fields(message)

	// Check if the command is PING
	if len(parts) > 0 && strings.ToUpper(parts[0]) == "PING" {
		// Respond with PONG for PING command
		conn.Write([]byte(resp.SerializeRESP("PONG")))
		return
	}

	// Check if the command is ECHO
	if len(parts) > 0 && strings.ToUpper(parts[0]) == "ECHO" {
		// Join the rest of the message and respond with the same message
		// Using RESP format for bulk string
		response := strings.Join(parts[1:], " ")
		conn.Write([]byte(resp.SerializeRESP(response)))
		return
	}

	// Default response for unknown command
	conn.Write([]byte(resp.SerializeRESP("Unknown command")))
}

func main() {
	// Set up listener for incoming connections
	listener, err := net.Listen("tcp", ":6379")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer listener.Close()

	fmt.Println("lyanDB server is running on port 6379...")

	// Continuously accept and handle incoming connections
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		// Handle each connection concurrently
		go handleConnection(conn)
	}
}
