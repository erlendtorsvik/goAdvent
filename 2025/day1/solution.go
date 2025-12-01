package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.ReadFile("input.txt")
	if err != nil {
		panic("naai")
	}

	dialNumber := 50

	password := part1(f, dialNumber)

	fmt.Println("password:", password)
}

func part1(f []byte, dialNumber int) int {
	password := 0

	for v := range strings.SplitSeq(strings.TrimSpace(string(f)), "\n") {
		direction := v[0]
		number, err := strconv.Atoi(v[1:])
		if err != nil {
			panic("klarte ikkje konvertere??")
		}

		if direction == 'L' {
			dialNumber -= number
		} else {
			dialNumber += number
		}

		if dialNumber%100 == 0 {
			password++
		}

	}

	return password
}
