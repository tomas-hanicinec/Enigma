package enigma

import (
	"fmt"
)

type RotorConfig = struct {
	RotorType     RotorType
	WheelPosition byte
	RingPosition  int
}

type rotor struct {
	rotorType            RotorType
	wiringMapIn          map[int]int // In = first pass through the rotors (from the plugboard to the reflector)
	wiringMapOut         map[int]int // Out = second pass (from the reflector to the plugboard)
	notchPositions       []int
	initialWheelPosition byte // necessary for rotor reset
	wheelPosition        int
	ringPosition         int
}

func newRotor(rotorType RotorType) rotor {
	if !rotorType.exists() {
		panic(fmt.Errorf("invalid rotor type"))
	}
	wiring := rotorType.getWiring()
	if !Alphabet.isValidWiring(wiring) {
		panic(fmt.Errorf("invalid rotor wiring %s", wiring))
	}
	notchPositions := make([]int, len(rotorType.getNotchPositions()))
	for i, notchPositionByte := range rotorType.getNotchPositions() {
		notchPositionInt, ok := Alphabet.charToInt(notchPositionByte)
		if !ok {
			panic(fmt.Errorf("invalid notch position %s", string(notchPositionByte)))
		}
		notchPositions[i] = notchPositionInt
	}

	in := make(map[int]int, Alphabet.getSize())
	out := make(map[int]int, Alphabet.getSize())
	for i, letter := range wiring {
		letterIndex, ok := Alphabet.charToInt(byte(letter))
		if !ok {
			panic(fmt.Errorf("unsupported wiring letter %s", string(letter))) // should not happen, we already checked the wiring validity
		}
		in[i] = letterIndex
		out[letterIndex] = i
	}

	return rotor{
		rotorType:            rotorType,
		wiringMapIn:          in,
		wiringMapOut:         out,
		notchPositions:       notchPositions,
		initialWheelPosition: Alphabet.intToChar(0),
		wheelPosition:        0, // start on the first position by default
		ringPosition:         1,
	}
}

func (r *rotor) setWheelPosition(letter byte) error {
	index, ok := Alphabet.charToInt(letter)
	if !ok {
		return fmt.Errorf("unsupported rotor wheel position \"%s\"", string(letter))
	}
	r.wheelPosition = index
	r.initialWheelPosition = letter
	return nil
}

func (r *rotor) getWheelPosition() int {
	return r.wheelPosition
}

func (r *rotor) setRingPosition(position int) error {
	if position < 1 || position > Alphabet.getSize() {
		return fmt.Errorf("invalid ring position %d, must be a number between 1 and %d", position, Alphabet.getSize())
	}
	r.ringPosition = position
	return nil
}

func (r *rotor) reset() {
	if err := r.setWheelPosition(r.initialWheelPosition); err != nil {
		panic(fmt.Errorf("failed to reset rotor %s: %w", r.rotorType, err))
	}
}

func (r *rotor) translateIn(input int) int {
	return r.translate(input, r.wiringMapIn)
}

func (r *rotor) translateOut(input int) int {
	return r.translate(input, r.wiringMapOut)
}

func (r *rotor) translate(input int, translateMap map[int]int) int {
	shiftSize := r.wheelPosition - r.ringPosition + 1
	rotatedInput := shift(input, shiftSize)     // shift according to the wheel and ring rotation
	rotatedOutput := translateMap[rotatedInput] // translate
	return shift(rotatedOutput, -shiftSize)     // shift back
}

func (r *rotor) rotate() {
	r.wheelPosition = (r.wheelPosition + 1) % Alphabet.getSize()
}

func (r *rotor) shouldRotateNext() bool {
	for _, notchPosition := range r.notchPositions {
		if r.wheelPosition == notchPosition {
			return true // double-stepping - we are about to cross a notch in the next step, next rotor should be rotated too then
		}
	}
	return false
}
