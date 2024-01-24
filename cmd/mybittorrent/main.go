package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"unicode"

	"github.com/codecrafters-io/bittorrent-starter-go/cmd/mybittorrent/decode"
)

func parseLists(str string) (interface{}, error) {
	m := map[string]string{
		"i": "e",
		"l": "e",
	}

	stack := make([]string, 0)
	curr := ""

	for i := 0; i < len(str); i++ {
		char := str[i]
		if complement, ok := m[string(char)]; ok {
			stack = append(stack, complement)
			curr += string(char)
		} else if unicode.IsDigit(rune(char)) {
			for i < len(str) && unicode.IsDigit(rune(str[i])) {
				curr += string(str[i])
				i++
			}
			stack = append(stack, curr)
			curr = ""
		} else {
			return nil, fmt.Errorf("unexpected character: %c", char)
		}
	}

	if curr != "" {
		stack = append(stack, curr)
	}

	return stack, nil
}

func decodeInt(str string) (int, int, error) {
	start := 0
	endNumberPos := start + 1
	for str[endNumberPos] != 'e' {
		endNumberPos++
	}

	numString := str[start:endNumberPos]
	num, err := strconv.Atoi(numString)
	if err != nil {
		return 0, 0, err
	}

	return num, endNumberPos + 1, nil
}

func main() {
	command := os.Args[1]

	if command == "decode" {
		bencodedValue := os.Args[2]

		decoded, _, err := decode.Debencode(bencodedValue)
		if err != nil {
			fmt.Println(err)
			return
		}

		jsonOutput, _ := json.Marshal(decoded)
		fmt.Println(string(jsonOutput))
	} else {
		fmt.Println("Unknown command: " + command)
		os.Exit(1)
	}
}
