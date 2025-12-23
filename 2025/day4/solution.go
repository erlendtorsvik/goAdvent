package main

import (
	"bytes"
	"fmt"
	"os"
	"time"
)

var valueMap = map[rune]int{
	'@': 1, '.': 0,
}

func main() {
	f, err := os.ReadFile("input.txt")
	if err != nil {
		panic("naai")
	}

	fmt.Println("part1", part1(f))
	fmt.Println("part2", part2(f))
}

func createNodeGrid(f []byte) [][]int {
	lines := bytes.Split(bytes.TrimSpace(f), []byte{'\n'})

	height := len(lines)
	width := len(lines[0])

	grid := make([][]int, height+1)
	grid[0] = make([]int, width+2)

	for i := 1; i < len(grid); i++ {
		grid[i] = make([]int, width+2)

		for j := 1; j < len(grid[0])-1; j++ {
			c := rune(lines[i-1][j-1])
			n := valueMap[c]
			grid[i][j] = n
		}

	}
	grid = append(grid, make([]int, width+2))

	return grid
}

func checkValidNodes(grid [][]int, rm bool) int {
	validNodes := 0

	height := len(grid)
	width := len(grid[0])
	for r := 1; r < height-1; r++ {
		for c := 1; c < width-1; c++ {
			if grid[r][c] == 0 {
				continue
			}

			score := 0

			score += grid[r-1][c]
			score += grid[r-1][c-1]
			score += grid[r-1][c+1]
			score += grid[r+1][c]
			score += grid[r+1][c-1]
			score += grid[r+1][c+1]
			score += grid[r][c-1]
			score += grid[r][c+1]

			if score < 4 {
				validNodes++
				if rm {
					grid[r][c] = 0
				}
			}
		}
	}

	return validNodes
}

func part1(f []byte) int {
	defer timeTrack(time.Now(), "part1")

	grid := createNodeGrid(f)

	return checkValidNodes(grid, false)
}

func part2(f []byte) int {
	defer timeTrack(time.Now(), "part2")
	sum := 0

	grid := createNodeGrid(f)

	for {
		add := checkValidNodes(grid, true)
		sum += add

		if add == 0 {
			break
		}
	}

	return sum
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %s\n", name, elapsed)
}
