package enigma

import (
	"fmt"
)

type Enigma struct {
	Model
	plugboard  plugboard
	entryWheel etw
	rotors     []rotor
	reflector  reflector
}

type RotorSlot int

const (
	Right  RotorSlot = 0
	Middle RotorSlot = 1
	Left   RotorSlot = 2
	Fourth RotorSlot = 3
)

func NewEnigma(model Model) (Enigma, error) {
	if !model.exists() {
		return Enigma{}, fmt.Errorf("unsupported model %s", model)
	}
	e := Enigma{
		Model:      model,
		plugboard:  newPlugboard(model.HasPlugboard()),
		entryWheel: newEtw(model.getEtwWiring()),
		rotors:     []rotor{},
		reflector:  newReflector(model.getDefaultReflectorType()),
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

	if len(rotors) > 0 {
		if err := e.RotorsSetup(rotors); err != nil {
			return Enigma{}, fmt.Errorf("failed to setup rotors: %w", err)
		}
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

func (e *Enigma) GetReflectorType() ReflectorType {
	return e.reflector.reflectorType
}

// -------------------------------------- SETUP --------------------------------------

func (e *Enigma) RotorsSetup(config map[RotorSlot]RotorConfig) error {
	types := map[RotorSlot]RotorType{}
	for i, rotor := range e.rotors {
		types[e.rotorIndexToSlot(i)] = rotor.rotorType // fill with current values
	}
	for slot, rotorConfig := range config {
		if !e.HasRotorSlot(slot) {
			return fmt.Errorf("unsupported rotor slot %d", slot)
		}
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
	availableSlots := e.GetAvailableRotorSlots()
	if len(rotorTypes) != len(availableSlots) {
		return nil, fmt.Errorf("%s model has %d rotors, but %d rotors selected", e.GetName(), len(availableSlots), len(rotorTypes))
	}

	rotors := make([]rotor, len(availableSlots))
	isDuplicateType := map[RotorType]struct{}{}
	for slot, rotorType := range rotorTypes {
		// can only populate slots supported by the current model
		if !e.HasRotorSlot(slot) {
			return nil, fmt.Errorf("unsupported rotor slot %d", slot)
		}
		// can only place supported rotor to the slot
		if !e.supportsRotorType(rotorType, slot) {
			return nil, fmt.Errorf("%s model does not support rotor %s in slot %d", e.GetName(), rotorType, slot)
		}
		// handle duplicates
		if _, ok := isDuplicateType[rotorType]; ok {
			return nil, fmt.Errorf("cannot select the rotor %s twice", rotorType)
		}

		// all good, add the rotor
		rotors[slot] = newRotor(rotorType)
		isDuplicateType[rotorType] = struct{}{}
	}

	return rotors, nil
}

func (e *Enigma) RotorSetWheel(slot RotorSlot, position byte) error {
	if !e.HasRotorSlot(slot) {
		return fmt.Errorf("unsupported rotor slot %d", slot)
	}
	return e.rotors[e.rotorSlotToIndex(slot)].setWheelPosition(position)
}

func (e *Enigma) RotorSetRing(slot RotorSlot, position int) error {
	if !e.HasRotorSlot(slot) {
		return fmt.Errorf("unsupported rotor slot %d", slot)
	}
	return e.rotors[e.rotorSlotToIndex(slot)].setRingPosition(position)
}

func (e *Enigma) RotorsReset() {
	for slot := range e.rotors {
		e.rotors[slot].reset()
	}
}

func (e *Enigma) ReflectorSetup(config ReflectorConfig) error {
	reflectorType := config.ReflectorType
	if reflectorType == "" {
		reflectorType = e.reflector.reflectorType // use current if not specified
	}
	ref, err := e.getReflector(reflectorType)
	if err != nil {
		return fmt.Errorf("failed to select reflector: %w", err)
	}

	if config.WheelPosition != 0 {
		if err = ref.setWheelPosition(config.WheelPosition); err != nil {
			return fmt.Errorf("failed to set reflector position: %w", err)
		}
	}

	if config.Wiring != "" {
		if err = ref.setWiring(config.Wiring); err != nil {
			return fmt.Errorf("failed to rewire reflector: %w", err)
		}
	}

	e.reflector = ref
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
	return nil
}

func (e *Enigma) ReflectorRewire(wiring string) error {
	err := e.reflector.setWiring(wiring)
	if err != nil {
		return err
	}
	return nil
}

func (e *Enigma) PlugboardSetup(plugConfig string) error {
	if !e.HasPlugboard() {
		return fmt.Errorf("%s model does not have a plugboard", e.GetName())
	}
	return e.plugboard.setup(plugConfig)
}

func (e *Enigma) rotorSlotToIndex(slot RotorSlot) int {
	return int(slot)
}

func (e *Enigma) rotorIndexToSlot(index int) RotorSlot {
	return RotorSlot(index)
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
	result := make([]byte, len(text))
	sequences := make([]EncryptionSequence, len(text))
	for i, letter := range text {
		sequence, err := e.translate(byte(letter))
		if err != nil {
			return "", nil, fmt.Errorf("failed to encode letter \"%s\": %w", string(letter), err)
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
	slots := e.GetAvailableRotorSlots()
	for _, slot := range slots {
		slotIndex := e.rotorSlotToIndex(slot)
		letter = e.rotors[slotIndex].translateIn(letter)
		sequence.addStep(fmt.Sprintf("rotor %d", slotIndex+1), letter)
	}

	// IV. reflector -> rotors
	letter = e.reflector.translate(letter)
	sequence.addStep("reflector", letter)

	// V. rotors -> ETW
	for i := len(slots) - 1; i >= 0; i-- {
		slotIndex := e.rotorSlotToIndex(slots[i])
		letter = e.rotors[slotIndex].translateOut(letter)
		sequence.addStep(fmt.Sprintf("rotor %d", slotIndex+1), letter)
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
