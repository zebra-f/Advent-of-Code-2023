package common

import "strconv"

func IsDigit(char byte) bool {
	if _, err := strconv.Atoi(string(char)); err == nil {
		return true
	}
	return false
}

