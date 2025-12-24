package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main() {
	f, err := os.ReadFile("input.txt")
	if err != nil {
		panic("naai")
	}

	fmt.Println(part1(f))
	fmt.Println(part2(f))
}

func part1(f []byte) int {
	defer timeTrack(time.Now(), "part1")
	validIDCount := 0
	idRanges, ids, _ := strings.Cut(string(f), "\n\n")

	var intRanges [][]int
	var intIds []int

	for idRange := range strings.SplitSeq(strings.TrimSpace(idRanges), "\n") {
		startID, endID, _ := strings.Cut(idRange, "-")
		startNumber, _ := strconv.Atoi(startID)
		endNumber, _ := strconv.Atoi(endID)
		rangeSlice := []int{startNumber, endNumber}
		intRanges = append(intRanges, rangeSlice)
	}

	for id := range strings.SplitSeq(strings.TrimSpace(ids), "\n") {
		intID, _ := strconv.Atoi(id)
		intIds = append(intIds, intID)
	}

	for _, id := range intIds {
		for _, v := range intRanges {
			if id >= v[0] && id <= v[1] {
				validIDCount++
				break
			}
		}
	}

	return validIDCount
}

func part2(f []byte) int {
	defer timeTrack(time.Now(), "part2")
	idRanges, _, _ := strings.Cut(string(f), "\n\n")
	var intRanges [][]int
	validIDCount := 0

	for idRange := range strings.SplitSeq(strings.TrimSpace(idRanges), "\n") {
		startID, endID, _ := strings.Cut(idRange, "-")
		startNumber, _ := strconv.Atoi(startID)
		endNumber, _ := strconv.Atoi(endID)
		rangeSlice := []int{startNumber, endNumber}
		intRanges = append(intRanges, rangeSlice)
	}

	sort.Slice(intRanges, func(i, j int) bool {
		return intRanges[i][0] < intRanges[j][0]
	})

	merged := make([][]int, 0, len(intRanges))
	cur := []int{intRanges[0][0], intRanges[0][1]}

	for i := 1; i < len(intRanges); i++ {
		nxt := intRanges[i]
		if nxt[0] <= cur[1] {
			if nxt[1] > cur[1] {
				cur[1] = nxt[1]
			}
		} else {
			merged = append(merged, cur)
			cur = []int{nxt[0], nxt[1]}
		}
	}

	merged = append(merged, cur)

	for _, r := range merged {
		validIDCount += r[1] - r[0] + 1
	}

	return validIDCount
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %s\n", name, elapsed)
}
