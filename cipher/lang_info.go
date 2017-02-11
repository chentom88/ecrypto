package cipher

type langInfo struct {
	letterFreqs   map[rune]*freqData
	bigramFreqs   map[string]*freqData
	trigramFreqs  map[string]*freqData
	letterTotal   float64
	bigramTotal   float64
	trigramTotal  float64
	letterWeight  float64
	bigramWeight  float64
	trigramWeight float64
}

func init() {
	initLangInfo(&englishInfo)
}

func initLangInfo(info *langInfo) {
	for _, f := range info.letterFreqs {
		info.letterTotal += f.freq
	}

	for _, f := range info.letterFreqs {
		f.scoreValue = (f.freq / info.letterTotal) * 100
	}

	for _, f := range info.bigramFreqs {
		info.bigramTotal += f.freq
	}

	for _, f := range info.bigramFreqs {
		f.scoreValue = (f.freq / info.bigramTotal) * 100
	}

	for _, f := range info.trigramFreqs {
		info.trigramTotal += f.freq
	}

	for _, f := range info.trigramFreqs {
		f.scoreValue = (f.freq / info.trigramTotal) * 100
	}
}

func englishLangInfo() *langInfo {
	return &englishInfo
}

type freqData struct {
	freq       float64
	scoreValue float64
}

var englishInfo = langInfo{
	letterFreqs: map[rune]*freqData{
		'E': &freqData{freq: 12.7},
		'T': &freqData{freq: 9.05},
		'A': &freqData{freq: 8.17},
		'O': &freqData{freq: 7.51},
		'I': &freqData{freq: 6.97},
		'N': &freqData{freq: 6.75},
		'S': &freqData{freq: 6.33},
		'H': &freqData{freq: 6.09},
		'R': &freqData{freq: 5.99},
		'D': &freqData{freq: 4.25},
		'L': &freqData{freq: 4.03},
		'C': &freqData{freq: 2.78},
		'U': &freqData{freq: 2.76},
		'M': &freqData{freq: 2.41},
		'W': &freqData{freq: 2.36},
		'F': &freqData{freq: 2.23},
		'G': &freqData{freq: 2.02},
		'Y': &freqData{freq: 1.97},
		'P': &freqData{freq: 1.93},
		'B': &freqData{freq: 1.49},
		'V': &freqData{freq: 0.98},
		'K': &freqData{freq: 0.77},
		'J': &freqData{freq: 0.15},
		'X': &freqData{freq: 0.15},
		'Q': &freqData{freq: 0.10},
		'Z': &freqData{freq: 0.07},
	},
	letterWeight: 33.3,
	bigramFreqs: map[string]*freqData{
		"TH": &freqData{freq: 2.71},
		"HE": &freqData{freq: 2.33},
		"IN": &freqData{freq: 2.03},
		"ER": &freqData{freq: 1.78},
		"AN": &freqData{freq: 1.62},
		"RE": &freqData{freq: 1.41},
		"ES": &freqData{freq: 1.32},
		"ON": &freqData{freq: 1.32},
		"ST": &freqData{freq: 1.25},
		"NT": &freqData{freq: 1.17},
		"EN": &freqData{freq: 1.13},
		"AT": &freqData{freq: 1.12},
		"ED": &freqData{freq: 1.08},
		"ND": &freqData{freq: 1.07},
		"TO": &freqData{freq: 1.07},
	},
	bigramWeight: 33.3,
	trigramFreqs: map[string]*freqData{
		"THE": &freqData{freq: 1.81},
		"AND": &freqData{freq: 0.73},
		"ING": &freqData{freq: 0.72},
		"ENT": &freqData{freq: 0.42},
		"ION": &freqData{freq: 0.42},
		"HER": &freqData{freq: 0.36},
		"FOR": &freqData{freq: 0.34},
		"THA": &freqData{freq: 0.33},
		"NTA": &freqData{freq: 0.33},
		"INT": &freqData{freq: 0.32},
	},
	trigramWeight: 33.3,
}
