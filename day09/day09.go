package main

import (
	"advent"
	"fmt"
)

func part1() {
	boards := advent.ReadBoards("day09/day09.txt", 100)
	var foundMin []int
	for i := 0; i < len(boards.Rows); i++ {
		for j := 0; j < len(boards.Rows[i]); j++ {
			checkMin := 0
			// check for above value
			above := i - 1
			if above >= 0 {
				if boards.Rows[above][j] > boards.Rows[i][j] {
					checkMin++
				}
			} else {
				checkMin++
			}

			// check for below value
			below := i + 1
			if below < len(boards.Rows) {
				if boards.Rows[below][j] > boards.Rows[i][j] {
					checkMin++
				}
			} else {
				checkMin++
			}

			// check for left value
			left := j - 1
			if left >= 0 {
				if boards.Rows[i][left] > boards.Rows[i][j] {
					checkMin++
				}
			} else {
				checkMin++
			}

			// check for right value
			right := j + 1
			if right < len(boards.Rows[i]) {
				if boards.Rows[i][right] > boards.Rows[i][j] {
					checkMin++
				}
			} else {
				checkMin++
			}

			// check if its overall minimum
			if checkMin == 4 {
				foundMin = append(foundMin, boards.Rows[i][j])
			}
		}
	}
	fmt.Println(foundMin)
	// calculate risk level
	var sum int
	for _, min := range foundMin {
		sum += min + 1
	}

	fmt.Println("Part1 Answer is: ", sum)
}

func part2() {
	boards := advent.ReadBoards("day09/day09.txt", 100)

	fmt.Println("Part2 Answer is: ", boards)
}

func main() {
	part1()
	part2()
}
