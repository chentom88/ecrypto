package cipher

import "github.com/chentom88/ecrypto/hex"

func EncodeSingleByteXOR(hexInput string, key byte) ([]byte, error) {
	input, err := hex.HexStringToString(hexInput)
	if err != nil {
		return nil, err
	}

	result := make([]byte, len(input))

	for i, inputByte := range input {
		result[i] = byte(inputByte) ^ key
	}

	return result, nil
}

func CrackSingleByteXOR(hexInput string) ([]*sbxResult, error) {
	results := make([]*sbxResult, 255)

	for i := 0; i < 255; i++ {
		key := byte(i + 1)
		tempResult, err := EncodeSingleByteXOR(hexInput, key)
		if err != nil {
			continue
		}

		results[i] = &sbxResult{
			decrypted: string(tempResult),
			key:       key,
		}

		results[i].ranking = scoreResult(tempResult)
	}

	return results, nil
}

func scoreResult(input []byte) float64 {
	return 100.0
}

type sbxResult struct {
	decrypted string
	ranking   float64
	key       byte
}
