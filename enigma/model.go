package enigma

import "fmt"

type Model string

const (
	SwissK Model = "Swiss-K"
	One          = "I"
	M3           = "M3"
	M4           = "M4"
	M4UKWD       = "M4-UKW-D"
)

type etwWiring string

const (
	etwStandard = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	etwSwiss    = "QWERTZUIOASDFGHJKPYXCVBNML"
)

func GetModels() []Model {
	return []Model{
		SwissK,
		One,
		M3,
		M4,
		M4UKWD,
	}
}

func (m Model) Exists() bool {
	_, ok := models[m]
	return ok
}

func (m Model) GetName() string {
	return models[m].name
}

func (m Model) GetDescription() string {
	return models[m].description
}

func (m Model) getEtwWiring() etwWiring {
	return models[m].etw
}

func (m Model) HasPlugboard() bool {
	return models[m].hasPlugboard
}

func (m Model) GetRotorCount() int {
	return models[m].rotorCount
}

func (m Model) GetAvailableRotors() []RotorType {
	return models[m].rotors
}

func (m Model) getDefaultRotorTypes() map[RotorSlot]RotorType {
	result := make(map[RotorSlot]RotorType)

	// basic three rotors
	slots := []RotorSlot{Right, Middle, Left}
	i := 0
	for _, rotor := range m.GetAvailableRotors() {
		// pick first 3 rotors that cannot be 4th
		if !rotor.CanBeFourth() {
			result[slots[i]] = rotor
			i++
		}
		if i >= len(slots) {
			break
		}
	}

	// fourth rotor
	if m.GetRotorCount() == 4 {
		// pick first rotor that can be 4th
		for _, rotor := range m.GetAvailableRotors() {
			if rotor.CanBeFourth() {
				result[Fourth] = rotor
				break
			}
		}
	}

	if len(result) != m.GetRotorCount() {
		panic(fmt.Errorf("error getting default rotors for %s model", m.GetName()))
	}

	return result
}

func (m Model) GetAvailableReflectors() []ReflectorType {
	return models[m].reflectors
}

func (m Model) getDefaultReflectorType() ReflectorType {
	return models[m].reflectors[0]
}

type modelDefinition struct {
	name         string
	description  string
	hasPlugboard bool
	rotorCount   int
	reflectors   []ReflectorType
	rotors       []RotorType
	etw          etwWiring
}

var models = map[Model]modelDefinition{
	// todo - better descriptions
	// todo - more models
	SwissK: {
		name:         "Swiss-K",
		description:  "1939, based on commercial D, movable reflector, different entry-wheel (QWERTZ)",
		hasPlugboard: false,
		rotorCount:   3,
		reflectors:   []ReflectorType{UKW_K},
		rotors:       []RotorType{Rotor_IK, Rotor_IIK, Rotor_IIIK},
		etw:          etwSwiss,
	},
	One: {
		name:         "Enigma I",
		description:  "1930–1938, based on commercial D, fixed reflector, added plugboard, notches on the movable rings, since 1935 air force used it too",
		hasPlugboard: true,
		rotorCount:   3,
		reflectors:   []ReflectorType{UKW_A, UKW_B},
		rotors:       []RotorType{Rotor_I, Rotor_II, Rotor_III},
		etw:          etwStandard,
	},
	M3: {
		name:         "Enigma M3",
		description:  "1934, navy version of Enigma I, added more rotors to choose from (1934 - 5, 1938 - 7, 1939 - 8)",
		hasPlugboard: true,
		rotorCount:   3,
		reflectors:   []ReflectorType{UKW_A, UKW_B, UKW_C, UKW_D},
		rotors:       []RotorType{Rotor_I, Rotor_II, Rotor_III, Rotor_IV, Rotor_V, Rotor_VI, Rotor_VII, Rotor_VIII},
		etw:          etwStandard,
	},
	M4: {
		name:         "Enigma M4",
		description:  "1942, reflector replaced by a thinner reflector and a fourth rotor, used for U-boat traffic",
		hasPlugboard: true,
		rotorCount:   4,
		reflectors:   []ReflectorType{UKW_BThin, UKW_CThin},
		rotors:       []RotorType{Rotor_I, Rotor_II, Rotor_III, Rotor_IV, Rotor_V, Rotor_VI, Rotor_VII, Rotor_VIII, Rotor_beta, Rotor_gamma},
		etw:          etwStandard,
	},
	M4UKWD: {
		name:         "Enigma M4 with UKW-D",
		description:  "1944, added reconfigurable reflector",
		hasPlugboard: true,
		rotorCount:   3,
		reflectors:   []ReflectorType{UKW_D},
		rotors:       []RotorType{Rotor_I, Rotor_II, Rotor_III, Rotor_IV, Rotor_V, Rotor_VI, Rotor_VII, Rotor_VIII},
		etw:          etwStandard,
	},
}
