package base64

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/chentom88/ecrypto/hex"
)

const baseCode = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/="
const empty = ""

var codeMap = make(map[rune]byte, len(baseCode))

func init() {
	for i, char := range baseCode {
		codeMap[char] = byte(i)
	}
}

func Encode(input string) (string, error) {
	if len(input) == 0 {
		return empty, errors.New("no string to encode")
	}

	terminator := len(baseCode) - 1

	inputBytes := []byte(input)
	lenInput := len(inputBytes)
	outputBytes := make([]byte, 0)

	for i := 0; i < lenInput; i += 3 {
		byte1 := inputBytes[i] >> 2
		byte2 := (inputBytes[i] & 0x03) << 4
		byte3 := byte(terminator)
		byte4 := byte(terminator)

		if i+1 < lenInput {
			byte2 = byte2 | ((inputBytes[i+1] & 0xF0) >> 4)
			byte3 = (inputBytes[i+1] & 0x0F) << 2
		}

		if i+2 < lenInput {
			byte3 = byte3 | ((inputBytes[i+2] & 0xC0) >> 6)
			byte4 = inputBytes[i+2] & 0x3F
		}

		outputBytes = append(outputBytes, baseCode[byte1], baseCode[byte2], baseCode[byte3], baseCode[byte4])
	}

	return string(outputBytes), nil
}

func Decode(input string) (string, error) {
	inputLen := len(input)
	if inputLen == 0 {
		return empty, errors.New("no string to decode")
	}

	if !validateBase64String(input) {
		return empty, errors.New("Invalid base64 encoded string")
	}

	terminator := codeMap['=']
	var result []byte
	for i := 0; i < inputLen; i += 4 {
		input1 := codeMap[rune(input[i])]
		input2 := codeMap[rune(input[i+1])]
		input3 := codeMap[rune(input[i+2])]
		input4 := codeMap[rune(input[i+3])]

		result = append(result, (input1<<2)|(input2>>4))
		if input3 != terminator {
			result = append(result, (input2<<4)|(input3>>2))

			if input4 != terminator {
				result = append(result, (input3<<6)|input4)
			}
		}
	}

	return string(result), nil
}

func EncodeHex(input string) (string, error) {
	asciiString, err := hex.HexStringToString(input)
	if err != nil {
		return empty, err
	}

	return Encode(asciiString)
}

func DecodeToHex(input string) (string, error) {
	if len(input) == 0 {
		return empty, errors.New("no string to decode")
	}

	decodedString, err := Decode(input)
	if err != nil {
		return empty, err
	}

	hexString, err := hex.StringToHexString(decodedString)
	if err != nil {
		return empty, err
	}

	return hexString, nil
}

func validateBase64String(input string) bool {
	inputLen := len(input)

	if inputLen == 0 {
		return true
	}

	if inputLen%4 != 0 {
		return false
	}

	indexTerminator := strings.IndexByte(input, '=')
	if indexTerminator >= 0 && indexTerminator < len(input)-2 {
		return false
	}

	matchString := fmt.Sprintf("[^%s]+", baseCode)
	matched, err := regexp.MatchString(matchString, input)

	if err != nil {
		return false
	}

	return !matched
}
