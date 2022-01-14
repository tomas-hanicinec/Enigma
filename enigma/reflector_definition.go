package enigma

type ReflectorType string

const (
	UKW_K     ReflectorType = "K"
	UKW_A     ReflectorType = "A"
	UKW_B     ReflectorType = "B"
	UKW_C     ReflectorType = "C"
	UKW_BThin ReflectorType = "BThin"
	UKW_CThin ReflectorType = "CThin"
	UKW_D     ReflectorType = "D"
	UKW_T     ReflectorType = "T"
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
	UKW_K: {
		isRewirable: false,
		isMovable:   true,
		isThin:      false,
		wiring:      "IMETCGFRAYSQBZXWLHKDVUPOJN",
	},
	UKW_A: {
		isRewirable: false,
		isMovable:   false,
		isThin:      false,
		wiring:      "EJMZALYXVBWFCRQUONTSPIKHGD",
	},
	UKW_B: {
		isRewirable: false,
		isMovable:   false,
		isThin:      false,
		wiring:      "YRUHQSLDPXNGOKMIEBFZCWVJAT",
	},
	UKW_C: {
		isRewirable: false,
		isMovable:   false,
		isThin:      false,
		wiring:      "FVPJIAOYEDRZXWGCTKUQSBNMHL",
	},
	UKW_BThin: {
		isRewirable: false,
		isMovable:   false,
		isThin:      true,
		wiring:      "ENKQAUYWJICOPBLMDXZVFTHRGS",
	},
	UKW_CThin: {
		isRewirable: false,
		isMovable:   false,
		isThin:      true,
		wiring:      "RDOBJNTKVEHMLFCWZAXGYIPSUQ",
	},
	UKW_D: {
		isRewirable: true,
		isMovable:   false,
		isThin:      false,
		wiring:      "FOWULAQYSRTEZVBXGJIKDNCPHM", // corresponds to the wiring "AV BO CT DM EZ FN GX HQ IS KR LU PW"
	},
	UKW_T: {
		isRewirable: false,
		isMovable:   false,
		isThin:      false,
		wiring:      "GEKPBTAUMOCNILJDXZYFHWVQSR",
	},
}
