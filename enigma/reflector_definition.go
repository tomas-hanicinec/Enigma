package enigma

type reflectorType string

const (
	UKW_A     reflectorType = "A"
	UKW_B                   = "B"
	UKW_C                   = "C"
	UKW_BThin               = "BThin"
	UKW_CThin               = "CThin"
	UKW_D                   = "D"
	UKW_K                   = "K"
)

func (r reflectorType) isRewirable() bool {
	return reflectorDefinitions[r].isRewirable
}

func (r reflectorType) isThin() bool {
	return reflectorDefinitions[r].isThin
}

func (r reflectorType) getPositions() []int {
	return reflectorDefinitions[r].positions
}

func (r reflectorType) getWiring() string {
	return reflectorDefinitions[r].wiring
}

type reflectorDefinition struct {
	isRewirable bool
	isThin      bool
	positions   []int
	wiring      string
}

var reflectorDefinitions = map[reflectorType]reflectorDefinition{
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
