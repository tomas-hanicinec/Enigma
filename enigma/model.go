package enigma

import "fmt"

type Model string

const (
	Commercial Model = "Commercial"
	One        Model = "I"
	M3         Model = "M3"
	M4         Model = "M4"
	M4UKWD     Model = "M4-UKW-D"
	SwissK     Model = "Swiss-K"
	Tripitz    Model = "Tripitz"
)

type etwWiring string

const (
	etwAbcdef  = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	etwQwertz  = "QWERTZUIOASDFGHJKPYXCVBNML"
	etwTripitz = "KZROUQHYAIGBLWVSTDXFPNMCJE"
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
	name           string
	description    string
	yearIntroduced int
	hasPlugboard   bool
	rotorCount     int
	reflectors     []ReflectorType
	rotors         []RotorType
	etw            etwWiring
}

var models = map[Model]modelDefinition{
	Commercial: {
		name:           "Commercial K",
		description:    "Based on Enigma C model and nearly identical to the D model introduced a year earlier, this was the most successful commercial Enigma model. Its core design with three swappable rotors and a movable reflector became the basis for all later Enigma models.",
		yearIntroduced: 1927,
		hasPlugboard:   false,
		rotorCount:     3,
		reflectors:     []ReflectorType{UKW_K},
		rotors:         []RotorType{Rotor_IK, Rotor_IIK, Rotor_IIIK},
		etw:            etwQwertz,
	},
	One: {
		name:           "Enigma I",
		description:    "Military version of the commercial Enigma D used by the German army and air force. Plugboard added for greater cryptographic security and contrary to the commercial Enigma models, the reflector was fixed. Originally supplied with just three rotors, later in 1938 two more were added.",
		yearIntroduced: 1932,
		hasPlugboard:   true,
		rotorCount:     3,
		reflectors:     []ReflectorType{UKW_A, UKW_B},
		rotors:         []RotorType{Rotor_I, Rotor_II, Rotor_III, Rotor_IV, Rotor_V},
		etw:            etwAbcdef,
	},
	M3: {
		name:           "Enigma M3",
		description:    "Variations of Enigma I developed for the German navy. Originally fully compatible with Enigma I with five rotors to choose from, but later three more rotors added exclusively for the navy",
		yearIntroduced: 1934,
		hasPlugboard:   true,
		rotorCount:     3,
		reflectors:     []ReflectorType{UKW_A, UKW_B, UKW_C, UKW_D},
		rotors:         []RotorType{Rotor_I, Rotor_II, Rotor_III, Rotor_IV, Rotor_V, Rotor_VI, Rotor_VII, Rotor_VIII},
		etw:            etwAbcdef,
	},
	M4: {
		name:           "Enigma M4",
		description:    "Four-rotor version of Enigma M3, developed secretly by the German navy and later used mainly for U-boat traffic. Reflector was replaced by a special thin reflector and a fourth (thinner) rotor to increase the number of key combinations.",
		yearIntroduced: 1942,
		hasPlugboard:   true,
		rotorCount:     4,
		reflectors:     []ReflectorType{UKW_BThin, UKW_CThin},
		rotors:         []RotorType{Rotor_I, Rotor_II, Rotor_III, Rotor_IV, Rotor_V, Rotor_VI, Rotor_VII, Rotor_VIII, Rotor_beta, Rotor_gamma},
		etw:            etwAbcdef,
	},
	M4UKWD: {
		name:           "Enigma M4 with UKW-D",
		description:    "Field-rewirable reflector UKW-D was introduced by the air force for the Enigma M4. Could be plugged instead of the (thin) reflector and the fourth rotor. The UKW-D settings were typically only changed once in 10 days.",
		yearIntroduced: 1944,
		hasPlugboard:   true,
		rotorCount:     3,
		reflectors:     []ReflectorType{UKW_D},
		rotors:         []RotorType{Rotor_I, Rotor_II, Rotor_III, Rotor_IV, Rotor_V, Rotor_VI, Rotor_VII, Rotor_VIII},
		etw:            etwAbcdef,
	},
	SwissK: {
		name:           "Swiss-K",
		description:    "Built for the Swiss army before WWII. Based on commercial K model, but with different rotor wiring and an extra lamp panel",
		yearIntroduced: 1938,
		hasPlugboard:   false,
		rotorCount:     3,
		reflectors:     []ReflectorType{UKW_K},
		rotors:         []RotorType{Rotor_ISK, Rotor_IISK, Rotor_IIISK},
		etw:            etwQwertz,
	},
	Tripitz: {
		name:           "Enigma T (Tripitz)",
		description:    "Used for communication between German and Japanese navies. No plugboard, but specific ETW wiring and multiple turnover notches on rotors to increase security",
		yearIntroduced: 1942,
		hasPlugboard:   false,
		rotorCount:     3,
		reflectors:     []ReflectorType{UKW_T},
		rotors:         []RotorType{Rotor_IT, Rotor_IIT, Rotor_IIIT, Rotor_IVT, Rotor_VT, Rotor_VIT, Rotor_VIIT, Rotor_VIIIT},
		etw:            etwTripitz,
	},
}
