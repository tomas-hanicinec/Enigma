package enigma

import (
	"fmt"
	"strings"
)

type ReflectorConfig struct {
	ReflectorType ReflectorType
	WheelPosition byte
	Wiring        string
}

func (r ReflectorConfig) isEmpty() bool {
	return r.ReflectorType == "" && r.WheelPosition == 0 && r.Wiring == ""
}

type reflector struct {
	reflectorType ReflectorType
	letterMap     map[int]int
	wheelPosition int
}

func newReflector(reflectorType ReflectorType) reflector {
	wiring := reflectorType.getWiring()
	if !Alphabet.isValidWiring(wiring) {
		panic(fmt.Errorf("invalid reflector wiring %s", wiring))
	}

	letterMap := make(map[int]int, Alphabet.getSize())
	for i, letter := range wiring {
		letterIndex, ok := Alphabet.charToInt(byte(letter))
		if !ok {
			panic(fmt.Errorf("unsupported wiring letter %s", string(letter))) // should not happen, we already checked the wiring validity
		}
		letterMap[i] = letterIndex
		letterMap[letterIndex] = i
	}

	return reflector{
		reflectorType: reflectorType,
		letterMap:     letterMap,
		wheelPosition: 0,
	}
}

func (r *reflector) setWheelPosition(letter byte) error {
	if !r.reflectorType.IsMovable() {
		return fmt.Errorf("reflector %s is fixed, cannot change position", r.reflectorType)
	}
	index, ok := Alphabet.charToInt(letter)
	if !ok {
		return fmt.Errorf("unsupported reflector position %s", string(letter))
	}

	r.wheelPosition = index
	return nil
}

func (r *reflector) setWiring(wiring string) error {
	if !r.reflectorType.IsRewirable() {
		return fmt.Errorf("reflector %s is not rewirable, cannot change wiring", r.reflectorType)
	}

	// UKW-D rewirable reflectors had different letter order (JY were always connected, the rest 12 pairs were configurable)
	ukwdOrder := "AJZXWVUTSRQPONYMLKIHGFEDCB"
	wiringMap := getDefaultLetterMap()
	wiringMap[strings.IndexByte(ukwdOrder, 'J')] = strings.IndexByte(ukwdOrder, 'Y')
	wiringMap[strings.IndexByte(ukwdOrder, 'Y')] = strings.IndexByte(ukwdOrder, 'J')

	// rewire the reflector
	pairs := strings.Split(wiring, " ")
	expectedSize := Alphabet.getSize()/2 - 1
	if len(pairs) != expectedSize {
		return fmt.Errorf("incomplete wiring of the reflector, must include %d pairs to cover the whole alphabet", expectedSize)
	}
	for _, pair := range pairs {
		// validate the pair
		if len(pair) != 2 {
			return fmt.Errorf("invalid pair %s, must be a pair of letters", pair)
		}
		if pair[0] == pair[1] {
			return fmt.Errorf("invalid pair %s, cannot connect reflector letter to itself", pair)
		}
		var letters [2]int
		for i := 0; i < 2; i++ {
			index := strings.IndexByte(ukwdOrder, pair[i])
			if index == -1 {
				return fmt.Errorf("invalid pair %s, unsupported letter %s", pair, string(pair[i]))
			}
			letters[i] = index
			if mapped, ok := wiringMap[letters[i]]; ok && mapped != letters[i] {
				return fmt.Errorf("invalid pair %s, letter %s (%s) already wired", pair, string(pair[i]), string(ukwdOrder[index]))
			}
		}

		// set to map
		wiringMap[letters[0]] = letters[1]
		wiringMap[letters[1]] = letters[0]
	}

	r.letterMap = wiringMap
	return nil
}

func (r *reflector) translate(input int) int {
	rotatedOutput := r.letterMap[shift(input, r.wheelPosition)]
	return shift(rotatedOutput, -r.wheelPosition) // don't forget to rotate back...
}
