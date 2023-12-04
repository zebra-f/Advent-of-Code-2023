package puzzle

import (
	"fmt"

	"github.com/zebra-f/Advent-of-Code-2023/day4/common"
)

func SecondPuzzle(numsList [][][]int) (answer float64) {
	winningCardMap := map[int]int{
		1: 1, 
	}
	
	for card, scratchcard := range numsList {
		card = card + 1
		if card > 1 {
			winningCardMap[card] += 1
		}
		
		winningNums := scratchcard[0]
		elfsNums := scratchcard[1]
		winningNumsSet := common.MakeSet(winningNums)
		var nextCard int = card
		for _, num := range elfsNums {
			if _, ok := winningNumsSet[num]; ok {
				nextCard += 1
				winningCardMap[nextCard] += 1 * winningCardMap[card]
			}
		}
	}

	for key := range winningCardMap {
		answer += float64(winningCardMap[key])
	}

	fmt.Println("Second puzzle answer: ", int(answer))
	return answer
}