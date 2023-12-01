package puzzle

import (
	"fmt"
	"strconv"
)

func FirstPuzzle(linesList []string) int {
	answer := 0
	for _, line := range linesList {
		var leftDigit string
		var rightDigit string
		l := 0
		r := len(line) - 1
		for l <= r && (leftDigit == "" || rightDigit == "") {
			// from left
			if _, err := strconv.Atoi(string(line[l])); err == nil {
				leftDigit = string(line[l])
			} else {
				l += 1
			}

			// from right
			if _, err := strconv.Atoi(string(line[r])); err == nil {
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

	fmt.Println("First puzzle answer: ", answer)
	return answer
}
