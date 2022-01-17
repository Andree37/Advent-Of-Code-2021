package main

import (
	"advent"
	"fmt"
)

func part1() {
	boards := advent.ReadBoards("day09/day09.txt", 10)
	for _, row := range boards.Rows {
		fmt.Println(row)
	}
}

func main() {
	part1()
}
