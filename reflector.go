package main

import (
	"fmt"
)

type Reflector struct {
	letterMap map[Char]Char
}

func NewReflector(config string) (Reflector, error) {
	if sortString(config) != Alphabet {
		return Reflector{}, fmt.Errorf("invalid reflector config %s", config)
	}

	letterMap := make(map[Char]Char, len(Alphabet))
	for i, letter := range config {
		letterMap[Alphabet[i]] = Char(letter)
		letterMap[Char(letter)] = Alphabet[i]
	}

	return Reflector{
		letterMap: letterMap,
	}, nil
}

func (r *Reflector) Translate(letter Char) (Char, error) {
	return translateLatter(r.letterMap, letter)
}
