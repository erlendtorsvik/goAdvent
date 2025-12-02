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
	password2 := part2(f, dialNumber)

	fmt.Println("password part1:", password)
	fmt.Println("password part2:", password2)
}

func part1(f []byte, dialNumber int) int {
	var password int

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

func part2(f []byte, dialNumber int) int {
	var password int

	for v := range strings.SplitSeq(strings.TrimSpace(string(f)), "\n") {
		dialNumber %= 100

		// Go fant ut en dritkul ting at de beholder fortegnet :))))))))))))))))))))))))
		// Så -18 modulo 100 = -18, ikkje 82
		// VERY COOL!!!!!
		if dialNumber < 0 {
			dialNumber += 100
		}

		direction := v[0]
		number, err := strconv.Atoi(v[1:])
		if err != nil {
			panic("klarte ikkje konvertere??")
		}

		if number == 0 {
			continue
		}

		if direction == 'L' {
			number *= -1
		}

		adder := 100
		if dialNumber == 0 {
			adder = 0
		}

		dialNumber += number

		var turns int

		if dialNumber < 1 {
			turns = (dialNumber*-1 + adder) / 100
		} else {
			turns = dialNumber / 100
		}

		password += turns

	}

	return password
}
