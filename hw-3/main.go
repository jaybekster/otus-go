package main

import (
	"strings"
	"strconv"
	"fmt"
	"unicode/utf8"
)

func checkIfNumber(ch rune) (int, bool) {
	number, err := strconv.Atoi(string(ch))

	if err != nil {
		return 0, false;
	}

	return number, true;
}

func decompressString(input string) {
	var output string
	var isInvert bool

	for _, ch := range input {
		chString := string(ch)

		if isInvert {
			output += chString
			isInvert = false
		} else if count, isNumber := checkIfNumber(ch); isNumber {
			if len(output) < 1 {
				output = "";
				break;
			}
			
			lastRune, _ := utf8.DecodeLastRuneInString(output)
			output += strings.Repeat(string(lastRune), count - 1);
		} else if chString == `\` {
			isInvert = true
		} else {
			output += chString;
		}
	}

	fmt.Println(output);
}

func main() {
	decompressString("a4bc2d5e")
	decompressString("abcd")
	decompressString("23")

	decompressString(`qwe\4\5`)
	decompressString(`qwe\45`)
	decompressString(`qwe\\5`)
}