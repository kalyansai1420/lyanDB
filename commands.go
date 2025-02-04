package main

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

func ExecuteCommand(parsedMessage []string) (string, error) {
	if len(parsedMessage) == 0 {
		return "", errors.New("ERR empty command")
	}

	command := strings.ToUpper(parsedMessage[0])

	switch command {
	case "PING":
		return "PONG", nil

	case "ECHO":
		if len(parsedMessage) < 2 {
			return "", errors.New("ERR missing argument for ECHO")
		}
		return strings.Join(parsedMessage[1:], " "), nil

	case "SET":
		if len(parsedMessage) < 3 {
			return "", errors.New("ERR wrong number of arguments for SET")
		}

		key := parsedMessage[1]
		value := parsedMessage[2]
		var expiryTime time.Time

		if len(parsedMessage) > 3 {
			if len(parsedMessage) < 5 {
				return "", errors.New("ERR syntax error for expiry option")
			}

			expiryOption := strings.ToUpper(parsedMessage[3])
			expiryValue, err := strconv.Atoi(parsedMessage[4])
			if err != nil {
				return "", errors.New("ERR invalid expiry time")
			}

			switch expiryOption {
			case "EX":
				expiryTime = time.Now().Add(time.Duration(expiryValue) * time.Second)
			case "PX":
				expiryTime = time.Now().Add(time.Duration(expiryValue) * time.Millisecond)
			case "EXAT":
				expiryTime = time.Unix(int64(expiryValue), 0)
			case "PXAT":
				expiryTime = time.Unix(0, int64(expiryValue)*int64(time.Millisecond))
			default:
				return "", errors.New("ERR unknown expiry option")
			}
		}

		SetKey(key, value, expiryTime)
		return "OK", nil

	case "GET":
		if len(parsedMessage) < 2 {
			return "", errors.New("ERR wrong number of arguments for GET")
		}

		key := parsedMessage[1]
		value, exists := GetKey(key)

		if !exists {
			return "ERR key not found", nil
		}

		return value, nil

	default:
		return "", errors.New("ERR unknown command")
	}
}
