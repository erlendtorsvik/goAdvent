package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var valueMap = map[rune]string{
	'F': "0", 'B': "1", 'L': "0", 'R': "1",
}

func main() {
	// er bare å ta binære verdien basert på F, B, L, R

	f, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	seatIDMap := make(map[int]bool)

	highestSeatID := part1BitShift(f, seatIDMap)

	fmt.Println(part1A(f))
	fmt.Println(part1B(f))
	fmt.Println(highestSeatID)
	fmt.Println(part2(seatIDMap, highestSeatID))
}

func part2(seatIDMap map[int]bool, highestSeatID int) int {
	for i := 1; i < highestSeatID; i++ {
		if !seatIDMap[i] && seatIDMap[i+1] && seatIDMap[i-1] {
			return i
		}
	}
	panic("couldnt find seat id")
}

func part1BitShift(f []byte, seatIDMap map[int]bool) int {
	var highestSeatID int

	for v := range strings.SplitSeq(strings.TrimSpace(string(f)), "\n") {
		var seatID int

		for _, ch := range v {
			// gigabrain bit shift til venstre, så 101 -> 1010
			seatID <<= 1

			// gjør en OR med 1010 basert på rune, hvis B eller R:
			// 1010 OR
			// 0001 =
			// 1011
			if ch == 'B' || ch == 'R' {
				seatID |= 1
			}
		}

		seatIDMap[seatID] = true

		if seatID > highestSeatID {
			highestSeatID = seatID
		}
	}
	return highestSeatID
}

func part1A(f []byte) int64 {
	var highestSeatID int64

	for v := range strings.SplitSeq(strings.TrimSpace(string(f)), "\n") {
		var b strings.Builder
		b.Grow(len(v))

		for _, ch := range v {
			b.WriteString(valueMap[ch])
		}

		seatID, _ := strconv.ParseInt(b.String(), 2, 64)

		if seatID > highestSeatID {
			highestSeatID = seatID
		}
	}
	return highestSeatID
}

func part1B(f []byte) int64 {
	var highestSeatID int64

	for v := range strings.SplitSeq(strings.TrimSpace(string(f)), "\n") {
		row := v[0:7]
		column := v[7:10]

		var rowB strings.Builder
		rowB.Grow(len(row))

		var columnB strings.Builder
		columnB.Grow(len(column))

		for _, ch := range row {
			rowB.WriteString(valueMap[ch])
		}

		for _, ch := range column {
			columnB.WriteString(valueMap[ch])
		}

		rowValue, _ := strconv.ParseInt(rowB.String(), 2, 64)
		columnValue, _ := strconv.ParseInt(columnB.String(), 2, 64)

		seatID := rowValue*8 + columnValue

		if seatID > highestSeatID {
			highestSeatID = seatID
		}

	}
	return highestSeatID
}
