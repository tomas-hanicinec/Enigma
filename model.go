package enigma

type etwWiring string

const (
	etwAbcdef  = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	etwQwertz  = "QWERTZUIOASDFGHJKPYXCVBNML"
	etwTripitz = "KZROUQHYAIGBLWVSTDXFPNMCJE"
)

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

func GetSupportedModels() []Model {
	return []Model{
		SwissK,
		One,
		M3,
		M4,
		M4UKWD,
		SwissK,
		Tripitz,
	}
}

func (m Model) exists() bool {
	_, ok := models[m]
	return ok
}

func (m Model) GetName() string {
	return models[m].name
}

func (m Model) GetDescription() string {
	return models[m].description
}

func (m Model) GetYear() int {
	return models[m].yearIntroduced
}

func (m Model) getEtwWiring() etwWiring {
	return models[m].etw
}

func (m Model) HasPlugboard() bool {
	return models[m].hasPlugboard
}

func (m Model) GetAvailableRotorSlots() []RotorSlot {
	// it is important that the slots are ordered right to left as this is the order the current flows through
	if models[m].hasFourthRotor {
		return []RotorSlot{Right, Middle, Left, Fourth}
	}
	return []RotorSlot{Right, Middle, Left}
}

func (m Model) HasRotorSlot(slot RotorSlot) bool {
	for _, supportedSlot := range m.GetAvailableRotorSlots() {
		if supportedSlot == slot {
			return true
		}
	}
	return false
}

func (m Model) GetAvailableRotors(slot RotorSlot) []RotorType {
	allAvailable := models[m].rotors
	normal := make([]RotorType, 0, len(allAvailable))
	thin := make([]RotorType, 0)
	for _, rotorType := range allAvailable {
		if rotorType.IsThin() {
			thin = append(thin, rotorType)
		} else {
			normal = append(normal, rotorType)
		}
	}

	if slot == Fourth {
		return thin // fourth rotor slot only supports thin rotors
	}
	return normal
}

func (m Model) supportsRotorType(rotorType RotorType, slot RotorSlot) bool {
	for _, rot := range m.GetAvailableRotors(slot) {
		if rot == rotorType {
			return true
		}
	}
	return false
}

func (m Model) getDefaultRotorTypes() map[RotorSlot]RotorType {
	placedTypes := map[RotorType]struct{}{} // cannot use the same rotor type twice
	result := map[RotorSlot]RotorType{}
	for _, slot := range m.GetAvailableRotorSlots() {
		for _, availableType := range m.GetAvailableRotors(slot) {
			if _, ok := placedTypes[availableType]; !ok {
				result[slot] = availableType // use the first available unused rotor type
				placedTypes[availableType] = struct{}{}
				break
			}
		}
	}

	return result
}

func (m Model) GetAvailableReflectors() []ReflectorType {
	return models[m].reflectors
}

func (m Model) supportsReflectorType(reflectorType ReflectorType) bool {
	for _, ref := range m.GetAvailableReflectors() {
		if ref == reflectorType {
			return true
		}
	}
	return false
}

func (m Model) getDefaultReflectorType() ReflectorType {
	return models[m].reflectors[0]
}

type modelDefinition struct {
	name           string
	description    string
	yearIntroduced int
	hasPlugboard   bool
	hasFourthRotor bool
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
		hasFourthRotor: false,
		reflectors:     []ReflectorType{UkwK},
		rotors:         []RotorType{RotorIK, RotorIIK, RotorIIIK},
		etw:            etwQwertz,
	},
	One: {
		name:           "Enigma I",
		description:    "Military version of the commercial Enigma D used by the German army and air force. Plugboard added for greater cryptographic security and contrary to the commercial Enigma models, the reflector was fixed. Originally supplied with just three rotors, later in 1938 two more were added.",
		yearIntroduced: 1932,
		hasPlugboard:   true,
		hasFourthRotor: false,
		reflectors:     []ReflectorType{UkwA, UkwB},
		rotors:         []RotorType{RotorI, RotorII, RotorIII, RotorIV, RotorV},
		etw:            etwAbcdef,
	},
	M3: {
		name:           "Enigma M3",
		description:    "Variations of Enigma I developed for the German navy. Originally fully compatible with Enigma I with five rotors to choose from, but later three more rotors added exclusively for the navy",
		yearIntroduced: 1934,
		hasPlugboard:   true,
		hasFourthRotor: false,
		reflectors:     []ReflectorType{UkwA, UkwB, UkwC, UkwD},
		rotors:         []RotorType{RotorI, RotorII, RotorIII, RotorIV, RotorV, RotorVI, RotorVII, RotorVIII},
		etw:            etwAbcdef,
	},
	M4: {
		name:           "Enigma M4",
		description:    "Four-rotor version of Enigma M3, developed secretly by the German navy and later used mainly for U-boat traffic. Reflector was replaced by a special thin reflector and a fourth (thinner) rotor to increase the number of key combinations.",
		yearIntroduced: 1942,
		hasPlugboard:   true,
		hasFourthRotor: true,
		reflectors:     []ReflectorType{UkwBThin, UkwCThin},
		rotors:         []RotorType{RotorI, RotorII, RotorIII, RotorIV, RotorV, RotorVI, RotorVII, RotorVIII, RotorBeta, RotorGamma},
		etw:            etwAbcdef,
	},
	M4UKWD: {
		name:           "Enigma M4 with UKW-D",
		description:    "Field-rewirable reflector UKW-D was introduced by the air force for the Enigma M4. Could be plugged instead of the (thin) reflector and the fourth rotor. The UKW-D settings were typically only changed once in 10 days.",
		yearIntroduced: 1944,
		hasPlugboard:   true,
		hasFourthRotor: false,
		reflectors:     []ReflectorType{UkwD},
		rotors:         []RotorType{RotorI, RotorII, RotorIII, RotorIV, RotorV, RotorVI, RotorVII, RotorVIII},
		etw:            etwAbcdef,
	},
	SwissK: {
		name:           "Swiss-K",
		description:    "Built for the Swiss army before WWII. Based on commercial K model, but with different rotor wiring and an extra lamp panel",
		yearIntroduced: 1938,
		hasPlugboard:   false,
		hasFourthRotor: false,
		reflectors:     []ReflectorType{UkwK},
		rotors:         []RotorType{RotorISK, RotorIISK, RotorIIISK},
		etw:            etwQwertz,
	},
	Tripitz: {
		name:           "Enigma T (Tripitz)",
		description:    "Used for communication between German and Japanese navies. No plugboard, but specific ETW wiring and multiple turnover notches on rotors to increase security",
		yearIntroduced: 1942,
		hasPlugboard:   false,
		hasFourthRotor: false,
		reflectors:     []ReflectorType{UkwT},
		rotors:         []RotorType{RotorIT, RotorIIT, RotorIIIT, RotorIVT, RotorVT, RotorVIT, RotorVIIT, RotorVIIIT},
		etw:            etwTripitz,
	},
}
