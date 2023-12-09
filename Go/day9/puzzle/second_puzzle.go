package puzzle

import "fmt"

func SecondPuzzle(lines [][]int) (answer int) {

	var getHistoryNextValue func ([]int) int
	getHistoryNextValue = func (currHistoryLine []int) int {
		var firstValue int = currHistoryLine[0]
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
		}

		if len(nextHistoryLine) == zeroCounter {
			return firstValue

		}
		return firstValue - getHistoryNextValue(nextHistoryLine)
	}

	for _, line := range lines{
		answer += getHistoryNextValue(line)
	}

	fmt.Println("Second puzzle answer: ", answer)
	return answer
}