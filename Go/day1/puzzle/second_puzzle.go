package puzzle

import (
	"fmt"
	"strconv"
)

func checkIfItIsDigit(line string, i int, initialLettersSlice []string) (string, bool) {
	strDigitToInt := map[string]string{
		"one": "1",
		"two": "2",
		"three": "3",
		"four": "4",
		"five": "5",
		"six": "6",
		"seven": "7",
		"eight": "8",
		"nine": "9",
		"ten": "10",
	}

	for j := 1; j <= 5 && i+j <= len(line); j++ {
		currWord := string(line[i:i+j])
		for _, strDigit := range initialLettersSlice {
			if currWord == strDigit {
				return strDigitToInt[currWord], true
			}
		}
	}
	return "", false
}

func SecondPuzzle(linesList []string) int {
	initialLetters := map[string][]string{
		"o": {"one"},
		"t": {"two", "three"},
		"f": {"four", "five"},
		"s": {"six", "seven"},
		"e": {"eight"},
		"n": {"nine"},
	}
	answer := 0
	for _, line := range linesList {
		var leftDigit string
		var rightDigit string
		l := 0
		r := len(line) - 1
		for l <= r && (leftDigit == "" || rightDigit == "") {
			// from left
			if initialLettersSlice, ok := initialLetters[string(line[l])]; ok {
				strDigit, itsDigit := checkIfItIsDigit(line, l, initialLettersSlice)
				if itsDigit {
					leftDigit = strDigit
				} else {
					l += 1
				}
			} else if _, err := strconv.Atoi(string(line[l])); err == nil {
				leftDigit = string(line[l])
			} else {
				l += 1
			}

			// from right
			if initialLettersSlice, ok := initialLetters[string(line[r])]; ok {
				strDigit, itsDigit := checkIfItIsDigit(line, r, initialLettersSlice)
				if itsDigit {
					rightDigit = strDigit
				} else {
					r -= 1
				}
			} else if _, err := strconv.Atoi(string(line[r])); err == nil {
				rightDigit = string(line[r])
			} else {
				r -= 1
			}
		}

		twoDigitNum, err := strconv.Atoi(leftDigit + rightDigit)
		if err != nil {
			panic(err)
		}
		answer += twoDigitNum
	}

	fmt.Println("Second puzzle answer: ", answer)
	return answer
}