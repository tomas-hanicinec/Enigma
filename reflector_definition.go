package enigma

type ReflectorType string

const (
	UkwK     ReflectorType = "K"
	UkwA     ReflectorType = "A"
	UkwB     ReflectorType = "B"
	UkwC     ReflectorType = "C"
	UkwBThin ReflectorType = "BThin"
	UkwCThin ReflectorType = "CThin"
	UkwD     ReflectorType = "D"
	UkwT     ReflectorType = "T"
)

func (r ReflectorType) IsRewirable() bool {
	return reflectorDefinitions[r].isRewirable
}

func (r ReflectorType) IsThin() bool {
	return reflectorDefinitions[r].isThin
}

func (r ReflectorType) IsMovable() bool {
	return reflectorDefinitions[r].isMovable
}

func (r ReflectorType) getWiring() string {
	return reflectorDefinitions[r].wiring
}

type reflectorDefinition struct {
	isRewirable bool
	isMovable   bool
	isThin      bool
	wiring      string
}

var reflectorDefinitions = map[ReflectorType]reflectorDefinition{
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
