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

func (r RotorType) IsThin() bool {
	return rotorDefinitions[r].isThin
}

func (r RotorType) getWiring() string {
	return rotorDefinitions[r].wiring
}

type rotorDefinition struct {
	notchPositions []byte
	isThin         bool
	wiring         string
}

var rotorDefinitions = map[RotorType]rotorDefinition{
	Rotor_IK: {
		notchPositions: []byte{'Y'},
		isThin:         false,
		wiring:         "LPGSZMHAEOQKVXRFYBUTNICJDW",
	},
	Rotor_IIK: {
		notchPositions: []byte{'E'},
		isThin:         false,
		wiring:         "SLVGBTFXJQOHEWIRZYAMKPCNDU",
	},
	Rotor_IIIK: {
		notchPositions: []byte{'N'},
		isThin:         false,
		wiring:         "CJGDPSHKTURAWZXFMYNQOBVLIE",
	},

	Rotor_I: {
		notchPositions: []byte{'Q'},
		isThin:         false,
		wiring:         "EKMFLGDQVZNTOWYHXUSPAIBRCJ",
	},
	Rotor_II: {
		notchPositions: []byte{'E'},
		isThin:         false,
		wiring:         "AJDKSIRUXBLHWTMCQGZNPYFVOE",
	},
	Rotor_III: {
		notchPositions: []byte{'V'},
		isThin:         false,
		wiring:         "BDFHJLCPRTXVZNYEIWGAKMUSQO",
	},
	Rotor_IV: {
		notchPositions: []byte{'J'},
		isThin:         false,
		wiring:         "ESOVPZJAYQUIRHXLNFTGKDCMWB",
	},
	Rotor_V: {
		notchPositions: []byte{'Z'},
		isThin:         false,
		wiring:         "VZBRGITYUPSDNHLXAWMJQOFECK",
	},
	Rotor_VI: {
		notchPositions: []byte{'Z', 'M'},
		isThin:         false,
		wiring:         "JPGVOUMFYQBENHZRDKASXLICTW",
	},
	Rotor_VII: {
		notchPositions: []byte{'Z', 'M'},
		isThin:         false,
		wiring:         "NZJHGRCXMYSWBOUFAIVLPEKQDT",
	},
	Rotor_VIII: {
		notchPositions: []byte{'Z', 'M'},
		isThin:         false,
		wiring:         "FKQHTLXOCBJSPDZRAMEWNIUYGV",
	},

	Rotor_beta: {
		notchPositions: []byte{},
		isThin:         true,
		wiring:         "LEYJVCNIXWPBQMDRTAKZGFUHOS",
	},
	Rotor_gamma: {
		notchPositions: []byte{},
		isThin:         true,
		wiring:         "FSOKANUERHMBTIYCWLQPZXVGJD",
	},

	Rotor_ISK: {
		notchPositions: []byte{'Y'},
		isThin:         false,
		wiring:         "PEZUOHXSCVFMTBGLRINQJWAYDK",
	},
	Rotor_IISK: {
		notchPositions: []byte{'E'},
		isThin:         false,
		wiring:         "ZOUESYDKFWPCIQXHMVBLGNJRAT",
	},
	Rotor_IIISK: {
		notchPositions: []byte{'N'},
		isThin:         false,
		wiring:         "EHRVXGAOBQUSIMZFLYNWKTPDJC",
	},

	Rotor_IT: {
		notchPositions: []byte{'W', 'Z', 'E', 'K', 'Q'},
		isThin:         false,
		wiring:         "KPTYUELOCVGRFQDANJMBSWHZXI",
	},
	Rotor_IIT: {
		notchPositions: []byte{'W', 'Z', 'F', 'L', 'R'},
		isThin:         false,
		wiring:         "UPHZLWEQMTDJXCAKSOIGVBYFNR",
	},
	Rotor_IIIT: {
		notchPositions: []byte{'W', 'Z', 'E', 'K', 'Q'},
		isThin:         false,
		wiring:         "QUDLYRFEKONVZAXWHMGPJBSICT",
	},
	Rotor_IVT: {
		notchPositions: []byte{'W', 'Z', 'F', 'L', 'R'},
		isThin:         false,
		wiring:         "CIWTBKXNRESPFLYDAGVHQUOJZM",
	},
	Rotor_VT: {
		notchPositions: []byte{'Y', 'C', 'F', 'K', 'R'},
		isThin:         false,
		wiring:         "UAXGISNJBVERDYLFZWTPCKOHMQ",
	},
	Rotor_VIT: {
		notchPositions: []byte{'X', 'E', 'I', 'M', 'Q'},
		isThin:         false,
		wiring:         "XFUZGALVHCNYSEWQTDMRBKPIOJ",
	},
	Rotor_VIIT: {
		notchPositions: []byte{'Y', 'C', 'F', 'K', 'R'},
		isThin:         false,
		wiring:         "BJVFTXPLNAYOZIKWGDQERUCHSM",
	},
	Rotor_VIIIT: {
		notchPositions: []byte{'X', 'E', 'I', 'M', 'Q'},
		isThin:         false,
		wiring:         "YMTPNZHWKODAJXELUQVGCBISFR",
	},
}
