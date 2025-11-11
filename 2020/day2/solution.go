package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Entry struct {
	Min      int
	Max      int
	Letter   rune
	Password string
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := f.Close(); err != nil {
			panic(err)
		}
	}()

	scanner := bufio.NewScanner(f)

	var entries []Entry

	for scanner.Scan() {
		line := scanner.Text()
		e, err := parseLine(line)
		if err != nil {
			panic(err)
		}
		entries = append(entries, e)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	validPasswords := validPasswordCount(entries)

	validPasswords2 := validPasswords2(entries)

	fmt.Println(len(entries), "entries loaded")
	fmt.Println("valid passwords: ", validPasswords)
	fmt.Println("valid passwords part two: ", validPasswords2)
}

func validPasswords2(entries []Entry) int {
	var count int

	for _, v := range entries {

		minLetter := v.Password[v.Min-1]
		maxLetter := v.Password[v.Max-1]

		minPosMatch := minLetter == byte(v.Letter)
		maxPosMatch := maxLetter == byte(v.Letter)

		if minPosMatch != maxPosMatch {
			count++
		}
	}
	return count
}

func validPasswordCount(entries []Entry) int {
	var count int

	for _, v := range entries {
		letter := string(v.Letter)
		letterAmount := strings.Count(v.Password, letter)

		if letterAmount >= v.Min && letterAmount <= v.Max {
			count++
		}

	}
	return count
}

func parseLine(line string) (Entry, error) {
	var e Entry
	_, err := fmt.Sscanf(line, "%d-%d %c: %s", &e.Min, &e.Max, &e.Letter, &e.Password)
	return e, err
}
