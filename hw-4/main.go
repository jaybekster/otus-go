package main

import (
	"strings"
	"fmt"
	"unicode"
)

func freqentWords(input string) map[string]int {
	output := make(map[string]int)

	for _, word := range strings.FieldsFunc(input, func(c rune) bool {
		return !unicode.IsLetter(c)
	}) {
		key := strings.ToLower(word)
		output[key] = output[key] + 1;
	}

	return output
}

func main() {
	fmt.Println(freqentWords("hello word hello"))

	fmt.Println(freqentWords("Hello word, hellO!"))
}