package cipher

import (
	"sort"

	"github.com/chentom88/ecrypto/hex"
)

func EncodeSingleByteXORString(input string, key string) (string, error) {
	if len(input) == 0 {
		return "", nil
	}

	resultByte := singleByteXOR(input, key[0])
	return hex.StringToHexString(string(resultByte))
}

func EncodeSingleByteXOR(hexInput string, key byte) ([]byte, error) {
	input, err := hex.HexStringToString(hexInput)
	if err != nil {
		return nil, err
	}

	return singleByteXOR(input, key), nil
}

func CrackSingleByteXOR(hexInput string) ([]*SbxResult, error) {
	results := make([]*SbxResult, 255)

	for i := 0; i < 255; i++ {
		key := byte(i + 1)
		tempResult, err := EncodeSingleByteXOR(hexInput, key)
		if err != nil {
			continue
		}

		results[i] = &SbxResult{
			Decrypted: string(tempResult),
			Key:       key,
			Ranking:   0.0,
		}

		results[i].Ranking, _ = scoreLangString(string(tempResult), &englishInfo)
	}

	sort.Sort(byScore(results))
	return results[:5], nil
}

func singleByteXOR(input string, key byte) []byte {
	result := make([]byte, len(input))

	for i, inputByte := range input {
		result[i] = byte(inputByte) ^ key
	}

	return result
}

type SbxResult struct {
	Decrypted string
	Ranking   float64
	Key       byte
}

type byScore []*SbxResult

func (s byScore) Len() int           { return len(s) }
func (s byScore) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s byScore) Less(i, j int) bool { return s[i].Ranking > s[j].Ranking }
