package common

import (
	"strconv"
	"strings"
)

func MaxColors(gameSet []string) (red, green, blue int) {
	for _, setStr := range gameSet {
		setList := strings.Split(setStr, ",")
		for _, valColor := range setList {
			valColorLen := len(valColor)
			color := valColor[valColorLen-1:]
			val, err := strconv.Atoi(valColor[:valColorLen-1])
			if err != nil {
				panic(err)
			}
			if color == "r" {
				red = max(red, val)
			} else if color == "g" {
				green = max(green, val)
			} else {
				blue = max(blue, val)
			}
		}
	}
	return red, green, blue
}