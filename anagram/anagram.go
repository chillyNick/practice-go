package anagram

import (
	"strings"
	"unicode"
)

func FindAnagrams(dictionary []string, word string) (result []string) {
	lw := strings.ToLower(word)

	rc := getRunesCount(lw)
	for _, w := range dictionary {
		if lw == strings.ToLower(w) || w == "" {
			continue
		}

		if isEqual(rc, getRunesCount(w)) {
			result = append(result, w)
		}
	}

	return
}

func getRunesCount(word string) (result map[rune]int) {
	result = make(map[rune]int)

	for _, r := range word {
		if unicode.IsSpace(r) || unicode.IsMark(r) {
			continue
		}

		result[unicode.ToLower(r)]++
	}

	return
}

func isEqual(x, y map[rune]int) bool {
	if len(x) != len(y) {
		return false
	}

	for k, _ := range x {
		if v, ok := y[k]; !ok || x[k] != v {
			return false
		}
	}

	return true
}
