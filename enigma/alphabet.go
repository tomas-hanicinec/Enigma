package enigma

import (
	"sort"
	"strings"
)

var Alphabet = newAlphabet("ABCDEFGHIJKLMNOPQRSTUVWXYZ")

type alphabet struct {
	letterMap string
	indexMap  map[byte]int
}

func newAlphabet(alphabetString string) alphabet {
	indexMap := make(map[byte]int)
	for i, letter := range alphabetString {
		indexMap[byte(letter)] = i
	}

	return alphabet{
		letterMap: alphabetString,
		indexMap:  indexMap,
	}
}

func (a *alphabet) charToInt(letter byte) (int, bool) {
	val, ok := a.indexMap[letter]
	return val, ok
}

func (a *alphabet) intToChar(index int) byte {
	return a.letterMap[index]
}

func (a *alphabet) getSize() int {
	return len(a.letterMap)
}

func (a *alphabet) isValidWiring(wiring string) bool {
	return sortString(wiring) == sortString(Alphabet.letterMap)
}

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func sortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}

// shift moves the given letter (specified by its index) byt the specified amount in the alphabet (Z wraps around back to A)
// accepts both positive and negative numbers, and it's cyclical (Z wraps around back to A and A back to Z)
func shift(input int, shiftBy int) int {
	result := (input + shiftBy) % Alphabet.getSize()
	if result < 0 {
		result = Alphabet.getSize() + result
	}
	return result
}

// getDefaultLetterMap generates mapping of each letter in the alphabet to itself
func getDefaultLetterMap() map[int]int {
	letterMap := make(map[int]int, Alphabet.getSize())
	for i := 0; i < Alphabet.getSize(); i++ {
		letterMap[i] = i
	}
	return letterMap
}

func Preprocess(text string) string {
	text = strings.ToUpper(text) // convert to uppercase
	// replace punctuations with double letters
	// todo - do this more elaborately (currently COMPLEX PROCEDURE -> COMPLE XPROCEDURE)
	text = strings.ReplaceAll(text, " ", "XX")
	text = strings.ReplaceAll(text, ".", "YY")
	text = strings.ReplaceAll(text, ",", "ZZ")

	return text
}

func Postprocess(text string) string {
	// convert back the punctuation
	text = strings.ReplaceAll(text, "XX", " ")
	text = strings.ReplaceAll(text, "YY", ".")
	text = strings.ReplaceAll(text, "ZZ", ",")
	return text
}
