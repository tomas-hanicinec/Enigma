package main

import (
	"fmt"
)

type rotor struct {
	name          string
	wiringMapIn   map[int]int // In = first pass through the rotors (from the plugboard to the reflector)
	wiringMapOut  map[int]int // Out = second pass (from the reflector to the plugboard)
	notchPosition int
	position      int
}

// todo - maybe validate only external inputs & let the "internal" errors (like invalid wiring) panic?
func newRotor(name string, wiring string, notchPosition char) (rotor, error) {
	if !Alphabet.isValidWiring(wiring) {
		return rotor{}, fmt.Errorf("invalid rotor wiring %s", wiring)
	}
	// todo - support multiple notches (later rotor models)
	notchIndex, ok := Alphabet.charToInt(notchPosition)
	if !ok {
		return rotor{}, fmt.Errorf("invalid notch position %s", string(notchPosition))
	}

	in := make(map[int]int, Alphabet.getSize())
	out := make(map[int]int, Alphabet.getSize())
	for i, letter := range wiring {
		letterIndex, ok := Alphabet.charToInt(char(letter))
		if !ok {
			return rotor{}, fmt.Errorf("unsupported wiring letter %s", string(letter)) // should not happen, we already checked the wiring validity
		}
		in[i] = letterIndex
		out[letterIndex] = i
	}

	return rotor{
		name:          name,
		wiringMapIn:   in,
		wiringMapOut:  out,
		notchPosition: notchIndex,
		position:      0, // start on the first position by default
	}, nil
}

func (r *rotor) setPosition(letter char) error {
	index, ok := Alphabet.charToInt(letter)
	if !ok {
		return fmt.Errorf("unsupported rotor position %s", string(letter))
	}
	r.position = index
	return nil
}

func (r *rotor) getPosition() int {
	return r.position
}

func (r *rotor) translateIn(letter int) int {
	letter = (letter + r.position) % Alphabet.getSize()                      // shift input according to the rotor current position
	letter = r.wiringMapIn[letter]                                           // translate
	letter = (letter + Alphabet.getSize() - r.position) % Alphabet.getSize() // shift output according to the rotor current position

	return letter
}

func (r *rotor) translateOut(letter int) int {
	letter = (letter + r.position) % Alphabet.getSize()                      // shift input according to the rotor current position
	letter = r.wiringMapOut[letter]                                          // translate
	letter = (letter + Alphabet.getSize() - r.position) % Alphabet.getSize() // shift output according to the rotor current position

	return letter
}

func (r *rotor) translate(translationMap map[int]int, letter int) (int, error) {
	translated, ok := translationMap[letter]
	if !ok {
		return 0, fmt.Errorf("invalid translation map, letter not there")
	}
	return translated, nil
}

func (r *rotor) rotate() bool {
	rotateNext := r.position == r.notchPosition
	r.position = (r.position + 1) % Alphabet.getSize()
	return rotateNext
}
