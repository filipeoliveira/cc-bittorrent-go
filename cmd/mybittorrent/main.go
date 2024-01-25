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

const (
	Announce    = "announce"
	Info        = "info"
	Length      = "length"
	PieceLength = "piece length"
	Pieces      = "pieces"
)

func parseDecodedTorrent(decoded interface{}) error {

	// Get announce
	decodedMap, err := utils.AssertMap(decoded)
	if err != nil {
		return err
	}

	announce, err := utils.AssertString(decodedMap[Announce])
	if err != nil {
		return err
	}
	fmt.Println("Tracker URL:", announce)

	// Get Length
	info, err := utils.AssertMap(decodedMap[Info])
	if err != nil {
		return err
	}

	length, err := utils.AssertInt(info[Length])
	if err != nil {
		return err
	}
	fmt.Println("Length:", length)

	// Get Info
	infoBencode, err := encode.Bencode(info)
	if err != nil {
		return err
	}

	infoHash := utils.EncodeToStringSha1(infoBencode)
	fmt.Println("Info Hash:", infoHash)

	// Get Piece Length
	pieceLength, err := utils.AssertInt(info[PieceLength])
	if err != nil {
		return err
	}
	fmt.Println("Piece Length:", pieceLength)

	// Get Pieces
	pieces, err := utils.AssertString(info[Pieces])
	if err != nil {
		return err
	}
	fmt.Println("Pieces:", pieces)

	return nil
}

func handleError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
