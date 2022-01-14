package enigma

import (
	"fmt"
	"strings"
)

type plugboard struct {
	isConfigurable bool
	letterMap      map[int]int
}

func newPlugboard(isConfigurable bool) plugboard {
	return plugboard{
		isConfigurable: isConfigurable,
		letterMap:      getDefaultLetterMap(),
	}
}

func (pb *plugboard) setup(plugConfig string) error {
	if !pb.isConfigurable {
		return fmt.Errorf("plugboard is locked, cannot configure")
	}

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
			letters[i], ok = Alphabet.charToInt(pair[i])
			if !ok {
				return fmt.Errorf("invalid pair %s, unsupported letter %s", pair, string(pair[i]))
			}
			if mapped, ok := letterMap[letters[i]]; ok && mapped != letters[i] {
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
