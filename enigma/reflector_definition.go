package enigma

type ReflectorType string

const (
	UKW_A     ReflectorType = "A"
	UKW_B     ReflectorType = "B"
	UKW_C     ReflectorType = "C"
	UKW_BThin ReflectorType = "BThin"
	UKW_CThin ReflectorType = "CThin"
	UKW_D     ReflectorType = "D"
	UKW_K     ReflectorType = "K"
)

func (r ReflectorType) IsRewirable() bool {
	return reflectorDefinitions[r].isRewirable
}

func (r ReflectorType) IsThin() bool {
	return reflectorDefinitions[r].isThin
}

func (r ReflectorType) IsMovable() bool {
	// todo - is there something else than fully rotatable reflector in commercial enigmas? if not, cancel the explicit positions array
	return len(reflectorDefinitions[r].positions) > 0
}

func (r ReflectorType) getPositions() []int {
	return reflectorDefinitions[r].positions
}

func (r ReflectorType) getWiring() string {
	return reflectorDefinitions[r].wiring
}

type reflectorDefinition struct {
	isRewirable bool
	isThin      bool
	positions   []int
	wiring      string
}

var reflectorDefinitions = map[ReflectorType]reflectorDefinition{
	UKW_K: {
		isRewirable: false,
		isThin:      false,
		positions:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26},
		wiring:      "IMETCGFRAYSQBZXWLHKDVUPOJN",
	},
	UKW_A: {
		isRewirable: false,
		isThin:      false,
		positions:   nil,
		wiring:      "EJMZALYXVBWFCRQUONTSPIKHGD",
	},
	UKW_B: {
		isRewirable: false,
		isThin:      false,
		positions:   nil,
		wiring:      "YRUHQSLDPXNGOKMIEBFZCWVJAT",
	},
	UKW_C: {
		isRewirable: false,
		isThin:      false,
		positions:   nil,
		wiring:      "FVPJIAOYEDRZXWGCTKUQSBNMHL",
	},
	UKW_BThin: {
		isRewirable: false,
		isThin:      true,
		positions:   nil,
		wiring:      "ENKQAUYWJICOPBLMDXZVFTHRGS",
	},
	UKW_CThin: {
		isRewirable: false,
		isThin:      true,
		positions:   nil,
		wiring:      "RDOBJNTKVEHMLFCWZAXGYIPSUQ",
	},
	UKW_D: {
		isRewirable: true,
		isThin:      false,
		positions:   nil,
		wiring:      "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
	},
}
