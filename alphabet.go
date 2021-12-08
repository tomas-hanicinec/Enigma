package main

type char byte

// todo - better handle aplhabet injection / configuration (alphabet is part of enigma type config)
var Alphabet = newAlphabet("ABCDEFGHIJKLMNOPQRSTUVWXYZ")

type alphabet struct {
	letterMap string
	indexMap  map[char]int
}

func newAlphabet(alphabetString string) alphabet {
	indexMap := make(map[char]int)
	for i, letter := range alphabetString {
		indexMap[char(letter)] = i
	}

	return alphabet{
		letterMap: alphabetString,
		indexMap:  indexMap,
	}
}

func (a *alphabet) charToInt(letter char) (int, bool) {
	val, ok := a.indexMap[letter]
	return val, ok
}

func (a *alphabet) intToChar(index int) char {
	return char(a.letterMap[index])
}

func (a *alphabet) getSize() int {
	return len(a.letterMap)
}

func (a *alphabet) isValidWiring(wiring string) bool {
	// todo - do this better so we don't need "special" utils
	return sortString(wiring) == sortString(Alphabet.letterMap)
}
