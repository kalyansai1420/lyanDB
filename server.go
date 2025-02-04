package main

import (
	"bufio"
	"fmt"
	"net"

)

func StartServer(port string) {
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

		go HandleConnection(conn)
	}
}

func HandleConnection(conn net.Conn) {
	defer conn.Close()
	fmt.Println("New client connected:", conn.RemoteAddr())

	reader := bufio.NewReader(conn)

	for {
		parsedMessage, err := DeserializeRESP(reader)
		if err != nil {
			conn.Write([]byte("-ERR Error parsing RESP command\r\n"))
			continue
		}
		fmt.Println("parsed: ", parsedMessage)
		
		if len(parsedMessage) == 0 {
			conn.Write([]byte("-ERR empty command\r\n"))
			continue
		}

		
		fmt.Println("Parsed command:", parsedMessage)

	
		if len(parsedMessage) == 0 || parsedMessage[0] == "" {
			conn.Write([]byte("-ERR empty command\r\n"))
			continue
		}

	
		response, err := ExecuteCommand(parsedMessage)
		if err != nil {
			conn.Write([]byte("-" + err.Error() + "\r\n")) 
		} else if response == "" {
			conn.Write([]byte("$-1\r\n")) 
		} else {
			conn.Write([]byte("+" + response + "\r\n"))
		}
	}
}
