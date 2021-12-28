package main

import (
	"advent"
	"fmt"
)

func part1() {
	input := advent.ReadNumberStringInput("day02/day02.txt")

	var verticalMov int
	var horizontalMov int

	for _, dict := range input {
		switch dict.Key {
		case "up":
			verticalMov -= dict.Value
		case "down":
			verticalMov += dict.Value
		case "forward":
			horizontalMov += dict.Value
		}
	}

	fmt.Println("Part 1 answer is: ", verticalMov*horizontalMov)
}

func part2() {
	input := advent.ReadNumberStringInput("day02/day02.txt")

	var aim int
	var horizontalMov int
	var verticalMov int

	for _, dict := range input {
		switch dict.Key {
		case "up":
			aim -= dict.Value
		case "down":
			aim += dict.Value
		case "forward":
			horizontalMov += dict.Value
			verticalMov += aim * dict.Value
		}
	}

	fmt.Println("Part 2 answer is: ", verticalMov*horizontalMov)
}

func main() {
	part1()
	part2()
}
