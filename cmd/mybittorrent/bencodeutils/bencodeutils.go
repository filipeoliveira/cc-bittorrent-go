package bencodeutils

import (
	"strconv"
	"unicode"
)

func IsInteger(str string) bool {
	return []rune(str)[0] == 'i'
}

func IsString(str string) bool {
	return unicode.IsDigit([]rune(str)[0])
}

func IsList(str string) bool {
	return []rune(str)[0] == 'l'
}

func IsDict(str string) bool {
	return []rune(str)[0] == 'd'
}

func parseIntegerFromChars(str string, start int) (int, int, error) {
	endDigitPos := start
	if str[endDigitPos] == '-' {
		endDigitPos++
	}
	for unicode.IsDigit(rune(str[endDigitPos])) {
		endDigitPos++
	}

	num, err := strconv.Atoi(str[start:endDigitPos])
	if err != nil {
		return 0, start, err
	}

	return num, endDigitPos, nil
}

func ParseInteger(str string) (int, int, error) {
	number, nextCharPos, err := parseIntegerFromChars(str, 1)
	if err != nil {
		return 0, 0, err
	}

	return number, nextCharPos + 1, nil
}

func ParseString(str string) (string, int, error) {
	stringLength, nextCharPos, err := parseIntegerFromChars(str, 0)
	if err != nil {
		return "", 0, err
	}

	start := nextCharPos + 1 // first char char after ':'
	end := start + stringLength
	parsedString := str[start:end]

	return parsedString, end, nil
}
