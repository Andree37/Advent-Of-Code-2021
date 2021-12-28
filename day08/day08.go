package main

import (
		"advent"
		"fmt"
)

func part1() {
		var _, outputs = advent.ReadStringInputOutput("day08/day08.txt")

		var count int

		for _, outputLine := range outputs {
				for _, output := range outputLine {
						if len(output) == 2 || len(output) == 3 || len(output) == 4 || len(output) == 7 {
								count++
						}
				}
		}

		fmt.Println("Part1 Answer is: ", count)
}

func main() {
		part1()
}
