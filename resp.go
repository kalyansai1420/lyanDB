package main

import (
	"errors"
	"strconv"
	"strings"
)

func serializeRESP(value string) string {
	if value == "nil" {
		return "$-1/r/n"
	}
	return "+" + value + "\r\n"
}

func deserializeRESP(input string) ([]string, error) {
	lines := strings.Split(strings.TrimSpace(input), "\r\n")
	if len(lines) < 1 {
		return nil, errors.New("invalid RESP format")
	}

	if lines[0][0] == '*' {
		numArgs, _ := strconv.Atoi(lines[0][1:])
		parsed := make([]string, 0, numArgs)
		for i := 2; i < len(lines); i += 2 {
			parsed = append(parsed, lines[i])
		}
		return parsed, nil
	}

	return []string{lines[0][1:]}, nil
}
