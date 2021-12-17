package main

import (
		"advent"
		"fmt"
		"math"
)

func part1() {
		input := advent.ReadNumberInput("day1/day1.txt")

		// calculate how many times a new max is calculated
		var prev = math.MaxInt
		var counter = 0

		for _, s := range input {
				if s > prev {
						counter++
				}
				prev = s
		}

		fmt.Println("Part1 Answer is: ", counter)
}

func part2() {
		// window is a 3 reading input interval
		// meaning that we would have three windows reading at the same time at most, each with their own inputs
		// we need to then sum the inputs from the windows when we reach completion

		windows := make([][]int, 4)

		input := advent.ReadNumberInput("day1/day1.txt")

		prev := math.MaxInt
		counter := 0

		for i, s := range input {
				for j := 0; j < len(windows); j++ {
						if (j <= i) && (j != (i+1)%4) {
								windows[j] = append(windows[j], s)
						}

						if len(windows[j]) == 3 {
								sum := 0
								for w := 0; w < 3; w++ {
										sum += windows[j][w]
								}

								if sum > prev {
										counter++
								}

								windows[j] = []int{}

								prev = sum
						}
				}
		}
		fmt.Println("Part2 Answer is: ", counter)
}

func main() {
		//part1()
		part2()
}
