package enigma

import (
	"fmt"
)

type Enigma struct {
	model
	entry     etw
	plugboard plugboard
	rotors    []rotor
	reflector reflector
}

type rotorSlot int

const (
	Left   rotorSlot = 1
	Middle           = 2
	Right            = 3
	Fourth           = 0
)

func NewEnigma(model model) Enigma {
	e := Enigma{
		model:     model,
		entry:     newEtw(model.GetEtwWiring()),
		plugboard: newPlugboard(model.HasPlugboard()),
		rotors:    []rotor{},
		reflector: newReflector(model.getDefaultReflectorType()),
	}

	// select default rotors to all the slots
	if err := e.RotorsSelect(e.getDefaultRotorTypes()); err != nil {
		panic(fmt.Errorf("failed to select default rotors in %s model: %w", e.GetName(), err))
	}

	//todo - set defaults for rotor / reflector positions & wiring

	return e
}

// -------------------------------------- SETUP --------------------------------------

// todo - do not use slice, but map[position]rotorType
func (e *Enigma) RotorsSelect(rotorTypes []RotorType) error {
	if len(rotorTypes) != e.GetRotorCount() {
		return fmt.Errorf("%s model has %d rotors, but %d rotors selected", e.GetName(), e.GetRotorCount(), len(rotorTypes))
	}

	rotors := make([]rotor, 0, e.GetRotorCount())
	for i, rotorType := range rotorTypes {
		if !rotorType.exists() {
			return fmt.Errorf("invalid rotor type %s", rotorType)
		}
		// handle duplicates
		for j, duplicateRotorType := range rotorTypes {
			if rotorType == duplicateRotorType && i != j {
				return fmt.Errorf("cannot select the rotor %s twice", rotorType)
			}
		}
		// handle 4th rotor (only 2 rotors can fit to the 4th slot and vice-versa)
		isFourth := e.GetRotorCount() == 4 && i == 0
		if rotorType.canBeFourth() != isFourth {
			if isFourth {
				return fmt.Errorf("rotor %s cannot be selected as the fourth rotor", rotorType)
			} else {
				return fmt.Errorf("rotor %s can only be selected as the fourth rotor", rotorType)
			}
		}
		// all good, add the rotor
		rotors = append(rotors, newRotor(rotorType))
	}

	e.rotors = rotors
	return nil
}

func (e *Enigma) RotorSetWheel(slot rotorSlot, position byte) error {
	if err := e.validateRotorSlot(slot); err != nil {
		return err
	}
	return e.rotors[e.rotorSlotToIndex(slot)].setWheelPosition(position)
}

func (e *Enigma) RotorSetRing(slot rotorSlot, position int) error {
	if err := e.validateRotorSlot(slot); err != nil {
		return err
	}
	return e.rotors[e.rotorSlotToIndex(slot)].setRingPosition(position)
}

func (e *Enigma) ReflectorSelect(reflectorType reflectorType) error {
	// check if the type is supported by the model
	supported := false
	for _, ref := range e.GetAvailableReflectors() {
		if ref == reflectorType {
			supported = true
			break
		}
	}
	if !supported {
		return fmt.Errorf("%s model does not support reflector %s", e.GetName(), reflectorType)
	}

	e.reflector = newReflector(reflectorType)
	return nil
}

func (e *Enigma) ReflectorRewire(wiring string) error {
	return e.reflector.setWiring(wiring)
}

func (e *Enigma) ReflectorSet(position byte) error {
	return e.reflector.setPosition(position)
}

func (e *Enigma) PlugboardSet(plugConfig string) error {
	if !e.HasPlugboard() {
		return fmt.Errorf("%s model does not have plugboard, cannot set", e.GetName())
	}
	return e.plugboard.setup(plugConfig)
}

func (e *Enigma) validateRotorSlot(slot rotorSlot) error {
	if e.GetRotorCount() != 4 && slot == Fourth {
		return fmt.Errorf("%s model does not support 4th rotor", e.GetName())
	}
	index := e.rotorSlotToIndex(slot)
	if index < 0 || index >= e.GetRotorCount() {
		return fmt.Errorf("invalid rotor slot %d, %s model has %d rotors", slot, e.GetName(), e.GetRotorCount())
	}
	return nil
}

func (e *Enigma) rotorSlotToIndex(slot rotorSlot) int {
	if e.GetRotorCount() == 4 {
		return int(slot)
	}
	return int(slot) - 1
}

func (e *Enigma) rotorIndexToSlot(index int) rotorSlot {
	if e.GetRotorCount() == 4 {
		return rotorSlot(index)
	}
	return rotorSlot(index + 1)
}

// -------------------------------------- ENCODING --------------------------------------

func (e *Enigma) Encode(text string) (string, error) {
	result := make([]byte, len(text))
	for i, letter := range text {
		sequence, err := e.translate(byte(letter))
		if err != nil {
			return "", fmt.Errorf("failed to encode letter [%s]: %w", string(letter), err)
		}
		result[i] = byte(sequence.getResult())
	}
	return string(result), nil
}

func (e *Enigma) translate(in byte) (encryptionSequence, error) {
	letter, ok := Alphabet.charToInt(in)
	if !ok {
		return encryptionSequence{}, fmt.Errorf("unsupported letter")
	}

	// rotate the rotors first and start sequence
	e.rotate()
	sequence := encryptionSequence{}
	sequence.start(e.rotors, letter)

	// I. keyboard -> ETW
	letter = e.entry.translateIn(letter)
	sequence.addStep("etw", letter)

	// II. ETW -> plugboard -> rotors
	letter = e.plugboard.translate(letter)
	sequence.addStep("plugboard", letter)

	// III. rotors -> reflector (reverse order of rotors, the letter goes from right to left)
	for i := len(e.rotors) - 1; i >= 0; i-- {
		letter = e.rotors[i].translateIn(letter)
		sequence.addStep(fmt.Sprintf("rotor %d", i+1), letter)
	}

	// IV. reflector -> rotors
	letter = e.reflector.translate(letter)
	sequence.addStep("reflector", letter)

	// V. rotors -> plugboard
	for i := 0; i < len(e.rotors); i++ {
		letter = e.rotors[i].translateOut(letter)
		sequence.addStep(fmt.Sprintf("rotor %d", i+1), letter)
	}

	// VI. plugboard -> ETW
	letter = e.plugboard.translate(letter)
	sequence.addStep("plugboard", letter)

	// VII. ETW -> output bulb
	letter = e.entry.translateOut(letter)
	sequence.addStep("etw", letter)

	sequence.finish(letter)
	return sequence, nil
}

func (e *Enigma) rotate() {
	// determine which rotors should be rotated in this step
	rotateMiddle := e.rotors[e.rotorSlotToIndex(Right)].shouldRotateNext()
	rotateLeft := e.rotors[e.rotorSlotToIndex(Middle)].shouldRotateNext()

	e.rotors[e.rotorSlotToIndex(Right)].rotate() // always rotate the right rotor
	if rotateMiddle {
		e.rotors[e.rotorSlotToIndex(Middle)].rotate()
	}
	if rotateLeft {
		// double-stepping - middle rotor rotates again if left rotor rotates
		e.rotors[e.rotorSlotToIndex(Middle)].rotate()
		e.rotors[e.rotorSlotToIndex(Left)].rotate()
	}
}
