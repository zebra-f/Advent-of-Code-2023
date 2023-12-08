package puzzle

import "fmt"

func FirstPuzzle(startPosition, endPosition, instructions string, instructionsMap map[string][]string) (answer int) {
	var steps int 	
	var currPosition string = startPosition
	for i := 0; true; i++ {
		if i >= len(instructions) {
			i = 0
		}

		if currPosition == endPosition {
			break
		}

		left, right := instructionsMap[currPosition][0], instructionsMap[currPosition][1]
		if instructions[i] == 'L' {
			currPosition = left
		} else {
			currPosition = right
		}

		steps += 1
	}

	answer += steps
	fmt.Println("First puzzle answer: ", answer)
	return answer
}