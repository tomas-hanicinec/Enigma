package enigma

type etwWiring string

const (
	etwAbcdef  = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	etwQwertz  = "QWERTZUIOASDFGHJKPYXCVBNML"
	etwTripitz = "KZROUQHYAIGBLWVSTDXFPNMCJE"
)

// Model specifies the Enigma machine model
type Model string

// all supported Enigma models
const (
	Commercial Model = "Commercial"
	One        Model = "I"
	M3         Model = "M3"
	M4         Model = "M4"
	M4UKWD     Model = "M4-UKW-D"
	SwissK     Model = "Swiss-K"
	Tripitz    Model = "Tripitz"
)

// GetSupportedModels returns all the supported Enigma models
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

// GetName returns full name of this model
func (m Model) GetName() string {
	return models[m].name
}

// GetDescription returns brief description of this model
func (m Model) GetDescription() string {
	return models[m].description
}

// GetYear returns the year in which this model was first introduced
func (m Model) GetYear() int {
	return models[m].yearIntroduced
}

func (m Model) getEtwWiring() etwWiring {
	return models[m].etw
}

// HasPlugboard shows if this model has a rewirable plugboard (not all Enigma models had that)
func (m Model) HasPlugboard() bool {
	return models[m].hasPlugboard
}

// GetAvailableRotorSlots returns all the rotor slots available in this model
func (m Model) GetAvailableRotorSlots() []RotorSlot {
	// it is important that the slots are ordered right to left as this is the order the current flows through
	if models[m].hasFourthRotor {
		return []RotorSlot{Right, Middle, Left, Fourth}
	}
	return []RotorSlot{Right, Middle, Left}
}

// HasRotorSlot determines if the given rotor slot was present in this Enigma model
func (m Model) HasRotorSlot(slot RotorSlot) bool {
	for _, supportedSlot := range m.GetAvailableRotorSlots() {
		if supportedSlot == slot {
			return true
		}
	}
	return false
}

// GetAvailableRotorModels returns the supported rotor set for this Enigma model
func (m Model) GetAvailableRotorModels(slot RotorSlot) []RotorModel {
	allAvailable := models[m].rotors
	normal := make([]RotorModel, 0, len(allAvailable))
	thin := make([]RotorModel, 0)
	for _, rotorModel := range allAvailable {
		if rotorModel.IsThin() {
			thin = append(thin, rotorModel)
		} else {
			normal = append(normal, rotorModel)
		}
	}

	if slot == Fourth {
		return thin // fourth rotor slot only supports thin rotors
	}
	return normal
}

func (m Model) supportsRotorModel(rotorModel RotorModel, slot RotorSlot) bool {
	for _, rot := range m.GetAvailableRotorModels(slot) {
		if rot == rotorModel {
			return true
		}
	}
	return false
}

func (m Model) getDefaultRotorModels() map[RotorSlot]RotorModel {
	placedModels := map[RotorModel]struct{}{} // cannot use the same rotor model twice
	result := map[RotorSlot]RotorModel{}
	for _, slot := range m.GetAvailableRotorSlots() {
		for _, availableModel := range m.GetAvailableRotorModels(slot) {
			if _, ok := placedModels[availableModel]; !ok {
				result[slot] = availableModel // use the first available unused rotor model
				placedModels[availableModel] = struct{}{}
				break
			}
		}
	}

	return result
}

// GetAvailableReflectorModels return all the reflectors that can be plugged into this Enigma model
func (m Model) GetAvailableReflectorModels() []ReflectorModel {
	return models[m].reflectors
}

func (m Model) supportsReflectorModel(reflectorModel ReflectorModel) bool {
	for _, ref := range m.GetAvailableReflectorModels() {
		if ref == reflectorModel {
			return true
		}
	}
	return false
}

func (m Model) getDefaultReflectorModel() ReflectorModel {
	return models[m].reflectors[0]
}

type modelDefinition struct {
	name           string
	description    string
	yearIntroduced int
	hasPlugboard   bool
	hasFourthRotor bool
	reflectors     []ReflectorModel
	rotors         []RotorModel
	etw            etwWiring
}

var models = map[Model]modelDefinition{
	Commercial: {
		name:           "Commercial K",
		description:    "Based on Enigma C model and nearly identical to the D model introduced a year earlier, this was the most successful commercial Enigma model. Its core design with three swappable rotors and a movable reflector became the basis for all later Enigma models.",
		yearIntroduced: 1927,
		hasPlugboard:   false,
		hasFourthRotor: false,
		reflectors:     []ReflectorModel{UkwK},
		rotors:         []RotorModel{RotorIK, RotorIIK, RotorIIIK},
		etw:            etwQwertz,
	},
	One: {
		name:           "Enigma I",
		description:    "Military version of the commercial Enigma D used by the German army and air force. Plugboard added for greater cryptographic security and contrary to the commercial Enigma models, the reflector was fixed. Originally supplied with just three rotors, later in 1938 two more were added.",
		yearIntroduced: 1932,
		hasPlugboard:   true,
		hasFourthRotor: false,
		reflectors:     []ReflectorModel{UkwA, UkwB},
		rotors:         []RotorModel{RotorI, RotorII, RotorIII, RotorIV, RotorV},
		etw:            etwAbcdef,
	},
	M3: {
		name:           "Enigma M3",
		description:    "Variations of Enigma I developed for the German navy. Originally fully compatible with Enigma I with five rotors to choose from, but later three more rotors added exclusively for the navy",
		yearIntroduced: 1934,
		hasPlugboard:   true,
		hasFourthRotor: false,
		reflectors:     []ReflectorModel{UkwA, UkwB, UkwC, UkwD},
		rotors:         []RotorModel{RotorI, RotorII, RotorIII, RotorIV, RotorV, RotorVI, RotorVII, RotorVIII},
		etw:            etwAbcdef,
	},
	M4: {
		name:           "Enigma M4",
		description:    "Four-rotor version of Enigma M3, developed secretly by the German navy and later used mainly for U-boat traffic. Reflector was replaced by a special thin reflector and a fourth (thinner) rotor to increase the number of key combinations.",
		yearIntroduced: 1942,
		hasPlugboard:   true,
		hasFourthRotor: true,
		reflectors:     []ReflectorModel{UkwBThin, UkwCThin},
		rotors:         []RotorModel{RotorI, RotorII, RotorIII, RotorIV, RotorV, RotorVI, RotorVII, RotorVIII, RotorBeta, RotorGamma},
		etw:            etwAbcdef,
	},
	M4UKWD: {
		name:           "Enigma M4 with UKW-D",
		description:    "Field-rewirable reflector UKW-D was introduced by the air force for the Enigma M4. Could be plugged instead of the (thin) reflector and the fourth rotor. The UKW-D settings were typically only changed once in 10 days.",
		yearIntroduced: 1944,
		hasPlugboard:   true,
		hasFourthRotor: false,
		reflectors:     []ReflectorModel{UkwD},
		rotors:         []RotorModel{RotorI, RotorII, RotorIII, RotorIV, RotorV, RotorVI, RotorVII, RotorVIII},
		etw:            etwAbcdef,
	},
	SwissK: {
		name:           "Swiss-K",
		description:    "Built for the Swiss army before WWII. Based on commercial K model, but with different rotor wiring and an extra lamp panel",
		yearIntroduced: 1938,
		hasPlugboard:   false,
		hasFourthRotor: false,
		reflectors:     []ReflectorModel{UkwK},
		rotors:         []RotorModel{RotorISK, RotorIISK, RotorIIISK},
		etw:            etwQwertz,
	},
	Tripitz: {
		name:           "Enigma T (Tripitz)",
		description:    "Used for communication between German and Japanese navies. No plugboard, but specific ETW wiring and multiple turnover notches on rotors to increase security",
		yearIntroduced: 1942,
		hasPlugboard:   false,
		hasFourthRotor: false,
		reflectors:     []ReflectorModel{UkwT},
		rotors:         []RotorModel{RotorIT, RotorIIT, RotorIIIT, RotorIVT, RotorVT, RotorVIT, RotorVIIT, RotorVIIIT},
		etw:            etwTripitz,
	},
}
