package main

import (
	"fmt"
	"sort"
)

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func sortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}

func translateLatter(letterMap map[Char]Char, letter Char) (Char, error) {
	translated, ok := letterMap[letter]
	if !ok {
		return 0, fmt.Errorf("unsupported letter %s", string(letter))
	}
	return translated, nil
}
