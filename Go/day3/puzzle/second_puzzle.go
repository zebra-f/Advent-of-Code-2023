package puzzle

import (
	"fmt"
	"strconv"

	"github.com/zebra-f/Advent-of-Code-2023/day3/common"
)

func getNumber(i, j int, linesList []string) (number int) {
	l := j
	r := j
	for common.IsDigit(linesList[i][l]) || common.IsDigit(linesList[i][r]) {
		if common.IsDigit(linesList[i][l]) {
			l -= 1
		}
		if common.IsDigit(linesList[i][r]) {
			r += 1
		}
	}
	var err error
	number, err = strconv.Atoi(string(linesList[i][l+1:r]))
	common.ErrorCheck(err)
	return number
}

func getGearsRatio(i, j int, linesList []string) (isGear bool, gearsRatio int) {
	gearsRatio = 1
	var topNumbers int
	if common.IsDigit(linesList[i+1][j]) {
		topNumbers = 1
		gearsRatio *= getNumber(i+1, j, linesList)
	} else {
		if common.IsDigit(linesList[i+1][j-1]) {
			topNumbers += 1
			gearsRatio *= getNumber(i+1, j-1, linesList)
		}
		if common.IsDigit((linesList[i+1][j+1])) {
			topNumbers += 1
			gearsRatio *= getNumber(i+1, j+1, linesList)
		}
	}

	var sideNumbers int
	if common.IsDigit(linesList[i][j-1]) {
		sideNumbers += 1
		gearsRatio *= getNumber(i, j-1, linesList)
	}
	if common.IsDigit(linesList[i][j+1]) {
		sideNumbers += 1
		gearsRatio *= getNumber(i, j+1, linesList)
	}

	var bottomNumbers int
	if common.IsDigit(linesList[i-1][j]) {
		bottomNumbers = 1
		gearsRatio *= getNumber(i-1, j, linesList)
	} else {
		if common.IsDigit(linesList[i-1][j-1]) {
			bottomNumbers += 1
			gearsRatio *= getNumber(i-1, j-1, linesList)
		}
		if common.IsDigit((linesList[i-1][j+1])) {
			bottomNumbers += 1
			gearsRatio *= getNumber(i-1, j+1, linesList)
		}
	}

	if topNumbers + sideNumbers + bottomNumbers == 2 {
		isGear = true
	}
	return isGear, gearsRatio
}

func SecondPuzzle(linesList []string) int {
	answer := 0
	for i, line := range linesList {
		for j := 0; j < len(line); j++ {
			if string(line[j]) == "*" {
				isGear, gearsRatio := getGearsRatio(i, j, linesList)
				if isGear {
					answer += gearsRatio
				}
			}
		}
	}

	fmt.Println("First puzzle answer: ", answer)
	return answer
}