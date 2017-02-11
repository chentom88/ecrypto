package cipher

import (
	"errors"
	"math"
	"strings"
	"unicode"
)

func scoreLangString(input string, info *langInfo) (float64, error) {
	if info == nil {
		return 0.0, errors.New("Language info is required")
	}

	if len(input) == 0 {
		return 0, nil
	}

	upperInput := strings.ToUpper(input)
	inputLen := len(upperInput)
	noSpaceLen := int64(len(strings.Replace(upperInput, " ", "", -1)))

	ltrCount := make(map[rune]*freqTally, len(info.letterFreqs))
	bigramCount := make(map[string]*freqTally, len(info.bigramFreqs))
	trigramCount := make(map[string]*freqTally, len(info.trigramFreqs))

	penalizer := 1.0
	for i, char := range upperInput {
		if penalizer > 10000 {
			return 0.0, nil
		}

		if !unicode.IsPrint(char) {
			penalizer *= 10
			continue
		}

		_, ok := info.letterFreqs[char]
		if !ok {
			continue
		}

		tallyLetter(ltrCount, char, noSpaceLen)

		if i+2 < inputLen {
			bigram := upperInput[i : i+2]
			tallyGram(bigramCount, info.bigramFreqs, bigram, noSpaceLen)
		}

		if i+3 < inputLen {
			trigram := upperInput[i : i+3]
			tallyGram(trigramCount, info.trigramFreqs, trigram, noSpaceLen)
		}
	}

	letterScore := info.letterWeight * scoreLetterFreqs(ltrCount, info.letterFreqs, penalizer)
	bigramScore := info.bigramWeight * scoreGramFreqs(bigramCount, info.bigramFreqs, penalizer)
	trigramScore := info.trigramWeight * scoreGramFreqs(trigramCount, info.trigramFreqs, penalizer)

	return letterScore + bigramScore + trigramScore, nil
}

func tallyLetter(counter map[rune]*freqTally, char rune, totalCount int64) {
	tempCnt, ok := counter[char]
	if !ok {
		counter[char] = &freqTally{
			count:   1,
			percent: float64(1 / totalCount),
		}
	} else {
		tempCnt.count += 1
		tempCnt.percent = float64(tempCnt.count / totalCount)
	}
}

func tallyGram(counter map[string]*freqTally, freqInfo map[string]*freqData, seq string, totalCount int64) {
	_, ok := freqInfo[seq]
	if !ok {
		return
	}

	tempCnt, ok := counter[seq]
	if !ok {
		counter[seq] = &freqTally{
			count:   1,
			percent: float64(1 / totalCount),
		}
	} else {
		tempCnt.count += 1
		tempCnt.percent = float64(tempCnt.count / totalCount)
	}
}

func scoreLetterFreqs(actual map[rune]*freqTally, expected map[rune]*freqData, penalizer float64) float64 {
	score := 0.0
	for char, freqData := range expected {
		tempCnt, ok := actual[char]
		if !ok {
			continue
		}

		diff := math.Abs(freqData.freq - (100 * tempCnt.percent))
		score += freqData.scoreValue / math.Max(1, diff)
	}

	return score / (100 * penalizer)
}

func scoreGramFreqs(actual map[string]*freqTally, expected map[string]*freqData, penalizer float64) float64 {
	score := 0.0
	for char, freqData := range expected {
		tempCnt, ok := actual[char]
		if !ok {
			continue
		}

		diff := math.Abs(freqData.freq - (100 * tempCnt.percent))
		score += freqData.scoreValue / math.Max(1, diff)
	}

	return score / (100 * penalizer)
}

type freqTally struct {
	count   int64
	percent float64
}
