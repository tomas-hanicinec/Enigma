package main

import (
	"fmt"
	"strings"
)

type Plugboard struct {
	letterMap map[Char]Char
}

func (pb *Plugboard) Setup(plugConfig string) error {
	pairs := strings.Split(plugConfig, " ")
	if len(pairs) != 10 {
		return fmt.Errorf("wrong number of plugboard pairs")
	}

	letterMap := make(map[Char]Char, len(Alphabet)) // start with map pointing each letter of the alphabet to itself
	for _, letter := range Alphabet {
		letterMap[Char(letter)] = Char(letter)
	}

	// connect the plugs
	for _, pair := range pairs {
		// check the pair is valid
		if err := pb.validatePair(pair, letterMap); err != nil {
			return fmt.Errorf("invalid pair %s: %w", pair, err)
		}
		// assign to map (temporary map for now, wait for the whole config to finish to make sure there are no errors)
		letterMap[pair[0]] = pair[1]
		letterMap[pair[1]] = pair[0]
	}

	// all good, set the new map
	pb.letterMap = letterMap
	return nil
}

func (pb *Plugboard) Translate(letter Char) (Char, error) {
	res, ok := pb.letterMap[letter]
	if !ok {
		return 0, fmt.Errorf("letter %s not supported", string(letter))
	}
	return res, nil
}

func (pb *Plugboard) validatePair(pair string, letterMap map[Char]Char) error {
	if len(pair) != 2 {
		return fmt.Errorf("must be a pair of letters")
	}
	if pair[0] == pair[1] {
		return fmt.Errorf("cannot plug a letter to itself")
	}
	// check there are no duplicate connections
	for i := 0; i < 2; i++ {
		mapped, ok := letterMap[pair[i]]
		if !ok {
			return fmt.Errorf("letter %s not supported", string(pair[i]))
		}
		if mapped != pair[i] {
			return fmt.Errorf("letter %s already connected to %s", string(pair[i]), string(mapped))
		}
	}

	return nil
}
