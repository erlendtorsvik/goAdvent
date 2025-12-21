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

func createNodeGrid(f []byte) [][]*node {
	lines := strings.Split(strings.TrimSpace(string(f)), "\n")

	height := len(lines)
	width := len(lines[0])

	grid := make([][]*node, height)

	for i, v := range lines {
		grid[i] = make([]*node, width)

		for j, c := range v {
			var n node
			n.value = valueMap[c]
			grid[i][j] = &n
		}
	}

	return grid
}

func linkAndCalcNodeGrid(grid [][]*node, rm bool) int {
	validNodes := 0
	height := len(grid)
	width := len(grid[0])

	for r := range height {
		for c := range width {
			score := 0
			n := grid[r][c]

			if n.value == 0 {
				continue
			}

			if r > 0 {
				n.up = grid[r-1][c]
				score += n.up.value
				if c > 0 {
					n.upleft = grid[r-1][c-1]
					score += n.upleft.value
				}
				if c < width-1 {
					n.upright = grid[r-1][c+1]
					score += n.upright.value
				}
			}

			if r < height-1 {
				n.down = grid[r+1][c]
				score += n.down.value
				if c > 0 {
					n.downleft = grid[r+1][c-1]
					score += n.downleft.value
				}
				if c < width-1 {
					n.downright = grid[r+1][c+1]
					score += n.downright.value
				}
			}

			if c > 0 {
				n.left = grid[r][c-1]
				score += n.left.value
			}

			if c < width-1 {
				n.right = grid[r][c+1]
				score += n.right.value
			}

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

	return linkAndCalcNodeGrid(grid, false)
}

func part2(f []byte) int {
	defer timeTrack(time.Now(), "part1")
	sum := 0

	grid := createNodeGrid(f)

	for {
		add := linkAndCalcNodeGrid(grid, true)
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
