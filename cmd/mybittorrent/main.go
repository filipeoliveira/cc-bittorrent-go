package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/codecrafters-io/bittorrent-starter-go/cmd/mybittorrent/decode"
	"github.com/codecrafters-io/bittorrent-starter-go/cmd/mybittorrent/encode"
	"github.com/codecrafters-io/bittorrent-starter-go/cmd/mybittorrent/utils"
)

func main() {
	command := os.Args[1]

	if command == "decode" {
		decoded, _, err := decode.Debencode(os.Args[2])
		handleError(err)

		jsonOutput, _ := json.Marshal(decoded)
		fmt.Println(string(jsonOutput))

	} else if command == "info" {
		fileContent, err := os.ReadFile(os.Args[2])
		handleError(err)

		decoded, _, err := decode.Debencode(string(fileContent))
		handleError(err)

		parseDecodedTorrent(decoded)
	} else {
		fmt.Println("Unknown command: " + command)
		os.Exit(1)
	}
}

func parseDecodedTorrent(decoded interface{}) {

	// Get announce
	decodedMap, err := utils.AssertType(decoded, "map")
	handleError(err)
	fmt.Println("Tracker URL:", decodedMap["announce"])

	// Get Length
	info, err := utils.GetMapProperty(decodedMap, "info")
	handleError(err)

	infoMap, err := utils.AssertType(info, "map")
	handleError(err)

	length, err := utils.GetMapProperty(infoMap, "length")
	handleError(err)
	fmt.Println("Length:", length)

	// Get Info
	infoBencode, err := encode.Bencode(info)
	handleError(err)

	infoHash := utils.EncodeToStringSha1(infoBencode)
	fmt.Println("Info Hash:", infoHash)

}

func handleError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
