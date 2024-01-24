package encode

import (
	"fmt"
	"strconv"
	"strings"
)

func Bencode(value interface{}) (string, error) {

	switch v := value.(type) {
	case map[string]interface{}:
		val, err := bencodeDict(v)
		if err != nil {
			return "", err
		}
		return val, nil
	}

	return "not suported...", fmt.Errorf("not supported")
}

func bencodeString(str string) string {
	return strconv.Itoa(len(str)) + ":" + str
}

func bencodeInteger(number int) string {
	return "i" + strconv.Itoa(number) + "e"
}

func bencodeList(list []interface{}) string {
	parts := []string{}

	for _, el := range list {
		switch v := el.(type) {
		case string:
			val := bencodeString(v)
			parts = append(parts, val)
		case int:
			val := bencodeInteger(v)
			parts = append(parts, val)
		}

	}
	return "l" + strings.Join(parts, "") + "e"
}

func bencodeDict(dict map[string]interface{}) (string, error) {
	parts := []string{}

	for k, v := range dict {
		key := bencodeString(k)
		parts = append(parts, key)

		switch v := v.(type) {
		case string:
			val := bencodeString(v)
			parts = append(parts, val)
		case int:
			val := bencodeInteger(v)
			parts = append(parts, val)
		case []interface{}:
			val := bencodeList(v)
			parts = append(parts, val)
		default:
			return "", fmt.Errorf("unsupported type: %T", v)
		}
	}

	return "d" + strings.Join(parts, "") + "e", nil
}
