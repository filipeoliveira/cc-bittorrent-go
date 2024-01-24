package utils

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
)

func AssertType(value interface{}, targetType string) (map[string]interface{}, error) {
	decodedMap, ok := value.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("Expected a %s, got: %T", targetType, value)
	}
	return decodedMap, nil
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
