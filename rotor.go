package main

import (
	"fmt"
)

type Rotor struct {
	name         string
	wiringMapIn  map[Char]Char // In = first pass through the rotors (from the plugboard to the reflector)
	wiringMapOut map[Char]Char // Out = second pass (from the reflector to the plugboard)
}

func NewRotor(name string, wiring string) (Rotor, error) {
	if sortString(wiring) != Alphabet {
		return Rotor{}, fmt.Errorf("invalid rotor wiring %s", wiring)
	}

	in := make(map[Char]Char, len(Alphabet))
	out := make(map[Char]Char, len(Alphabet))
	for i, letter := range wiring {
		in[Alphabet[i]] = Char(letter)
		out[Char(letter)] = Alphabet[i]
	}

	return Rotor{
		name:         name,
		wiringMapIn:  in,
		wiringMapOut: out,
	}, nil
}

func (r *Rotor) TranslateIn(letter Char) (Char, error) {
	return translateLatter(r.wiringMapIn, letter)
}

func (r *Rotor) TranslateOut(letter Char) (Char, error) {
	return translateLatter(r.wiringMapOut, letter)
}
