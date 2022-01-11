package enigma

import (
	"fmt"
	"strings"
)

type reflector struct {
	reflectorType  reflectorType
	translationMap map[int]int
	position       int
}

func newReflector(reflectorType reflectorType) reflector {
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
		reflectorType:  reflectorType,
		translationMap: letterMap,
		position:       0,
	}
}

func (r *reflector) setWiring(wiring string) error {
	// todo - do not allow UKW-D without custom wiring (or provide default one)
	if !r.reflectorType.isRewirable() {
		return fmt.Errorf("reflector %s is not rewirable, cannot change wiring", r.reflectorType)
	}

	// UKW-D rewirable reflectors had different letter order (JY were always connected, the rest 12 pairs were configurable)
	ukwdOrder := "AJZXWVUTSRQPONYMLKIHGFEDCB"
	wiringMap := getDefaultWiring()
	wiring += " JY" //todo - do better (check wiring before changing it to avoid misleading error messages)

	// rewire the reflector
	pairs := strings.Split(wiring, " ")
	expectedSize := Alphabet.getSize() / 2
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

	r.translationMap = wiringMap
	return nil
}

func (r *reflector) setPosition(position byte) error {
	if len(r.reflectorType.getPositions()) <= 1 {
		return fmt.Errorf("reflector %s is fixed, cannot change position", r.reflectorType)
	}
	index, ok := Alphabet.charToInt(position)
	if !ok {
		return fmt.Errorf("invalid reflector position %s", string(position))
	}
	supported := false
	for _, pos := range r.reflectorType.getPositions() {
		if pos == index {
			supported = true
			break
		}
	}
	if !supported {
		return fmt.Errorf("reflector %s does not support position %s", r.reflectorType, string(position))
	}

	r.position = index
	return nil
}

func (r *reflector) translate(input int) int {
	rotatedOutput := r.translationMap[shift(input, r.position)]
	return shift(rotatedOutput, -r.position) // don't forget to rotate back...
}

func getDefaultWiring() map[int]int {
	letterMap := make(map[int]int, Alphabet.getSize())
	for i := 0; i < Alphabet.getSize(); i++ {
		letterMap[i] = i
	}
	return letterMap
}