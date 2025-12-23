package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

type node struct {
	value     int
	up        *node
	upright   *node
	upleft    *node
	down      *node
	downright *node
	downleft  *node
	right     *node
	left      *node
}

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

func createNodeGrid(f []byte) [][]node {
	lines := strings.Split(strings.TrimSpace(string(f)), "\n")

	height := len(lines)
	width := len(lines[0])

	grid := make([][]node, height+1)
	grid[0] = make([]node, width+2)

	for i := 1; i < len(grid); i++ {
		grid[i] = make([]node, width+2)

		for j := 1; j < len(grid[0])-1; j++ {
			c := rune(lines[i-1][j-1])
			var n node
			n.value = valueMap[c]
			grid[i][j] = n
		}

	}
	grid = append(grid, make([]node, width+2))

	linkNodes(grid)

	return grid
}

func linkNodes(grid [][]node) {
	for r := 1; r < len(grid)-1; r++ {
		for c := 1; c < len(grid[0])-1; c++ {
			n := &grid[r][c]

			n.up = &grid[r-1][c]
			n.upleft = &grid[r-1][c-1]
			n.upright = &grid[r-1][c+1]
			n.down = &grid[r+1][c]
			n.downleft = &grid[r+1][c-1]
			n.downright = &grid[r+1][c+1]
			n.left = &grid[r][c-1]
			n.right = &grid[r][c+1]

		}
	}
}

func checkValidNodes(grid [][]node, rm bool) int {
	validNodes := 0

	height := len(grid)
	width := len(grid[0])
	for r := 1; r < height-1; r++ {
		for c := 1; c < width-1; c++ {
			n := &grid[r][c]
			if n.value == 0 {
				continue
			}

			score := 0

			score += n.up.value
			score += n.upleft.value
			score += n.upright.value
			score += n.down.value
			score += n.downleft.value
			score += n.downright.value
			score += n.left.value
			score += n.right.value

			if score < 4 {
				validNodes++
				if rm {
					n.value = 0
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
