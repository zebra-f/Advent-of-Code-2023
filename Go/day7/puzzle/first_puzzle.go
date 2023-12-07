package puzzle

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	Go "github.com/zebra-f/Advent-of-Code-2023"
)

func FirstPuzzle(lines [][]string) (answer int) {
 	typeMap := map[int][]string{
		5: {},
		4: {},
		32: {},
		3: {},
		2: {},
		1: {},
		0: {},
	}
	handBidMap := map[string]int{}

	for _, line := range lines {
		hand := line[0]
		bid, err := strconv.Atoi(line[1])
		Go.ErrorCheck(err)
		
		handBidMap[hand] = bid
		
		handSlice := strings.Split(hand, "")
		sort.Strings(handSlice)

		counter := 1
		sliceCounter := make([]int, 6)
		for j, char := range handSlice {
			if j == 0 {
				continue
			}
			if char == handSlice[j-1] {
				counter += 1	
			} else {
				sliceCounter[counter] += 1
				counter = 1
			}
		}
		sliceCounter[counter] += 1

		flag := false
		for j := 5; j > 1; j-- {
			if sliceCounter[j] > 0 {
				flag = true

				if j == 2 {
					typeMap[sliceCounter[2]] = append(typeMap[sliceCounter[2]], hand)
					break
				}
				if j == 3 && sliceCounter[2] > 0 {
					typeMap[32] = append(typeMap[32], hand)
					break
				}

				typeMap[j] = append(typeMap[j], hand)
				break
			}
		}
		if !flag {
			typeMap[0] = append(typeMap[0], hand)
		}
	}

	rank := 1
	for i := 0; i <= 5; i++ {
		sort.Strings(typeMap[i])
		for j := 0; j < len(typeMap[i]); j++ {
			answer += rank * handBidMap[typeMap[i][j]]
			rank += 1
		}
		if i == 3 {
			sort.Strings(typeMap[32])
			for j := 0; j < len(typeMap[32]); j++ {

				answer += rank * handBidMap[typeMap[32][j]]
				rank += 1
			}
		}
	}

	fmt.Println("First puzzle answer: ", answer)
	return answer
}