package hex

import (
	"errors"
	"regexp"
)

func HexStringToString(input string) (string, error) {
	if len(input) == 0 {
		return "", errors.New("no string to encode")
	}

	if !validateHexString(input) {
		return "", errors.New("invalid hex string")
	}

	output := make([]byte, len(input)/2)

	for i := 0; i < len(output); i++ {
		output[i] = (hexMap[input[2*i]] << 4) | hexMap[input[2*i+1]]
	}

	return string(output), nil
}

func StringToHexString(input string) (string, error) {
	if len(input) == 0 {
		return "", errors.New("no string to encode")
	}

	hexString := make([]byte, len(input)*2)
	for i, char := range input {
		hexString[2*i] = hexMap[byte(char)>>4]
		hexString[(2*i)+1] = hexMap[byte(char)&0x0f]
	}

	return string(hexString), nil
}

func FixedXOR(input1, input2 string) (string, error) {
	inputLen1 := len(input1)
	inputLen2 := len(input2)

	if inputLen1 == 0 || inputLen2 == 0 || (inputLen1 != inputLen2) {
		return "", errors.New("invalid input")
	}

	if !validateHexString(input1) || !validateHexString(input2) {
		return "", errors.New("invalid hex strings")
	}

	converted1, err := HexStringToString(input1)
	if err != nil {
		return "", err
	}

	converted2, err := HexStringToString(input2)
	if err != nil {
		return "", err
	}

	convertedLen := len(converted1)
	result := make([]byte, convertedLen)
	for i := 0; i < convertedLen; i++ {
		result[i] = converted1[i] ^ converted2[i]
	}

	return StringToHexString(string(result))
}

func validateHexString(input string) bool {
	if len(input)%2 != 0 {
		return false
	}

	matched, err := regexp.MatchString("[^0-9abcdefABCDEF]+", input)

	if err != nil {
		return false
	}

	return !matched
}

var hexMap = map[byte]byte{
	'0':  0x00,
	'1':  0x01,
	'2':  0x02,
	'3':  0x03,
	'4':  0x04,
	'5':  0x05,
	'6':  0x06,
	'7':  0x07,
	'8':  0x08,
	'9':  0x09,
	'a':  0x0a,
	'A':  0x0a,
	'b':  0x0b,
	'B':  0x0b,
	'c':  0x0c,
	'C':  0x0c,
	'd':  0x0d,
	'D':  0x0d,
	'e':  0x0e,
	'E':  0x0e,
	'f':  0x0f,
	'F':  0x0f,
	0x00: '0',
	0x01: '1',
	0x02: '2',
	0x03: '3',
	0x04: '4',
	0x05: '5',
	0x06: '6',
	0x07: '7',
	0x08: '8',
	0x09: '9',
	0x0a: 'a',
	0x0b: 'b',
	0x0c: 'c',
	0x0d: 'd',
	0x0e: 'e',
	0x0f: 'f',
}
