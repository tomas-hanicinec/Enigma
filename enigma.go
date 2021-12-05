package main

import "fmt"

type Char = byte

const Alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

type Enigma struct {
	rotors    []Rotor
	reflector Reflector
	plugboard Plugboard
}

func NewEnigma() (Enigma, error) {

	rotorWiring := map[string]string{
		"I":   "JGDQOXUSCAMIFRVTPNEWKBLZYH",
		"II":  "NTZPSFBOKMWRCJDIVLAEYUXHGQ",
		"III": "JVIUBHTCDYAKEQZPOSGXNRMWFL",
	}
	rotors := make([]Rotor, 0, 3)
	for name, wiring := range rotorWiring {
		rotor, err := NewRotor(name, wiring)
		if err != nil {
			return Enigma{}, fmt.Errorf("failed to create rotor %s: %w", name, err)
		}
		rotors = append(rotors, rotor)
	}

	/*
		Reflector A	EJMZALYXVBWFCRQUONTSPIKHGD
		Reflector B	YRUHQSLDPXNGOKMIEBFZCWVJAT
		Reflector C	FVPJIAOYEDRZXWGCTKUQSBNMHL
	*/
	reflector, err := NewReflector("EJMZALYXVBWFCRQUONTSPIKHGD")
	if err != nil {
		return Enigma{}, fmt.Errorf("failed to create reflector: %w", err)
	}

	return Enigma{
		rotors:    rotors,
		reflector: reflector,
		plugboard: Plugboard{},
	}, nil
}

func (e *Enigma) Configure(rotorPositions string, plugConfig string) error {
	//@todo - setup the rotors

	if err := e.plugboard.Setup(plugConfig); err != nil {
		return fmt.Errorf("failed to setup plugboard: %w", err)
	}

	return nil
}

func (e *Enigma) Encode(text string) (string, error) {
	result := make([]Char, len(text))
	err := error(nil)
	for i, letter := range text {
		result[i], err = e.encodeLetter(Char(letter))
		if err != nil {
			return "", fmt.Errorf("failed to encode string: %w", err)
		}
	}
	return string(result), nil
}

func (e *Enigma) encodeLetter(letter Char) (Char, error) {
	// I. keyboard -> plugboard -> rotors
	letter, err := e.plugboard.Translate(letter)
	if err != nil {
		return 0, fmt.Errorf("failed to translate incomming letter [%s] in plugboard: %w", string(letter), err)
	}

	// II. rotors -> reflector
	for i := 0; i < len(e.rotors); i++ {
		letter, err = e.rotors[i].TranslateIn(letter)
		if err != nil {
			return 0, fmt.Errorf("failed to translate incomming letter [%s] in rotor %s", string(letter), e.rotors[i].name)
		}
	}

	// III. reflector -> rotors
	letter, err = e.reflector.Translate(letter)
	if err != nil {
		return 0, fmt.Errorf("failed to translate letter [%s] in reflector: %w", string(letter), err)
	}

	// IV. rotors -> plugboard
	for i := len(e.rotors) - 1; i >= 0; i-- {
		letter, err = e.rotors[i].TranslateOut(letter)
		if err != nil {
			return 0, fmt.Errorf("failed to translate outgoing letter [%s] in rotor %s", string(letter), e.rotors[i].name)
		}
	}

	// V. plugboard -> output
	letter, err = e.plugboard.Translate(letter)
	if err != nil {
		return 0, fmt.Errorf("failed to translate outgoing letter [%s] in plugboard: %w", string(letter), err)
	}

	// VI. rotate the rotors for the next letter
	e.rotate()

	return letter, nil
}

func (e *Enigma) rotate() {
	//@todo - change the rotors position
}
