package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"unicode"
	// bencode "github.com/jackpal/bencode-go" // Available if you need it!
)

func decodeBencode(bencodedString string) (interface{}, error) {

	lastPos := len(bencodedString) - 1

	if rune(bencodedString[0]) == 'i' && rune(bencodedString[lastPos]) == 'e' {

		endIndex := lastPos
		numberStr := bencodedString[1:endIndex]

		valid, err := strconv.Atoi(numberStr) // checking if numberStr is a valid integer
		if err != nil {
			return "", fmt.Errorf("invalid integer: %v", err)
		}

		return valid, nil
	}

	if unicode.IsDigit(rune(bencodedString[0])) {
		var firstColonIndex int

		for i := 0; i < len(bencodedString); i++ {
			if bencodedString[i] == ':' {
				firstColonIndex = i
				break
			}
		}
		lengthStr := bencodedString[:firstColonIndex]

		length, err := strconv.Atoi(lengthStr) // string to int.
		if err != nil {
			return "", err
		}

		return bencodedString[firstColonIndex+1 : firstColonIndex+1+length], nil
	} else {
		return "", fmt.Errorf("only strings are supported at the moment")
	}
}

func main() {
	command := os.Args[1]

	if command == "decode" {
		bencodedValue := os.Args[2]

		decoded, err := decodeBencode(bencodedValue)
		if err != nil {
			fmt.Println(err)
			return
		}

		// marshal ?
		jsonOutput, _ := json.Marshal(decoded)
		fmt.Println(string(jsonOutput))
	} else {
		fmt.Println("Unknown command: " + command)
		os.Exit(1)
	}
}
