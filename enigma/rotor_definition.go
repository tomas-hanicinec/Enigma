package enigma

type RotorType string

const (
	Rotor_I    RotorType = "I"
	Rotor_II   RotorType = "II"
	Rotor_III  RotorType = "III"
	Rotor_IV   RotorType = "IV"
	Rotor_V    RotorType = "V"
	Rotor_VI   RotorType = "VI"
	Rotor_VII  RotorType = "VII"
	Rotor_VIII RotorType = "VIII"

	Rotor_beta  RotorType = "beta"
	Rotor_gamma RotorType = "gamma"

	Rotor_IK   RotorType = "I-K"
	Rotor_IIK  RotorType = "II-K"
	Rotor_IIIK RotorType = "III-K"
)

func (r RotorType) exists() bool {
	_, ok := rotorDefinitions[r]
	return ok
}

func (r RotorType) HasRing() bool {
	return rotorDefinitions[r].hasRing
}

func (r RotorType) getNotchPositions() []byte {
	return rotorDefinitions[r].notchPositions
}

func (r RotorType) CanBeFourth() bool {
	return rotorDefinitions[r].canBeFourth
}

func (r RotorType) getWiring() string {
	return rotorDefinitions[r].wiring
}

type rotorDefinition struct {
	hasRing        bool
	notchPositions []byte
	canBeFourth    bool
	wiring         string
}

var rotorDefinitions = map[RotorType]rotorDefinition{
	Rotor_I: {
		hasRing:        true,
		notchPositions: []byte{'Q'},
		canBeFourth:    false,
		wiring:         "EKMFLGDQVZNTOWYHXUSPAIBRCJ",
	},
	Rotor_II: {
		hasRing:        true,
		notchPositions: []byte{'E'},
		canBeFourth:    false,
		wiring:         "AJDKSIRUXBLHWTMCQGZNPYFVOE",
	},
	Rotor_III: {
		hasRing:        true,
		notchPositions: []byte{'V'},
		canBeFourth:    false,
		wiring:         "BDFHJLCPRTXVZNYEIWGAKMUSQO",
	},
	Rotor_IV: {
		hasRing:        true,
		notchPositions: []byte{'J'},
		canBeFourth:    false,
		wiring:         "ESOVPZJAYQUIRHXLNFTGKDCMWB",
	},
	Rotor_V: {
		hasRing:        true,
		notchPositions: []byte{'Z'},
		canBeFourth:    false,
		wiring:         "VZBRGITYUPSDNHLXAWMJQOFECK",
	},
	Rotor_VI: {
		hasRing:        true,
		notchPositions: []byte{'Z', 'M'},
		canBeFourth:    false,
		wiring:         "JPGVOUMFYQBENHZRDKASXLICTW",
	},
	Rotor_VII: {
		hasRing:        true,
		notchPositions: []byte{'Z', 'M'},
		canBeFourth:    false,
		wiring:         "NZJHGRCXMYSWBOUFAIVLPEKQDT",
	},
	Rotor_VIII: {
		hasRing:        true,
		notchPositions: []byte{'Z', 'M'},
		canBeFourth:    false,
		wiring:         "FKQHTLXOCBJSPDZRAMEWNIUYGV",
	},

	Rotor_beta: {
		hasRing:        true,
		notchPositions: []byte{},
		canBeFourth:    true,
		wiring:         "LEYJVCNIXWPBQMDRTAKZGFUHOS",
	},
	Rotor_gamma: {
		hasRing:        true,
		notchPositions: []byte{},
		canBeFourth:    true,
		wiring:         "FSOKANUERHMBTIYCWLQPZXVGJD",
	},
	Rotor_IK: {
		hasRing:        true,
		notchPositions: []byte{'Y'},
		canBeFourth:    false,
		wiring:         "PEZUOHXSCVFMTBGLRINQJWAYDK",
	},
	Rotor_IIK: {
		hasRing:        true,
		notchPositions: []byte{'E'},
		canBeFourth:    false,
		wiring:         "ZOUESYDKFWPCIQXHMVBLGNJRAT",
	},
	Rotor_IIIK: {
		hasRing:        true,
		notchPositions: []byte{'N'},
		canBeFourth:    false,
		wiring:         "EHRVXGAOBQUSIMZFLYNWKTPDJC",
	},
}
