package puzzle

import (
	"fmt"
	"strconv"

	"github.com/zebra-f/Advent-of-Code-2023/day3/common"
)

func checkPartNumber(i, start, end int, linesList []string) (isPart bool, partNumber int) {
	if !common.IsDigit(linesList[i-1][start-1]) && string(linesList[i-1][start-1]) != "." {
		isPart = true
	} else if !common.IsDigit(linesList[i][start-1]) && string(linesList[i][start-1]) != "." {
		isPart = true
	} else if !common.IsDigit(linesList[i+1][start-1]) && string(linesList[i+1][start-1]) != "." {
		isPart = true
	}

	for j := start; j <= end; j++ {
		if !common.IsDigit(linesList[i+1][j]) && string(linesList[i+1][j]) != "." {
			isPart = true
		} else if !common.IsDigit(linesList[i-1][j]) && string(linesList[i-1][j]) != "." {
			isPart = true
		}
		if isPart {
			break
		}
	}
	
	if !common.IsDigit(linesList[i-1][end+1]) && string(linesList[i-1][end+1]) != "." {
		isPart = true
	} else if !common.IsDigit(linesList[i][end+1]) && string(linesList[i][end+1]) != "." {
		isPart = true
	} else if !common.IsDigit(linesList[i+1][end+1]) && string(linesList[i+1][end+1]) != "." {
		isPart = true
	}

	if isPart {
		partNumber_, err := strconv.Atoi(linesList[i][start:end+1])
		common.ErrorCheck(err)
		partNumber = partNumber_
	}

	return isPart, partNumber	
}

func FirstPuzzle(linesList []string) int {
	answer := 0
	for i, line := range linesList {
		for j := 0; j < len(line); j++ {
			flag := false
			k := j
			for _, err := strconv.Atoi(string(line[k])); err == nil; {
				k += 1
				flag = true
 				_, err = strconv.Atoi(string(line[k]));
			}
			if flag {
				isPart, partNumber := checkPartNumber(i, j, k-1, linesList)
				if isPart {
					answer += partNumber
				}
			}
			j = k
		}
	}

	fmt.Println("First puzzle answer: ", answer)
	return answer
}