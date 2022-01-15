package enigma

// RotorModel specifies the type of rotor (different Enigma models supported different sets of rotors)
type RotorModel string

// all supported rotor models
const (
	RotorIK   RotorModel = "I-K"
	RotorIIK  RotorModel = "II-K"
	RotorIIIK RotorModel = "III-K"

	RotorI    RotorModel = "I"
	RotorII   RotorModel = "II"
	RotorIII  RotorModel = "III"
	RotorIV   RotorModel = "IV"
	RotorV    RotorModel = "V"
	RotorVI   RotorModel = "VI"
	RotorVII  RotorModel = "VII"
	RotorVIII RotorModel = "VIII"

	RotorBeta  RotorModel = "beta"
	RotorGamma RotorModel = "gamma"

	RotorISK   RotorModel = "I-SK"
	RotorIISK  RotorModel = "II-SK"
	RotorIIISK RotorModel = "III-SK"

	RotorIT    RotorModel = "I-T"
	RotorIIT   RotorModel = "II-T"
	RotorIIIT  RotorModel = "III-T"
	RotorIVT   RotorModel = "IV-T"
	RotorVT    RotorModel = "V-T"
	RotorVIT   RotorModel = "VI-T"
	RotorVIIT  RotorModel = "VII-T"
	RotorVIIIT RotorModel = "VIII-T"
)

func (r RotorModel) exists() bool {
	_, ok := rotorDefinitions[r]
	return ok
}

func (r RotorModel) getNotchPositions() []byte {
	return rotorDefinitions[r].notchPositions
}

// IsThin determines if this rotor model is thin or normal size,
// only thin rotors can be placed into the last slot of 4-rotor Enigma models
func (r RotorModel) IsThin() bool {
	return rotorDefinitions[r].isThin
}

func (r RotorModel) getWiring() string {
	return rotorDefinitions[r].wiring
}

type rotorDefinition struct {
	notchPositions []byte
	isThin         bool
	wiring         string
}

var rotorDefinitions = map[RotorModel]rotorDefinition{
	RotorIK: {
		notchPositions: []byte{'Y'},
		isThin:         false,
		wiring:         "LPGSZMHAEOQKVXRFYBUTNICJDW",
	},
	RotorIIK: {
		notchPositions: []byte{'E'},
		isThin:         false,
		wiring:         "SLVGBTFXJQOHEWIRZYAMKPCNDU",
	},
	RotorIIIK: {
		notchPositions: []byte{'N'},
		isThin:         false,
		wiring:         "CJGDPSHKTURAWZXFMYNQOBVLIE",
	},

	RotorI: {
		notchPositions: []byte{'Q'},
		isThin:         false,
		wiring:         "EKMFLGDQVZNTOWYHXUSPAIBRCJ",
	},
	RotorII: {
		notchPositions: []byte{'E'},
		isThin:         false,
		wiring:         "AJDKSIRUXBLHWTMCQGZNPYFVOE",
	},
	RotorIII: {
		notchPositions: []byte{'V'},
		isThin:         false,
		wiring:         "BDFHJLCPRTXVZNYEIWGAKMUSQO",
	},
	RotorIV: {
		notchPositions: []byte{'J'},
		isThin:         false,
		wiring:         "ESOVPZJAYQUIRHXLNFTGKDCMWB",
	},
	RotorV: {
		notchPositions: []byte{'Z'},
		isThin:         false,
		wiring:         "VZBRGITYUPSDNHLXAWMJQOFECK",
	},
	RotorVI: {
		notchPositions: []byte{'Z', 'M'},
		isThin:         false,
		wiring:         "JPGVOUMFYQBENHZRDKASXLICTW",
	},
	RotorVII: {
		notchPositions: []byte{'Z', 'M'},
		isThin:         false,
		wiring:         "NZJHGRCXMYSWBOUFAIVLPEKQDT",
	},
	RotorVIII: {
		notchPositions: []byte{'Z', 'M'},
		isThin:         false,
		wiring:         "FKQHTLXOCBJSPDZRAMEWNIUYGV",
	},

	RotorBeta: {
		notchPositions: []byte{},
		isThin:         true,
		wiring:         "LEYJVCNIXWPBQMDRTAKZGFUHOS",
	},
	RotorGamma: {
		notchPositions: []byte{},
		isThin:         true,
		wiring:         "FSOKANUERHMBTIYCWLQPZXVGJD",
	},

	RotorISK: {
		notchPositions: []byte{'Y'},
		isThin:         false,
		wiring:         "PEZUOHXSCVFMTBGLRINQJWAYDK",
	},
	RotorIISK: {
		notchPositions: []byte{'E'},
		isThin:         false,
		wiring:         "ZOUESYDKFWPCIQXHMVBLGNJRAT",
	},
	RotorIIISK: {
		notchPositions: []byte{'N'},
		isThin:         false,
		wiring:         "EHRVXGAOBQUSIMZFLYNWKTPDJC",
	},

	RotorIT: {
		notchPositions: []byte{'W', 'Z', 'E', 'K', 'Q'},
		isThin:         false,
		wiring:         "KPTYUELOCVGRFQDANJMBSWHZXI",
	},
	RotorIIT: {
		notchPositions: []byte{'W', 'Z', 'F', 'L', 'R'},
		isThin:         false,
		wiring:         "UPHZLWEQMTDJXCAKSOIGVBYFNR",
	},
	RotorIIIT: {
		notchPositions: []byte{'W', 'Z', 'E', 'K', 'Q'},
		isThin:         false,
		wiring:         "QUDLYRFEKONVZAXWHMGPJBSICT",
	},
	RotorIVT: {
		notchPositions: []byte{'W', 'Z', 'F', 'L', 'R'},
		isThin:         false,
		wiring:         "CIWTBKXNRESPFLYDAGVHQUOJZM",
	},
	RotorVT: {
		notchPositions: []byte{'Y', 'C', 'F', 'K', 'R'},
		isThin:         false,
		wiring:         "UAXGISNJBVERDYLFZWTPCKOHMQ",
	},
	RotorVIT: {
		notchPositions: []byte{'X', 'E', 'I', 'M', 'Q'},
		isThin:         false,
		wiring:         "XFUZGALVHCNYSEWQTDMRBKPIOJ",
	},
	RotorVIIT: {
		notchPositions: []byte{'Y', 'C', 'F', 'K', 'R'},
		isThin:         false,
		wiring:         "BJVFTXPLNAYOZIKWGDQERUCHSM",
	},
	RotorVIIIT: {
		notchPositions: []byte{'X', 'E', 'I', 'M', 'Q'},
		isThin:         false,
		wiring:         "YMTPNZHWKODAJXELUQVGCBISFR",
	},
}
