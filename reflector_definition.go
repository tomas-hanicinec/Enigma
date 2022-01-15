package enigma

// ReflectorModel specifies the type of reflector (different Enigma models supported different reflectors)
type ReflectorModel string

// all supported reflector models
const (
	UkwK     ReflectorModel = "K"
	UkwA     ReflectorModel = "A"
	UkwB     ReflectorModel = "B"
	UkwC     ReflectorModel = "C"
	UkwBThin ReflectorModel = "BThin"
	UkwCThin ReflectorModel = "CThin"
	UkwD     ReflectorModel = "D"
	UkwT     ReflectorModel = "T"
)

// IsThin shows whether this reflector model is thin, or normal size,
// only thin reflectors can fit to the small slot in 4-rotor Enigma model
func (r ReflectorModel) IsThin() bool {
	return reflectorDefinitions[r].isThin
}

// IsMovable shows if this reflector model can rotate
func (r ReflectorModel) IsMovable() bool {
	return reflectorDefinitions[r].isMovable
}

// IsRewirable shows if this reflector model can be custom-rewired
func (r ReflectorModel) IsRewirable() bool {
	return reflectorDefinitions[r].isRewirable
}

func (r ReflectorModel) getWiring() string {
	return reflectorDefinitions[r].wiring
}

type reflectorDefinition struct {
	isRewirable bool
	isMovable   bool
	isThin      bool
	wiring      string
}

var reflectorDefinitions = map[ReflectorModel]reflectorDefinition{
	UkwK: {
		isRewirable: false,
		isMovable:   true,
		isThin:      false,
		wiring:      "IMETCGFRAYSQBZXWLHKDVUPOJN",
	},
	UkwA: {
		isRewirable: false,
		isMovable:   false,
		isThin:      false,
		wiring:      "EJMZALYXVBWFCRQUONTSPIKHGD",
	},
	UkwB: {
		isRewirable: false,
		isMovable:   false,
		isThin:      false,
		wiring:      "YRUHQSLDPXNGOKMIEBFZCWVJAT",
	},
	UkwC: {
		isRewirable: false,
		isMovable:   false,
		isThin:      false,
		wiring:      "FVPJIAOYEDRZXWGCTKUQSBNMHL",
	},
	UkwBThin: {
		isRewirable: false,
		isMovable:   false,
		isThin:      true,
		wiring:      "ENKQAUYWJICOPBLMDXZVFTHRGS",
	},
	UkwCThin: {
		isRewirable: false,
		isMovable:   false,
		isThin:      true,
		wiring:      "RDOBJNTKVEHMLFCWZAXGYIPSUQ",
	},
	UkwD: {
		isRewirable: true,
		isMovable:   false,
		isThin:      false,
		wiring:      "FOWULAQYSRTEZVBXGJIKDNCPHM", // corresponds to the wiring "AV BO CT DM EZ FN GX HQ IS KR LU PW"
	},
	UkwT: {
		isRewirable: false,
		isMovable:   false,
		isThin:      false,
		wiring:      "GEKPBTAUMOCNILJDXZYFHWVQSR",
	},
}
