package puzzle

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	Go "github.com/zebra-f/Advent-of-Code-2023"
	"github.com/zebra-f/Advent-of-Code-2023/day7/common"
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

	answer = common.GetTotalWinnings(typeMap, handBidMap)
	fmt.Println("First puzzle answer: ", answer)
	return answer
}