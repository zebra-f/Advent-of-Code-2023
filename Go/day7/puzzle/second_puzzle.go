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
		// these are current replacement (used also for the first puzzle): 
		// line = strings.Replace(line, "T", "B", -1)
		// line = strings.Replace(line, "J", "C", -1)
		// line = strings.Replace(line, "Q", "D", -1)
		// line = strings.Replace(line, "K", "E", -1)
		// line = strings.Replace(line, "A", "F", -1), they are ordered from lowest ("T"-"B") to highest value ("A"-"F") 
		// replace "C" (Joker) with "1"
		hand = strings.Replace(hand, "C", "1", -1)
		
		bid, err := strconv.Atoi(line[1])
		Go.ErrorCheck(err)
		
		// simple mapping of hand and it's value
		handBidMap[hand] = bid
		
		handSlice := strings.Split(hand, "")
		// sort, so for example slice of "51F51" becomse [1 1 5 5 F]
		sort.Strings(handSlice)
		counter := 1
		jokerCounter := 0
		// a little bit confusing
		// [
		// 	0 stores: nothing,
		// 	1 number of `high cards` either 1 or 5,
		// 	2 number of `two pairs`- either 1 or 2,
		// 	3 number Three of a kind (1 or 0) etc.
		// 	4
		// 	5
		// ]
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
			if sliceCounter[5] == 1 {
				fmt.Println(hand)
			}
			typeMap[5] = append(typeMap[5], hand)
		}  else if sliceCounter[4] == 1 {  // four of kind
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
			if jokerCounter >= 2 {
				typeMap[4] = append(typeMap[4], hand)
			} else if jokerCounter == 1 {
				typeMap[3] = append(typeMap[3], hand)
			} else {
				typeMap[1] = append(typeMap[1], hand)
			}
		} else {  // uniqe cards
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