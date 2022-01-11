package enigma

import (
	"fmt"
	"strings"
)

// todo - improve/ expand this so it is clear what happens (now confusing because of implicit rotation)
type encryptionSequence struct {
	rotorPositions []int
	in             int
	out            int
	steps          []encryptionStep
}

type encryptionStep struct {
	title string
	out   int
}

func (es *encryptionSequence) start(rotors []rotor, letterToEncrypt int) {
	es.in = letterToEncrypt
	es.rotorPositions = make([]int, len(rotors))
	for i := range rotors {
		es.rotorPositions[i] = rotors[i].getPosition()
	}
}

func (es *encryptionSequence) addStep(title string, encodedLetter int) {
	step := encryptionStep{
		title: title,
		out:   encodedLetter,
	}
	es.steps = append(es.steps, step)
}

func (es *encryptionSequence) finish(encodedLetter int) {
	es.out = encodedLetter
}

func (es *encryptionSequence) getResult() char {
	return Alphabet.intToChar(es.out)
}

func (es *encryptionSequence) format() string {
	separator := "---------------------------------\n"
	positions := make([]string, len(es.rotorPositions))
	for i, position := range es.rotorPositions {
		positions[i] = string(Alphabet.intToChar(position))
	}
	result := fmt.Sprintf("INPUT: %s\n", string(Alphabet.intToChar(es.in)))
	result += fmt.Sprintf("rotor positions: %s\n", strings.Join(positions, ", "))
	for _, step := range es.steps {
		result += fmt.Sprintf("%s: %s\n", step.title, string(Alphabet.intToChar(step.out)))
	}
	result += fmt.Sprintf("OUTPUT: %s\n", string(Alphabet.intToChar(es.out)))

	return separator + result + separator
}
