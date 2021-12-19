package main

import (
	"advent"
	"fmt"
)

func part1() {
	lines := advent.ReadLines("day05/day05.txt")

	board := make([][]uint8, 990)
	for i := range board {
		board[i] = make([]uint8, 990)
	}

	var sum int
	for _, line := range lines {
		for _, point := range line.Points {
			fmt.Println(point.X, point.Y)
			board[point.X][point.Y]++
			// check for overlap
			if board[point.X][point.Y] == 2 {
				sum++
			}
		}
	}

	fmt.Println("Parts Answer is: ", sum)
}

func main() {
	part1()
}
