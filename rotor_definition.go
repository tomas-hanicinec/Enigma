package enigma

type RotorType string

const (
	RotorIK   RotorType = "I-K"
	RotorIIK  RotorType = "II-K"
	RotorIIIK RotorType = "III-K"

	RotorI    RotorType = "I"
	RotorII   RotorType = "II"
	RotorIII  RotorType = "III"
	RotorIV   RotorType = "IV"
	RotorV    RotorType = "V"
	RotorVI   RotorType = "VI"
	RotorVII  RotorType = "VII"
	RotorVIII RotorType = "VIII"

	RotorBeta  RotorType = "beta"
	RotorGamma RotorType = "gamma"

	RotorISK   RotorType = "I-SK"
	RotorIISK  RotorType = "II-SK"
	RotorIIISK RotorType = "III-SK"

	RotorIT    RotorType = "I-T"
	RotorIIT   RotorType = "II-T"
	RotorIIIT  RotorType = "III-T"
	RotorIVT   RotorType = "IV-T"
	RotorVT    RotorType = "V-T"
	RotorVIT   RotorType = "VI-T"
	RotorVIIT  RotorType = "VII-T"
	RotorVIIIT RotorType = "VIII-T"
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
