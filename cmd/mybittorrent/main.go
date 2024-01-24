package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/codecrafters-io/bittorrent-starter-go/cmd/mybittorrent/decode"
)

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
	} else if command == "info" {
		filename := os.Args[2]

		fileContent, err := os.ReadFile(filename)
		if err != nil {
			fmt.Println(err)
			return
		}

		decoded, _, err := decode.Debencode(string(fileContent))
		if err != nil {
			fmt.Println(err)
			return
		}

		decodedMap, ok := decoded.(map[string]interface{})
		if !ok {
			fmt.Println("Decoded value is not a map")
			return
		}

		fmt.Println("Tracker URL:", decodedMap["announce"])

		info, ok := decodedMap["info"].(map[string]interface{})
		if !ok {
			fmt.Println("Property 'info' is not a map")
			return
		}

		length, exists := info["length"]
		if !exists {
			fmt.Println("Property 'length' does not exist")
			return
		}

		fmt.Println("Length:", length)

	} else {
		fmt.Println("Unknown command: " + command)
		os.Exit(1)
	}
}
