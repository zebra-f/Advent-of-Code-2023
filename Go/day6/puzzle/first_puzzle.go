package puzzle

import (
	"fmt"

	"github.com/zebra-f/Advent-of-Code-2023/day6/common"
)

func FirstPuzzle(times, distances []int) (answer int) {
	answer = 1
	for i, time := range times {
		distance := distances[i]
		lowerBound := common.LowerBoundBinarySearch(0, time, distance + 1)
		upperBound := common.UpperBoundBinarySearch(0, time, distance + 1)
		answer *= upperBound - lowerBound
	}	

	fmt.Println("First puzzle answer: ", answer)
	return answer
}