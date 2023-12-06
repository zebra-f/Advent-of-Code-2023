package puzzle

import (
	"fmt"
	"strconv"

	Go "github.com/zebra-f/Advent-of-Code-2023"
	"github.com/zebra-f/Advent-of-Code-2023/day6/common"
)

func SecondPuzzle(times, distances []int) (answer int) {
	var timeStr string
	for _, time := range times {
		timeStr = timeStr + strconv.Itoa(time)
	}
	
	var distanceStr string
	for _, distance := range distances {
		distanceStr = distanceStr + strconv.Itoa(distance)
	}
	
	time, err := strconv.Atoi(timeStr)
	Go.ErrorCheck(err)
	distance, err := strconv.Atoi(distanceStr)
	Go.ErrorCheck(err)

	lowerBound := common.LowerBoundBinarySearch(0, time, distance + 1)
	upperBound := common.UpperBoundBinarySearch(0, time, distance + 1)
	answer = upperBound - lowerBound

	fmt.Println("Second puzzle answer: ", answer)
	return answer
}