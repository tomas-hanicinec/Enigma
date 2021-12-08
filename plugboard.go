package main

import (
	"fmt"
	"strings"
)

type plugboard struct {
	letterMap map[int]int
}

func newPlugboard() plugboard {
	return plugboard{
		letterMap: getDefaultLetterMap(),
	}
}

func (pb *plugboard) setup(plugConfig string) error {
	// start with default map
	letterMap := getDefaultLetterMap()

	// connect the plugs
	pairs := strings.Split(plugConfig, " ")
	for _, pair := range pairs {
		// validate the pair
		if len(pair) != 2 {
			return fmt.Errorf("invalid pair %s, must be a pair of letters", pair)
		}
		if pair[0] == pair[1] {
			return fmt.Errorf("invalid pair %s, cannot plug a letter to itself", pair)
		}
		var letters [2]int
		ok := false
		for i := 0; i < 2; i++ {
			letters[i], ok = Alphabet.charToInt(char(pair[0]))
			if !ok {
				return fmt.Errorf("invalid pair %s, unsupported letter %s", pair, string(pair[0]))
			}
			if _, ok = letterMap[letters[i]]; ok {
				return fmt.Errorf("invalid pair %s, letter %s already connected", pair, string(pair[i]))
			}
		}

		// set to map (both directions)
		letterMap[letters[0]] = letters[1]
		letterMap[letters[1]] = letters[0]
	}

	// all good, set the new map to plugboard
	pb.letterMap = letterMap
	return nil
}

func (pb *plugboard) translate(letter int) int {
	return pb.letterMap[letter]
}

func getDefaultLetterMap() map[int]int {
	letterMap := make(map[int]int, Alphabet.getSize())
	for i := 0; i < Alphabet.getSize(); i++ {
		letterMap[i] = i
	}
	return letterMap
}
