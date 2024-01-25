package utils

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
)

func AssertMap(val interface{}) (map[string]interface{}, error) {
	m, ok := val.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("expected map, got %T", val)
	}
	return m, nil
}

func AssertString(val interface{}) (string, error) {
	s, ok := val.(string)
	if !ok {
		return "", fmt.Errorf("expected string, got %T", val)
	}
	return s, nil
}

func AssertInt(val interface{}) (int, error) {
	i, ok := val.(int)
	if !ok {
		return 0, fmt.Errorf("expected int, got %T", val)
	}
	return i, nil
}

func GetMapProperty(mapValue map[string]interface{}, propertyName string) (interface{}, error) {
	propertyValue, exists := mapValue[propertyName]
	if !exists {
		return nil, fmt.Errorf("Property '%s' does not exist", propertyName)
	}
	return propertyValue, nil
}

func EncodeToStringSha1(str string) string {
	hasher := sha1.New()
	hasher.Write([]byte(str))
	sha := hasher.Sum(nil)
	return hex.EncodeToString(sha)
}

func ConvertToHexStrings(chunks [][]byte) []string {
	var hexStrings []string

	for _, chunk := range chunks {
		hexString := hex.EncodeToString(chunk)
		hexStrings = append(hexStrings, hexString)
	}

	return hexStrings
}

func SplitIntoChunks(data string, chunkSize int) [][]byte {
	byteData := []byte(data)
	var chunks [][]byte

	for i := 0; i < len(byteData); i += chunkSize {
		end := i + chunkSize

		// Ensure not to go beyond the data length
		if end > len(byteData) {
			end = len(byteData)
		}

		chunks = append(chunks, byteData[i:end])
	}

	return chunks
}
