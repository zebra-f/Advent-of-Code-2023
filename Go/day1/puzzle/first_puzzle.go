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
			if _, err := strconv.Atoi(string(line[l])); err == nil {
				leftDigit = string(line[l])
			} else {
				l += 1
			}
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
	fmt.Println("answer: ", answer)
	return answer
}
