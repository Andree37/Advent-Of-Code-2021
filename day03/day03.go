package main

import (
		"advent"
		"fmt"
		"log"
		"math"
		"strconv"
)

type bitCounter struct {
		bits []int
		sum  int
}

func digit(num, place int) int {
		r := num % int(math.Pow(10, float64(place)))
		return r / int(math.Pow(10, float64(place-1)))
}

func part1() {
		input := advent.ReadNumberInput("day03/day03.txt")

		var bits = []bitCounter{
				bitCounter{
						sum: 0, bits: []int{},
				},
				{sum: 0, bits: []int{}},
				{sum: 0, bits: []int{}},
				{sum: 0, bits: []int{}},
				{sum: 0, bits: []int{}},
				{sum: 0, bits: []int{}},
				{sum: 0, bits: []int{}},
				{sum: 0, bits: []int{}},
				{sum: 0, bits: []int{}},
				{sum: 0, bits: []int{}},
				{sum: 0, bits: []int{}},
				{sum: 0, bits: []int{}},
		}

		for _, in := range input {
				for j := 0; j < 12; j++ {
						d := digit(in, j+1)
						bits[j].bits = append(bits[j].bits, d)
						if d == 0 {
								bits[j].sum += -1
						} else {
								bits[j].sum += d
						}
				}
		}

		gammaBin := ""
		epsilonBin := ""

		for i := 11; i >= 0; i-- {
				if bits[i].sum >= 0 {
						gammaBin += strconv.FormatInt(int64(1), 10)
						epsilonBin += strconv.FormatInt(int64(0), 10)
				} else {
						gammaBin += strconv.FormatInt(int64(0), 10)
						epsilonBin += strconv.FormatInt(int64(1), 10)
				}
		}

		gamma, err := strconv.ParseInt(gammaBin, 2, 64)
		if err != nil {
				log.Fatal(err)
		}

		epsilon, err := strconv.ParseInt(epsilonBin, 2, 64)
		fmt.Println("Part1 Answer is: ", gamma*epsilon)
}

func part2() {
		input := advent.ReadStringInput("day03/day03.txt")

		var filteredInput []string
		sum := 0
		for i, s := range input {
				if s[i] == '0' {
						sum += -1
				} else {
						sum += 1
				}
		}

		for i, s := range input {
				if sum > 0 && s[i] == '1' {

				}
		}
}

func main() {
		part1()
}
