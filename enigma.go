package main

import (
	"fmt"
)

type Enigma struct {
	rotors    []rotor
	reflector reflector
	plugboard plugboard
}

// todo - implement better configuration of various Enigma models
type rotorConfig struct {
	name          string
	wiring        string
	notchPosition char
}

func NewEnigma() (Enigma, error) {

	rotorConfigs := []rotorConfig{
		{
			name:          "I",
			wiring:        "EKMFLGDQVZNTOWYHXUSPAIBRCJ",
			notchPosition: 'Q',
		},
		{
			name:          "II",
			wiring:        "AJDKSIRUXBLHWTMCQGZNPYFVOE",
			notchPosition: 'E',
		},
		{
			name:          "III",
			wiring:        "BDFHJLCPRTXVZNYEIWGAKMUSQO",
			notchPosition: 'V',
		},
	}
	rotors := make([]rotor, 0, 3)
	for _, conf := range rotorConfigs {
		rotor, err := newRotor(conf.name, conf.wiring, conf.notchPosition)
		if err != nil {
			return Enigma{}, fmt.Errorf("failed to create rotor %s: %w", conf.name, err)
		}
		rotors = append(rotors, rotor)
	}

	reflector, err := newReflector("YRUHQSLDPXNGOKMIEBFZCWVJAT")
	if err != nil {
		return Enigma{}, fmt.Errorf("failed to create reflector: %w", err)
	}

	return Enigma{
		rotors:    rotors,
		reflector: reflector,
		plugboard: newPlugboard(),
	}, nil
}

// todo - add support for ring settings (only in some models)
func (e *Enigma) Configure(rotorPositions []char, plugConfig string) error {
	if len(rotorPositions) != len(e.rotors) {
		return fmt.Errorf("invalid rotor configurartion, different size than number of rotors")
	}
	for i, letter := range rotorPositions {
		if err := e.rotors[i].setPosition(letter); err != nil {
			return fmt.Errorf("failed to set position of rotor %s", e.rotors[i].name)
		}
	}

	if err := e.plugboard.setup(plugConfig); err != nil {
		return fmt.Errorf("failed to setup plugboard: %w", err)
	}

	return nil
}

// todo - write some basic tests (enabling refacoring / expansion)
func (e *Enigma) Encode(text string) (string, error) {
	result := make([]byte, len(text))
	for i, letter := range text {
		sequence, err := e.translate(char(letter))
		if err != nil {
			return "", fmt.Errorf("failed to encode letter [%s]: %w", string(letter), err)
		}
		result[i] = byte(sequence.getResult())
	}
	return string(result), nil
}

func (e *Enigma) translate(in char) (encryptionSequence, error) {
	letter, ok := Alphabet.charToInt(in)
	if !ok {
		return encryptionSequence{}, fmt.Errorf("unsupported letter")
	}

	// rotate the rotors first and start sequence
	e.rotate()
	sequence := encryptionSequence{}
	sequence.start(e.rotors, letter)

	// I. keyboard -> plugboard -> rotors
	letter = e.plugboard.translate(letter)
	sequence.addStep("plugboard", letter)

	// II. rotors -> reflector (reverse order of rotors, the letter goes from right to left)
	for i := len(e.rotors) - 1; i >= 0; i-- {
		letter = e.rotors[i].translateIn(letter)
		sequence.addStep(fmt.Sprintf("rotor %d", i+1), letter)
	}

	// III. reflector -> rotors
	letter = e.reflector.translate(letter)
	sequence.addStep("reflector", letter)

	// IV. rotors -> plugboard
	for i := 0; i < len(e.rotors); i++ {
		letter = e.rotors[i].translateOut(letter)
		sequence.addStep(fmt.Sprintf("rotor %d", i+1), letter)
	}

	// V. plugboard -> output
	letter = e.plugboard.translate(letter)
	sequence.addStep("plugboard", letter)

	sequence.finish(letter)
	return sequence, nil
}

func (e *Enigma) rotate() {
	rotateNext := true // first rotor from the right (the one with the highest index) always rotates
	for i := len(e.rotors) - 1; i >= 0; i-- {
		if rotateNext {
			rotateNext = e.rotors[i].rotate()
		}
	}
}
