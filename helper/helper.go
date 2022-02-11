package helper

import (
	"strings"
)

func RepeatedCharacters(text string) bool {
	valuesToCheck := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "x", "y", "z"}
	var resultStorage []int
	hasRepeatedCharacter := false

	for _, letter := range valuesToCheck {
		result := strings.Count(text, letter)
		resultStorage = append(resultStorage, result)
	}
	for _, result := range resultStorage {
		if result > 1 {
			hasRepeatedCharacter = true
			break
		}
	}
	return hasRepeatedCharacter
}
