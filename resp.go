package main

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

func SerializeRESP(value string) string {
	if value == "nil" {
		return "$-1/r/n"
	}
	return "+" + value + "\r\n"
}

func DeserializeRESP(reader *bufio.Reader) ([]string, error) {
	var result []string

	
	prefix, err := reader.ReadByte()
	if err != nil {
		return nil, err
	}

	switch prefix {
	case '*': 
		lengthStr, _ := reader.ReadString('\n')
		length, _ := strconv.Atoi(strings.TrimSpace(lengthStr))
		for i := 0; i < length; i++ {
			elem, err := DeserializeRESP(reader)
			if err != nil {
				return nil, err
			}
			result = append(result, elem...)
		}
		return result, nil

	case '$': 
		lengthStr, _ := reader.ReadString('\n')
		length, _ := strconv.Atoi(strings.TrimSpace(lengthStr))
		if length < 0 {
			return nil, nil
		}
		buf := make([]byte, length+2)
		_, err := io.ReadFull(reader, buf)
		if err != nil {
			return nil, err
		}
		return []string{string(buf[:length])}, nil

	default:
		
		reader.UnreadByte()
		line, _ := reader.ReadString('\n')
		return strings.Fields(strings.TrimSpace(line)), nil
	}
}

