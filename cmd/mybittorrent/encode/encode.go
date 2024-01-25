package encode

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func Bencode(value interface{}) (string, error) {
	switch v := value.(type) {
	case int:
		val := bencodeInteger(v)
		return val, nil
	case string:
		val := bencodeString(v)
		return val, nil
	case map[string]interface{}:
		val, err := bencodeDict(v)
		if err != nil {
			return "", err
		}
		return val, nil
	case []interface{}:
		val := bencodeList(v)
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
	keys := make([]string, 0, len(dict))
	for key := range dict {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	parts := make([]string, 0, len(keys)*2)
	for _, key := range keys {
		value, ok := dict[key]
		if !ok {
			return "", fmt.Errorf("invalid dict: key %s not found", key)
		}

		// encode key
		keyEncoded := bencodeString(key)
		parts = append(parts, keyEncoded)

		// encode value
		switch v := value.(type) {
		case int:
			parts = append(parts, bencodeInteger(v))
		case map[string]interface{}:
			data, err := bencodeDict(v)
			if err != nil {
				return "", err
			}
			parts = append(parts, data)
		case string:
			parts = append(parts, bencodeString(v))
		default:
			return "", fmt.Errorf("unsupported value type: %T", v)
		}
	}

	return "d" + strings.Join(parts, "") + "e", nil
}
