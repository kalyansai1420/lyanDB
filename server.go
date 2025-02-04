package main

import (
	"fmt"
	"net"
)

func startServer(port string) {
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}

	defer listener.Close()

	fmt.Println("lyanDB server is running on port", port)

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
	fmt.Println("New client connected:", conn.RemoteAddr())

	buf := make([]byte, 1024)

	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Client disconnected:", err)
			return
		}

		message := string(buf[:n])
		parsedMessage, err := deserializeRESP(message)
		if err != nil {
			conn.Write([]byte(serializeRESP("Error parsing command")))
			continue
		}

		response, err := ExecuteCommand(parsedMessage)
		if err != nil {
			conn.Write([]byte(serializeRESP(err.Error())))
		} else if response == "" {
			conn.Write([]byte(serializeRESP("nil")))
		} else {
			conn.Write([]byte(serializeRESP(response)))
		}

	}
}
