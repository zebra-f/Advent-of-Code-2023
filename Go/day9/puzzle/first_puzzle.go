package puzzle

import "fmt"

func FirstPuzzle(lines [][]int) (answer int) {

	var getHistoryNextValue func ([]int) int
	getHistoryNextValue = func (currHistoryLine []int) int {
		var lastValue int
		var zeroCounter int
		nextHistoryLine := []int{} 	
		for i := 1; i < len(currHistoryLine); i++ {
			currValue := currHistoryLine[i]
			prevValue := currHistoryLine[i-1]
			nextHistoryLineValue := currValue - prevValue

			if nextHistoryLineValue == 0 {
				zeroCounter += 1
			}
			
			nextHistoryLine = append(nextHistoryLine, nextHistoryLineValue)
			lastValue = currValue
		}

		if len(nextHistoryLine) == zeroCounter {
			return lastValue

		}
		return lastValue + getHistoryNextValue(nextHistoryLine)
	}

	for _, line := range lines{
		answer += getHistoryNextValue(line)
	}

	fmt.Println("First puzzle answer: ", answer)
	return answer
}