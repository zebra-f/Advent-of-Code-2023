package puzzle

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	Go "github.com/zebra-f/Advent-of-Code-2023"
	"github.com/zebra-f/Advent-of-Code-2023/day7/common"
)

func SecondPuzzle(lines [][]string) (answer int) {
 	typeMap := map[int][]string{
		5: {},  // Five of a kind
		4: {},  // Four of a kind
		32: {},  // Full house
		3: {},  // Three of a kind
		2: {},  // Two pair
		1: {},  // One pair
		0: {},  // High card
	}
	handBidMap := map[string]int{}
	for _, line := range lines {
		hand := line[0]
		hand = strings.Replace(hand, "C", "1", -1)
		
		bid, err := strconv.Atoi(line[1])
		Go.ErrorCheck(err)
		
		handBidMap[hand] = bid
		
		handSlice := strings.Split(hand, "")
		sort.Strings(handSlice)
		counter := 1
		jokerCounter := 0
		sliceCounter := make([]int, 6)
		for j, char := range handSlice {
			if j == 0 {
				if char == "1" {
					jokerCounter += 1
				}
				continue
			}
			if char == "1" {
				jokerCounter += 1
				continue
			}
			if j > 0 && handSlice[j-1] == "1" {
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
		
		if jokerCounter == 5 || jokerCounter == 4 || sliceCounter[5] == 1  {
			typeMap[5] = append(typeMap[5], hand)
		} else if sliceCounter[4] == 1 {  // four of kind
			if jokerCounter == 1 {
				typeMap[5] = append(typeMap[5], hand)
			} else {
				typeMap[4] = append(typeMap[4], hand)
			}
		} else if sliceCounter[3] == 1 {  // three of kind
			if jokerCounter == 2 {
				typeMap[5] = append(typeMap[5], hand)
			} else if jokerCounter == 1 {
				typeMap[4] = append(typeMap[4], hand)
			} else if sliceCounter[2] == 1 {
				typeMap[32] = append(typeMap[32], hand)
			} else {
				typeMap[3] = append(typeMap[3], hand)
			}
		} else if sliceCounter[2] == 2 {  // two pairs
			if jokerCounter == 1 {
				typeMap[32] = append(typeMap[32], hand)
			} else {
				typeMap[2] = append(typeMap[2], hand)
			}
		} else if sliceCounter[2] == 1 { // one pair
			if jokerCounter == 3 {
				typeMap[5] = append(typeMap[5], hand)
			} else if jokerCounter == 2 {
				typeMap[4] = append(typeMap[4], hand)
			} else if jokerCounter == 1 {
				typeMap[3] = append(typeMap[3], hand)
			} else {
				typeMap[1] = append(typeMap[1], hand)
			}
		} else {  // high card
			if jokerCounter == 3 {
				typeMap[4] = append(typeMap[4], hand)
			} else if jokerCounter == 2 {
				typeMap[3] = append(typeMap[3], hand)
			} else if jokerCounter == 1 {
				typeMap[1] = append(typeMap[1], hand)
			} else {
				typeMap[0] = append(typeMap[0], hand)
			}
		}
	}

	answer = common.GetTotalWinnings(typeMap, handBidMap)
	fmt.Println("Second puzzle answer: ", answer)
	return answer
}