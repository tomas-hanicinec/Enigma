package enigma

import (
	"fmt"
)

type etw struct {
	letterMapIn  map[int]int
	letterMapOut map[int]int
}

func newEtw(wiring etwWiring) etw {
	letterMapIn := map[int]int{}
	letterMapOut := map[int]int{}
	isDuplicate := map[int]struct{}{}
	for i := 0; i < Alphabet.getSize(); i++ {
		if i > len(wiring) {
			panic(fmt.Errorf("invalid ETW wiring, does not cover the whole alphabet"))
		}
		mappedIndex, ok := Alphabet.charToInt(wiring[i])
		if !ok {
			panic(fmt.Errorf("invalid ETW wiring pair %s->%s", string(Alphabet.intToChar(i)), string(wiring[i])))
		}
		if _, ok := isDuplicate[mappedIndex]; ok {
			panic(fmt.Errorf("invalid ETW wiring, letter %s is duplicate", string(wiring[i])))
		}
		letterMapIn[mappedIndex] = i
		letterMapOut[i] = mappedIndex
		isDuplicate[mappedIndex] = struct{}{}
	}

	return etw{
		letterMapIn:  letterMapIn,
		letterMapOut: letterMapOut,
	}
}

func (e *etw) translateIn(letter int) int {
	return e.letterMapIn[letter]
}

func (e *etw) translateOut(letter int) int {
	return e.letterMapOut[letter]
}
