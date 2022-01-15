package enigma

import (
	"fmt"
	"strings"
)

// EncryptionSequence contains detailed information about encryption process of a single letter,
// can be used for debugging via Enigma.EncodeVerbose
type EncryptionSequence struct {
	rotorPositions []int
	in             int
	out            int
	steps          []encryptionStep
}

type encryptionStep struct {
	title string
	out   int
}

func (es *EncryptionSequence) start(rotors []rotor, letterToEncrypt int) {
	es.in = letterToEncrypt
	es.rotorPositions = make([]int, len(rotors))
	for i := range rotors {
		es.rotorPositions[i] = rotors[i].getWheelPosition()
	}
}

func (es *EncryptionSequence) addStep(title string, encodedLetter int) {
	step := encryptionStep{
		title: title,
		out:   encodedLetter,
	}
	es.steps = append(es.steps, step)
}

func (es *EncryptionSequence) finish(encodedLetter int) {
	es.out = encodedLetter
}

// GetResult returns the final encrypted letter
func (es *EncryptionSequence) GetResult() byte {
	return Alphabet.intToChar(es.out)
}

// Format returns human-readable string representation of the sequence
func (es *EncryptionSequence) Format() string {
	separator := "---------------------------------\n"
	positions := make([]string, len(es.rotorPositions))
	for i, position := range es.rotorPositions {
		positions[len(es.rotorPositions)-i-1] = string(Alphabet.intToChar(position))
	}
	result := fmt.Sprintf("INPUT: %s\n", string(Alphabet.intToChar(es.in)))
	result += fmt.Sprintf("rotor wheel positions: %s\n", strings.Join(positions, ", "))
	for _, step := range es.steps {
		result += fmt.Sprintf("%s: %s\n", step.title, string(Alphabet.intToChar(step.out)))
	}
	result += fmt.Sprintf("OUTPUT: %s\n", string(Alphabet.intToChar(es.out)))

	return separator + result + separator
}
