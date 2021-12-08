package main

import (
	"fmt"
)

type reflector struct {
	translationMap map[int]int
}

func newReflector(wiring string) (reflector, error) {
	if !Alphabet.isValidWiring(wiring) {
		return reflector{}, fmt.Errorf("invalid reflector wiring %s", wiring)
	}

	letterMap := make(map[int]int, Alphabet.getSize())
	for i, letter := range wiring {
		letterIndex, ok := Alphabet.charToInt(char(letter))
		if !ok {
			return reflector{}, fmt.Errorf("unsupported wiring letter %s", string(letter)) // should not happen, we already checked the wiring validity
		}
		letterMap[i] = letterIndex
		letterMap[letterIndex] = i
	}

	return reflector{
		translationMap: letterMap,
	}, nil
}

func (r *reflector) translate(letter int) int {
	return r.translationMap[letter]
}
