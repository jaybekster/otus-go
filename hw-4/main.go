package main

import (
	"strings"
	"fmt"
	"unicode"
	"sort"
)

func freqentWords(input string) []string {
	type WordsArrayItem struct {
		name string
		total int
	}

	counter := make(map[string]int)
	wordsArray := []WordsArrayItem{}
	output := make([]string, 0)

	for _, word := range strings.FieldsFunc(input, func(c rune) bool {
		return !unicode.IsLetter(c)
	}) {
		key := strings.ToLower(word)
		counter[key] = counter[key] + 1;
	}

	for k, v := range counter {
		wordsArray = append(wordsArray, WordsArrayItem{ k, v })
	}

	sort.SliceStable(wordsArray, func(i, j int) bool {
		return wordsArray[i].total > wordsArray[j].total
	})

	for i := 0; i < 10; i++ {
		if len(wordsArray) <= i {
			break;
		}
		output = append(output, wordsArray[i].name)
	}

	return output
}

func main() {
	fmt.Println(freqentWords("hello world hello"))

	fmt.Println(freqentWords("Hello world, hellO!"))

	fmt.Println(freqentWords("ops oops oops oops oops asd asd, hellO!"))
}