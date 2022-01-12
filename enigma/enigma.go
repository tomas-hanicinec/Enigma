package enigma

import (
	"fmt"
)

type Enigma struct {
	Model
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

func NewEnigma(model Model) (Enigma, error) {
	if !model.Exists() {
		return Enigma{}, fmt.Errorf("unsupported model %s", model)
	}
	e := Enigma{
		Model:            model,
		plugboard:        newPlugboard(model.HasPlugboard()),
		entryWheel:       newEtw(model.getEtwWiring()),
		rotors:           []rotor{},
		reflector:        newReflector(model.getDefaultReflectorType()),
		reflectorIsWired: false,
	}

	// select default rotors to all the slots
	if err := e.RotorsSelect(e.getDefaultRotorTypes()); err != nil {
		panic(fmt.Errorf("failed to select default rotors in %s model: %w", e.GetName(), err))
	}

	return e, nil
}

func NewEnigmaWithSetup(model Model, rotors map[RotorSlot]RotorConfig, reflector ReflectorConfig, plugboard string) (Enigma, error) {
	e, err := NewEnigma(model)
	if err != nil {
		return Enigma{}, err
	}

	if err := e.RotorsSetup(rotors); err != nil {
		return Enigma{}, fmt.Errorf("failed to setup rotors: %w", err)
	}

	if !reflector.isEmpty() {
		if err := e.ReflectorSetup(reflector); err != nil {
			return Enigma{}, fmt.Errorf("failed to setup reflector: %w", err)
		}
	}

	if plugboard != "" {
		if err := e.PlugboardSetup(plugboard); err != nil {
			return Enigma{}, fmt.Errorf("failed to setup plugboard: %w", err)
		}
	}

	return e, nil
}

// -------------------------------------- SETUP --------------------------------------

func (e *Enigma) RotorsSetup(config map[RotorSlot]RotorConfig) error {
	types := map[RotorSlot]RotorType{}
	for slot, rotorConfig := range config {
		types[slot] = rotorConfig.RotorType
	}
	rotors, err := e.getRotors(types)
	if err != nil {
		return fmt.Errorf("failed to set rotor types: %w", err)
	}

	for slot, rotorConfig := range config {
		if rotorConfig.WheelPosition != 0 {
			if err = rotors[e.rotorSlotToIndex(slot)].setWheelPosition(rotorConfig.WheelPosition); err != nil {
				return fmt.Errorf("failed to set wheel position for rotor %s: %w", rotorConfig.RotorType, err)
			}
		}
		if rotorConfig.RingPosition != 0 {
			if err = rotors[e.rotorSlotToIndex(slot)].setRingPosition(rotorConfig.RingPosition); err != nil {
				return fmt.Errorf("failed to set ring position for rotor %s: %w", rotorConfig.RotorType, err)
			}
		}
	}

	e.rotors = rotors
	return nil
}

func (e *Enigma) RotorsSelect(rotorTypes map[RotorSlot]RotorType) error {
	rotors, err := e.getRotors(rotorTypes)
	if err == nil {
		e.rotors = rotors
	}
	return err
}

func (e *Enigma) getRotors(rotorTypes map[RotorSlot]RotorType) ([]rotor, error) {
	if len(rotorTypes) != e.GetRotorCount() {
		return nil, fmt.Errorf("%s model has %d rotors, but %d rotors selected", e.GetName(), e.GetRotorCount(), len(rotorTypes))
	}

	rotors := make([]rotor, e.GetRotorCount())
	isDuplicateType := map[RotorType]struct{}{}
	for slot, rotorType := range rotorTypes {
		if !rotorType.exists() {
			return nil, fmt.Errorf("invalid rotor type %s", rotorType)
		}
		if !e.supportsRotorType(rotorType) {
			return nil, fmt.Errorf("%s model does not support rotor %s", e.GetName(), rotorType)
		}
		// handle duplicates
		if _, ok := isDuplicateType[rotorType]; ok {
			return nil, fmt.Errorf("cannot select the rotor %s twice", rotorType)
		}
		// can only populate as many slots as the current model supports
		if int(slot) >= e.GetRotorCount() {
			return nil, fmt.Errorf("cannot select rotor for slot number %d, %s model only has %d rotors", slot, e.GetName(), e.GetRotorCount())
		}
		// handle 4th rotor (only 2 rotors can fit to the 4th slot and vice-versa)
		isFourth := slot == Fourth
		if rotorType.CanBeFourth() != isFourth {
			if isFourth {
				return nil, fmt.Errorf("rotor %s cannot be selected as the fourth rotor", rotorType)
			} else {
				return nil, fmt.Errorf("rotor %s can only be selected as the fourth rotor", rotorType)
			}
		}
		// all good, add the rotor
		rotors[slot] = newRotor(rotorType)
		isDuplicateType[rotorType] = struct{}{}
	}

	return rotors, nil
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

func (e *Enigma) ReflectorSetup(config ReflectorConfig) error {
	ref, err := e.getReflector(config.ReflectorType)
	if err != nil {
		return fmt.Errorf("failed to select reflector: %w", err)
	}

	if config.WheelPosition != 0 {
		if err = ref.setWheelPosition(config.WheelPosition); err != nil {
			return fmt.Errorf("failed to set reflector position: %w", err)
		}
	}

	isWired := false
	if config.Wiring != "" {
		if err = ref.setWiring(config.Wiring); err != nil {
			return fmt.Errorf("failed to rewire reflector: %w", err)
		}
		isWired = true
	}

	e.reflector = ref
	e.reflectorIsWired = isWired
	return nil
}

func (e *Enigma) ReflectorSelect(reflectorType ReflectorType) error {
	ref, err := e.getReflector(reflectorType)
	if err == nil {
		e.reflector = ref
	}
	return err
}

func (e *Enigma) getReflector(reflectorType ReflectorType) (reflector, error) {
	if !e.supportsReflectorType(reflectorType) {
		return reflector{}, fmt.Errorf("%s model does not support reflector %s", e.GetName(), reflectorType)
	}

	return newReflector(reflectorType), nil
}

func (e *Enigma) ReflectorSetWheel(position byte) error {
	err := e.reflector.setWheelPosition(position)
	if err != nil {
		return err
	}
	e.reflectorIsWired = false
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

func (e *Enigma) PlugboardSetup(plugConfig string) error {
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

func (e *Enigma) supportsRotorType(rotorType RotorType) bool {
	for _, rot := range e.GetAvailableRotors() {
		if rot == rotorType {
			return true
		}
	}
	return false
}

func (e *Enigma) supportsReflectorType(reflectorType ReflectorType) bool {
	for _, ref := range e.GetAvailableReflectors() {
		if ref == reflectorType {
			return true
		}
	}
	return false
}

func (e *Enigma) rotorSlotToIndex(slot RotorSlot) int {
	return int(slot)
}

func (e *Enigma) rotorIndexToSlot(index int) RotorSlot {
	return RotorSlot(index)
}

func (e *Enigma) validateSettings() error {
	if e.reflector.reflectorType.IsRewirable() && !e.reflectorIsWired {
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
