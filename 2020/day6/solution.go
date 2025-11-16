package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	f, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	sum := 0
	sum2 := 0

	for group := range strings.SplitSeq(strings.TrimSpace(string(f)), "\n\n") {
		sum += uniqueCharacterCount(group)
		sum2 += allYesCount(group)
	}

	fmt.Println(sum)
	fmt.Println(sum2)
}

func allYesCount(group string) int {
	persons := strings.Split(group, "\n")

	freq := make(map[rune]int)

	for _, person := range persons {
		for _, ch := range person {
			freq[ch]++
		}
	}

	allYes := 0

	for _, c := range freq {
		if c == len(persons) {
			allYes++
		}
	}

	return allYes
}

func uniqueCharacterCount(str string) int {
	charSet := make(map[rune]bool)

	for _, char := range str {
		if char != '\n' {
			charSet[char] = true
		}
	}

	return len(charSet)
}
