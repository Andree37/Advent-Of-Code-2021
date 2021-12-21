package main

import (
	"advent"
	"fmt"
	"math"
	"sort"
)

func part1() {
	input := advent.ReadNumberInlineInput("day07/day07.txt")

	sort.Ints(input)

	median := input[len(input)/2]

	var fuel float64
	for _, s := range input {
		fuel += math.Abs(float64(s - median))
	}

	fmt.Println("Part1 Answer is: ", fuel)
}

func part2() {
	input := advent.ReadNumberInlineInput("day07/day07.txt")

	sort.Ints(input)

	var sum int
	for _, s := range input {
		sum += s
	}

	mean := float64(sum) / float64(len(input))
	mean = math.Floor(mean)
	// calculate fuel
	var fuel int
	for _, s := range input {
		effort := int(math.Abs(float64(s) - mean))
		fuel += (effort * (effort + 1)) / 2
	}

	fmt.Println("Part2 Answer is: ", fuel)
}

func main() {
	//part1()
	part2()
}
