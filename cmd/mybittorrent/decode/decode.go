package decode

import (
	"strconv"

	"github.com/codecrafters-io/bittorrent-starter-go/cmd/mybittorrent/bencodeutils"
)

func Debencode(str string) (interface{}, int, error) {
	start := 0

	if bencodeutils.IsInteger(str, start) {
		return decodeInteger(str, start)
	} else if bencodeutils.IsString(str, start) {
		return decodeString(str, start)
	} else if bencodeutils.IsList(str, start) {
		return decodeList(str, start)
	} else if bencodeutils.IsDict(str, start) {
		return decodeDict(str, start)
	}

	return "", -1, nil
}

func decodeInteger(str string, start int) (interface{}, int, error) {
	start++ // jumping first char (i)
	value, end, err := bencodeutils.ParseInteger(str, start)
	if err != nil {
		return nil, 0, err
	}

	return value, end, nil
}

func decodeString(str string, start int) (interface{}, int, error) {
	value, end, err := bencodeutils.ParseString(str, start)
	if err != nil {
		return nil, 0, err
	}
	return value, end, nil
}

func decodeList(str string, start int) (interface{}, int, error) {
	start++
	output := []interface{}{}

	if str[start] == 'e' {
		return output, start + 1, nil
	}

	for start < len(str) {
		value, end, err := Debencode(str[start:])
		output = append(output, value)

		curr := start + end

		if err != nil {
			return nil, 0, err
		}

		if string(str[curr]) == "e" {
			return output, curr + 1, nil
		}

		start = curr
	}

	return output, -1, nil
}

func decodeDict(str string, start int) (interface{}, int, error) {
	start++
	output := map[string]interface{}{}
	missingKey := true
	var lastKey string

	if str[start] == 'e' {
		return output, start + 1, nil
	}

	for start < len(str) {
		value, end, err := Debencode(str[start:])
		if missingKey {
			lastKey, _ = value.(string)

			if intValue, ok := value.(int); ok {
				lastKey = strconv.Itoa(intValue)
			}

			missingKey = false
		} else {
			output[lastKey] = value
			missingKey = true
		}

		curr := start + end

		if err != nil {
			return nil, 0, err
		}

		if string(str[curr]) == "e" {
			return output, curr + 1, nil
		}

		start = curr
	}

	return output, -1, nil
}
