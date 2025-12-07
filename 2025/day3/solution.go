package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	f, err := os.ReadFile("input.txt")
	if err != nil {
		panic("naai")
	}

	fmt.Println("part1", part1A(f))
	fmt.Println("part1B", part1and2(f, 2))
	fmt.Println("part2", part1and2(f, 12))
}

func part1A(f []byte) int {
	defer timeTrack(time.Now(), "part1")
	totalJoltage := 0

	for bank := range strings.SplitSeq(strings.TrimSpace(string(f)), "\n") {
		maxJoltage := 0

		if strings.Count(bank, "9") >= 2 {
			maxJoltage = 99
			totalJoltage += maxJoltage
			continue
		}

		for i := 0; i < len(bank); i++ {
			for j := i + 1; j < len(bank); j++ {

				first := bank[i] - '0'
				second := bank[j] - '0'

				value := int(first)*10 + int(second)

				if maxJoltage < value {
					maxJoltage = value
				}
			}
		}

		totalJoltage += maxJoltage

	}

	return totalJoltage
}

func part1and2(f []byte, n int) int {
	defer timeTrack(time.Now(), "part1B")
	totalJoltage := 0
	for bank := range strings.SplitSeq(strings.TrimSpace(string(f)), "\n") {
		var maxJoltage int

		remaining := len(bank) - n
		stack := make([]int, 0, len(bank))

		for _, battery := range bank {
			for len(stack) > 0 && remaining > 0 && stack[len(stack)-1] < int(battery-'0') {
				stack = stack[:len(stack)-1]
				remaining--
			}
			stack = append(stack, int(battery-'0'))
		}

		if len(stack) > n {
			stack = stack[:n]
		}

		for _, battery := range stack {
			maxJoltage = maxJoltage*10 + battery
		}

		totalJoltage += maxJoltage

	}

	return totalJoltage
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %s\n", name, elapsed)
}
