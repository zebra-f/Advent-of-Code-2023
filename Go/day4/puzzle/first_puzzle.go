package puzzle

import (
	"fmt"
	"math"

	"github.com/zebra-f/Advent-of-Code-2023/day4/common"
)

func FirstPuzzle(numsList [][][]int) (answer float64) {
	for _, scratchcard := range numsList {
		winningNums := scratchcard[0]
		elfsNums := scratchcard[1]

		winningNumsSet := common.MakeSet(winningNums)
		var counter int
		for _, num := range elfsNums {
			if _, ok := winningNumsSet[num]; ok {
				counter += 1
			}
		}
		if counter != 0 {
			answer += math.Pow(2, float64(counter-1))
		}
	}

	fmt.Println("First puzzle answer: ", answer)
	return answer
}