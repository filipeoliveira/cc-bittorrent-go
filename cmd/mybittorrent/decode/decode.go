package decode

import (
	"fmt"
	"strconv"

	"github.com/codecrafters-io/bittorrent-starter-go/cmd/mybittorrent/bencodeutils"
)

func Debencode(str string) (interface{}, int, error) {
	if bencodeutils.IsInteger(str) {
		return decodeInteger(str)
	} else if bencodeutils.IsString(str) {
		return decodeString(str)
	} else if bencodeutils.IsList(str) {
		return decodeList(str)
	} else if bencodeutils.IsDict(str) {
		return decodeDict(str)
	} else {
		return "", -1, fmt.Errorf("can't parse string anymore %s", str)
	}
}

func decodeInteger(str string) (interface{}, int, error) {
	value, end, err := bencodeutils.ParseInteger(str)
	if err != nil {
		return nil, 0, err
	}

	return value, end, nil
}

func decodeString(str string) (interface{}, int, error) {
	value, end, err := bencodeutils.ParseString(str)
	if err != nil {
		return nil, 0, err
	}
	return value, len([]byte(str[:end])), nil
}

func decodeList(str string) (interface{}, int, error) {
	start := 1 // getting next char after 'l'
	output := []interface{}{}

	if str[start] == 'e' {
		return output, start + 1, nil
	}

	for start < len(str) {
		value, end, err := Debencode(str[start:])
		if err != nil {
			return nil, 0, err
		}

		output = append(output, value)
		curr := start + end

		if str[curr] == 'e' {
			return output, curr + 1, nil
		}

		start = curr
	}

	return output, -1, fmt.Errorf("unexpected end of string while decoding list")
}

func decodeDict(str string) (interface{}, int, error) {
	start := 1 // getting next char after 'd'
	output := map[string]interface{}{}
	missingKey := true
	var lastKey string

	if str[start] == 'e' {
		return output, start + 1, nil
	}

	for start < len(str) {
		value, end, err := Debencode(str[start:])
		if err != nil {
			return nil, 0, fmt.Errorf("error decoding dict: %w", err)
		}

		if missingKey {
			switch v := value.(type) {
			case int:
				lastKey = strconv.Itoa(v)
			case string:
				lastKey = v
			default:
				return nil, 0, fmt.Errorf("expected key to be string or int, got %T", v)
			}

			missingKey = false
		} else {
			output[lastKey] = value
			missingKey = true
		}

		curr := start + end

		if str[curr] == 'e' {
			return output, curr + 1, nil
		}

		start = curr
	}

	return output, -1, fmt.Errorf("unexpected end of string while decoding dict")
}
