package bencodeutils

import (
	"strconv"
	"unicode"
)

func IsInteger(str string, i int) bool {
	return str[i] == 'i'
}

func IsString(str string, i int) bool {
	return unicode.IsDigit(rune(str[i]))
}

func IsList(str string, i int) bool {
	return str[i] == 'l'
}

func IsDict(str string, i int) bool {
	return str[i] == 'd'
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

func ParseInteger(str string, i int) (int, int, error) {
	number, nextCharPos, err := parseIntegerFromChars(str, i)
	if err != nil {
		return 0, 0, err
	}

	return number, nextCharPos + 1, nil
}

func ParseString(str string, i int) (string, int, error) {
	stringLength, nextCharPos, err := parseIntegerFromChars(str, i)
	if err != nil {
		return "", 0, err
	}

	start := nextCharPos + 1 // first char char after ':'
	end := start + stringLength
	parsedString := str[start:end]

	return parsedString, end, nil
}
