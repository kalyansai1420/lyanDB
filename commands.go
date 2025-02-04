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

	case "EXISTS":
		if len(parsedMessage) < 2 {
			return "", errors.New("ERR wrong number of arguements for EXISTS")
		}
		key := parsedMessage[1]
		exists := checkIfKeyExists(key)

		if exists {
			return "1", nil
		}
		return "0", nil

	case "DEL":
		if len(parsedMessage) < 2 {
			return "",errors.New("ERR wrong number of arguments for DEL")
		}
		deletedCount := 0
		for _, key := range parsedMessage[1:] {
			if deleteKey(key) {
				deletedCount++
			}
		}
		return strconv.Itoa(deletedCount),nil

	case "INCR":
		if len(parsedMessage) < 2 {
			return "",errors.New("ERR wrong number of arguments for INCR")
		}

		key := parsedMessage[1]
		currentValue, exists := GetKey(key)

		if !exists {
			return "", errors.New("ERR key does not exists")
		}

		intValue, err := strconv.Atoi(currentValue)
		if err != nil {
			return "",errors.New("ERR value is not an integer")
		}
		newValue := intValue +1
		SetKey(key,strconv.Itoa(newValue),time.Time{})
		return strconv.Itoa(newValue),nil

	case "DECR":
		if len(parsedMessage) < 2 {
			return "",errors.New("ERR wrong number of arguments for INCR")
		}

		key := parsedMessage[1]
		currentValue, exists := GetKey(key)

		if !exists {
			return "", errors.New("ERR key does not exists")
		}

		intValue, err := strconv.Atoi(currentValue)
		if err != nil {
			return "",errors.New("ERR value is not an integer")
		}
		newValue := intValue -1
		SetKey(key,strconv.Itoa(newValue),time.Time{})
		return strconv.Itoa(newValue),nil

	case "LPUSH":
		if len(parsedMessage) < 3 {
			return "", errors.New("ERR wrong number of arguments for LPUSH")
		}
	
		key := parsedMessage[1]
		values := parsedMessage[2:]
	
		existingValue, exists := GetKey(key)
		var list []string
	
		if exists {
			list = strings.Split(existingValue, ",") 
		}
	
		list = append(values, list...)
	
		SetKey(key, strings.Join(list, ","), time.Time{})
		return strconv.Itoa(len(list)), nil
	
	case "RPUSH":
		if len(parsedMessage) < 3 {
			return "", errors.New("ERR wrong number of arguments for RPUSH")
		}
	
		key := parsedMessage[1]
		values := parsedMessage[2:]
	
		existingValue, exists := GetKey(key)
		var list []string
	
		if exists {
			list = strings.Split(existingValue, ",")
		}
	
		list = append(list, values...)
	
		SetKey(key, strings.Join(list, ","), time.Time{})
		return strconv.Itoa(len(list)), nil
	
	case "SAVE":
		err := SaveDatabase("lyanDB.rdb")
		if err != nil {
			return "", errors.New("ERR failed to save database")
		}
		return "OK", nil
	case "CONFIG":
		if len(parsedMessage) < 2 {
			return "", errors.New("ERR wrong number of arguments for CONFIG")
		}
		
		subCommand := strings.ToUpper(parsedMessage[1])
		if subCommand == "GET" {
			if len(parsedMessage) < 3 {
				return "", errors.New("ERR wrong number of arguments for CONFIG GET")
			}
			key := parsedMessage[2]
			
			
			if key == "databases" {
				return "1", nil
			}
			return "", errors.New("ERR unsupported CONFIG GET key")
		}
		return "", errors.New("ERR unknown CONFIG subcommand")

	default:
		return "", errors.New("ERR unknown command")
	}
}
