package enigma

type RotorType string

const (
	Rotor_IK   RotorType = "I-K"
	Rotor_IIK  RotorType = "II-K"
	Rotor_IIIK RotorType = "III-K"

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

	Rotor_ISK   RotorType = "I-SK"
	Rotor_IISK  RotorType = "II-SK"
	Rotor_IIISK RotorType = "III-SK"

	Rotor_IT    RotorType = "I-T"
	Rotor_IIT   RotorType = "II-T"
	Rotor_IIIT  RotorType = "III-T"
	Rotor_IVT   RotorType = "IV-T"
	Rotor_VT    RotorType = "V-T"
	Rotor_VIT   RotorType = "VI-T"
	Rotor_VIIT  RotorType = "VII-T"
	Rotor_VIIIT RotorType = "VIII-T"
)

func (r RotorType) exists() bool {
	_, ok := rotorDefinitions[r]
	return ok
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
	notchPositions []byte
	canBeFourth    bool
	wiring         string
}

var rotorDefinitions = map[RotorType]rotorDefinition{
	Rotor_IK: {
		notchPositions: []byte{'Y'},
		canBeFourth:    false,
		wiring:         "LPGSZMHAEOQKVXRFYBUTNICJDW",
	},
	Rotor_IIK: {
		notchPositions: []byte{'E'},
		canBeFourth:    false,
		wiring:         "SLVGBTFXJQOHEWIRZYAMKPCNDU",
	},
	Rotor_IIIK: {
		notchPositions: []byte{'N'},
		canBeFourth:    false,
		wiring:         "CJGDPSHKTURAWZXFMYNQOBVLIE",
	},

	Rotor_I: {
		notchPositions: []byte{'Q'},
		canBeFourth:    false,
		wiring:         "EKMFLGDQVZNTOWYHXUSPAIBRCJ",
	},
	Rotor_II: {
		notchPositions: []byte{'E'},
		canBeFourth:    false,
		wiring:         "AJDKSIRUXBLHWTMCQGZNPYFVOE",
	},
	Rotor_III: {
		notchPositions: []byte{'V'},
		canBeFourth:    false,
		wiring:         "BDFHJLCPRTXVZNYEIWGAKMUSQO",
	},
	Rotor_IV: {
		notchPositions: []byte{'J'},
		canBeFourth:    false,
		wiring:         "ESOVPZJAYQUIRHXLNFTGKDCMWB",
	},
	Rotor_V: {
		notchPositions: []byte{'Z'},
		canBeFourth:    false,
		wiring:         "VZBRGITYUPSDNHLXAWMJQOFECK",
	},
	Rotor_VI: {
		notchPositions: []byte{'Z', 'M'},
		canBeFourth:    false,
		wiring:         "JPGVOUMFYQBENHZRDKASXLICTW",
	},
	Rotor_VII: {
		notchPositions: []byte{'Z', 'M'},
		canBeFourth:    false,
		wiring:         "NZJHGRCXMYSWBOUFAIVLPEKQDT",
	},
	Rotor_VIII: {
		notchPositions: []byte{'Z', 'M'},
		canBeFourth:    false,
		wiring:         "FKQHTLXOCBJSPDZRAMEWNIUYGV",
	},

	Rotor_beta: {
		notchPositions: []byte{},
		canBeFourth:    true,
		wiring:         "LEYJVCNIXWPBQMDRTAKZGFUHOS",
	},
	Rotor_gamma: {
		notchPositions: []byte{},
		canBeFourth:    true,
		wiring:         "FSOKANUERHMBTIYCWLQPZXVGJD",
	},

	Rotor_ISK: {
		notchPositions: []byte{'Y'},
		canBeFourth:    false,
		wiring:         "PEZUOHXSCVFMTBGLRINQJWAYDK",
	},
	Rotor_IISK: {
		notchPositions: []byte{'E'},
		canBeFourth:    false,
		wiring:         "ZOUESYDKFWPCIQXHMVBLGNJRAT",
	},
	Rotor_IIISK: {
		notchPositions: []byte{'N'},
		canBeFourth:    false,
		wiring:         "EHRVXGAOBQUSIMZFLYNWKTPDJC",
	},

	Rotor_IT: {
		notchPositions: []byte{'W', 'Z', 'E', 'K', 'Q'},
		canBeFourth:    false,
		wiring:         "KPTYUELOCVGRFQDANJMBSWHZXI",
	},
	Rotor_IIT: {
		notchPositions: []byte{'W', 'Z', 'F', 'L', 'R'},
		canBeFourth:    false,
		wiring:         "UPHZLWEQMTDJXCAKSOIGVBYFNR",
	},
	Rotor_IIIT: {
		notchPositions: []byte{'W', 'Z', 'E', 'K', 'Q'},
		canBeFourth:    false,
		wiring:         "QUDLYRFEKONVZAXWHMGPJBSICT",
	},
	Rotor_IVT: {
		notchPositions: []byte{'W', 'Z', 'F', 'L', 'R'},
		canBeFourth:    false,
		wiring:         "CIWTBKXNRESPFLYDAGVHQUOJZM",
	},
	Rotor_VT: {
		notchPositions: []byte{'Y', 'C', 'F', 'K', 'R'},
		canBeFourth:    false,
		wiring:         "UAXGISNJBVERDYLFZWTPCKOHMQ",
	},
	Rotor_VIT: {
		notchPositions: []byte{'X', 'E', 'I', 'M', 'Q'},
		canBeFourth:    false,
		wiring:         "XFUZGALVHCNYSEWQTDMRBKPIOJ",
	},
	Rotor_VIIT: {
		notchPositions: []byte{'Y', 'C', 'F', 'K', 'R'},
		canBeFourth:    false,
		wiring:         "BJVFTXPLNAYOZIKWGDQERUCHSM",
	},
	Rotor_VIIIT: {
		notchPositions: []byte{'X', 'E', 'I', 'M', 'Q'},
		canBeFourth:    false,
		wiring:         "YMTPNZHWKODAJXELUQVGCBISFR",
	},
}
