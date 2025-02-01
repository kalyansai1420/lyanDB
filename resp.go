package main

import (
	"fmt"
	"strings"
)

// Serialize the RESP message (Convert into Redis protocol format)
func serializeRESP(message string) string {
	// We are handling Simple Strings for now (e.g., "+PONG\r\n")
	return fmt.Sprintf("+%s\r\n", message)
}

// Deserialize the RESP message (Convert from Redis protocol format)
func deserializeRESP(data []byte) string {
	// For simplicity, we're assuming the message is a simple string (e.g., "+PING\r\n")
	message := strings.TrimSpace(string(data))
	if strings.HasPrefix(message, "*") {
		messageParts := strings.Split(message, "\r\n")
		if len(messageParts) >= 2 {
			return messageParts[2] // Extract the command, e.g., "PING"
		}
	}
	return ""
}
