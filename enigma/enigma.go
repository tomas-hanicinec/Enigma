package enigma

import (
	"fmt"
)

type Enigma struct {
	model
	plugboard        plugboard
	entryWheel       etw
	rotors           []rotor
	reflector        reflector
	reflectorIsWired bool
}

type RotorSlot int

const (
	Right  RotorSlot = 0
	Middle RotorSlot = 1
	Left   RotorSlot = 2
	Fourth RotorSlot = 3
)

func NewEnigma(model model) Enigma {
	e := Enigma{
		model:            model,
		plugboard:        newPlugboard(model.HasPlugboard()),
		entryWheel:       newEtw(model.GetEtwWiring()),
		rotors:           []rotor{},
		reflector:        newReflector(model.getDefaultReflectorType()),
		reflectorIsWired: false,
	}

	// select default rotors to all the slots
	if err := e.RotorsSelect(e.getDefaultRotorTypes()); err != nil {
		panic(fmt.Errorf("failed to select default rotors in %s model: %w", e.GetName(), err))
	}

	return e
}

// -------------------------------------- SETUP --------------------------------------

func (e *Enigma) RotorsSelect(rotorTypes map[RotorSlot]RotorType) error {
	if len(rotorTypes) != e.GetRotorCount() {
		return fmt.Errorf("%s model has %d rotors, but %d rotors selected", e.GetName(), e.GetRotorCount(), len(rotorTypes))
	}

	rotors := make([]rotor, e.GetRotorCount())
	isDuplicateSlot := map[RotorSlot]struct{}{}
	isDuplicateType := map[RotorType]struct{}{}
	for slot, rotorType := range rotorTypes {
		if !rotorType.exists() {
			return fmt.Errorf("invalid rotor type %s", rotorType)
		}
		// handle duplicates
		if _, ok := isDuplicateSlot[slot]; ok {
			return fmt.Errorf("cannot select rotor for slot number %d twice", slot)
		}
		if _, ok := isDuplicateType[rotorType]; ok {
			return fmt.Errorf("cannot select the rotor %s twice", rotorType)
		}
		// can only populate as many slots as the current model supports
		if int(slot) >= e.GetRotorCount() {
			return fmt.Errorf("cannot select rotor for slot number %d, Enigma model %s only has %d rotors", slot, e.GetName(), e.GetRotorCount())
		}
		// handle 4th rotor (only 2 rotors can fit to the 4th slot and vice-versa)
		isFourth := slot == Fourth
		if rotorType.canBeFourth() != isFourth {
			if isFourth {
				return fmt.Errorf("rotor %s cannot be selected as the fourth rotor", rotorType)
			} else {
				return fmt.Errorf("rotor %s can only be selected as the fourth rotor", rotorType)
			}
		}
		// all good, add the rotor
		rotors[slot] = newRotor(rotorType)
		isDuplicateSlot[slot] = struct{}{}
		isDuplicateType[rotorType] = struct{}{}
	}

	e.rotors = rotors
	return nil
}

func (e *Enigma) RotorSetWheel(slot RotorSlot, position byte) error {
	if err := e.validateRotorSlot(slot); err != nil {
		return err
	}
	return e.rotors[e.rotorSlotToIndex(slot)].setWheelPosition(position)
}

func (e *Enigma) RotorSetRing(slot RotorSlot, position int) error {
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
	err := e.reflector.setWiring(wiring)
	if err != nil {
		return err
	}
	e.reflectorIsWired = true
	return nil
}

func (e *Enigma) ReflectorSet(position byte) error {
	err := e.reflector.setPosition(position)
	if err != nil {
		return err
	}
	e.reflectorIsWired = false
	return nil
}

func (e *Enigma) PlugboardSet(plugConfig string) error {
	if !e.HasPlugboard() {
		return fmt.Errorf("%s model does not have plugboard, cannot set", e.GetName())
	}
	return e.plugboard.setup(plugConfig)
}

func (e *Enigma) validateRotorSlot(slot RotorSlot) error {
	if e.GetRotorCount() != 4 && slot == Fourth {
		return fmt.Errorf("%s model does not support 4th rotor", e.GetName())
	}
	index := e.rotorSlotToIndex(slot)
	if index < 0 || index >= e.GetRotorCount() {
		return fmt.Errorf("invalid rotor slot %d, %s model has %d rotors", slot, e.GetName(), e.GetRotorCount())
	}
	return nil
}

func (e *Enigma) rotorSlotToIndex(slot RotorSlot) int {
	return int(slot)
}

func (e *Enigma) rotorIndexToSlot(index int) RotorSlot {
	return RotorSlot(index)
}

func (e *Enigma) validateSettings() error {
	if e.reflector.reflectorType.isRewirable() && !e.reflectorIsWired {
		return fmt.Errorf("must specify reflector wiring")
	}

	return nil
}

// -------------------------------------- ENCODING --------------------------------------

func (e *Enigma) Encode(text string) (string, error) {
	result, _, err := e.doEncode(text)
	return result, err
}

func (e *Enigma) EncodeVerbose(text string) ([]EncryptionSequence, error) {
	_, sequences, err := e.doEncode(text)
	return sequences, err
}

func (e *Enigma) doEncode(text string) (string, []EncryptionSequence, error) {
	if err := e.validateSettings(); err != nil {
		return "", nil, fmt.Errorf("this Enigma machine is not fully configured yet: %w", err)
	}

	result := make([]byte, len(text))
	sequences := make([]EncryptionSequence, len(text))
	for i, letter := range text {
		sequence, err := e.translate(byte(letter))
		if err != nil {
			return "", nil, fmt.Errorf("failed to encode letter [%s]: %w", string(letter), err)
		}
		result[i] = sequence.GetResult()
		sequences[i] = sequence
	}
	return string(result), sequences, nil
}

func (e *Enigma) translate(in byte) (EncryptionSequence, error) {
	letter, ok := Alphabet.charToInt(in)
	if !ok {
		return EncryptionSequence{}, fmt.Errorf("unsupported letter")
	}

	// rotate the rotors first and start sequence
	e.rotate()
	sequence := EncryptionSequence{}
	sequence.start(e.rotors, letter)

	// I. plugboard -> ETW
	if e.HasPlugboard() {
		letter = e.plugboard.translate(letter)
		sequence.addStep("plugboard", letter)
	}

	// II. ETW -> rotors
	letter = e.entryWheel.translateIn(letter)
	sequence.addStep("etw", letter)

	// III. rotors -> reflector (reverse order of rotors, the letter goes from right to left)
	for i := 0; i < e.GetRotorCount(); i++ {
		letter = e.rotors[i].translateIn(letter)
		sequence.addStep(fmt.Sprintf("rotor %d", i+1), letter)
	}

	// IV. reflector -> rotors
	letter = e.reflector.translate(letter)
	sequence.addStep("reflector", letter)

	// V. rotors -> ETW
	for i := e.GetRotorCount() - 1; i >= 0; i-- {
		letter = e.rotors[i].translateOut(letter)
		sequence.addStep(fmt.Sprintf("rotor %d", i+1), letter)
	}

	// VI. ETW -> plugboard
	letter = e.entryWheel.translateOut(letter)
	sequence.addStep("etw", letter)

	// VII. plugboard -> output bulb
	if e.HasPlugboard() {
		letter = e.plugboard.translate(letter)
		sequence.addStep("plugboard", letter)
	}

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
