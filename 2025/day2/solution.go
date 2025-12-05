package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/dlclark/regexp2"
)

var reRepeat = regexp2.MustCompile(`^(\d+)\1+$`, regexp2.None)

func main() {
	f, err := os.ReadFile("input.txt")
	if err != nil {
		panic("oops")
	}

	fmt.Println(part1(f))
	fmt.Println(part2(f))
}

func part1(f []byte) int {
	defer timeTrack(time.Now(), "part1")
	idsAdded := 0
	for idRange := range strings.SplitSeq(strings.TrimSpace(string(f)), ",") {
		ids := strings.SplitN(idRange, "-", 2)

		min, _ := strconv.Atoi(ids[0])
		max, _ := strconv.Atoi(ids[1])

		for i := min; i <= max; i++ {
			length := lenLoop(i)
			if length%2 != 0 {
				continue
			}

			s := strconv.Itoa(i)

			mid := len(s) / 2

			if s[:mid] == s[mid:] {
				idsAdded += i
			}

		}

	}
	return idsAdded
}

func part2(f []byte) int {
	defer timeTrack(time.Now(), "part2")
	idsAdded := 0
	for idRange := range strings.SplitSeq(strings.TrimSpace(string(f)), ",") {
		ids := strings.SplitN(idRange, "-", 2)

		min, _ := strconv.Atoi(ids[0])
		max, _ := strconv.Atoi(ids[1])

		for i := min; i <= max; i++ {

			s := strconv.Itoa(i)

			if match, _ := reRepeat.MatchString(s); match {
				idsAdded += i
			}

		}

	}
	return idsAdded
}

func lenLoop(i int) int {
	if i == 0 {
		return 1
	}
	count := 0
	for i != 0 {
		i /= 10
		count++
	}
	return count
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %s\n", name, elapsed)
}
