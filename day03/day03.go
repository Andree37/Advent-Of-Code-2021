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
		if err != nil {
				log.Fatal(err)
		}
		fmt.Println("Part1 Answer is: ", gamma*epsilon)
}

func part2() {
		input := advent.ReadStringInput("day03/day03.txt")

		var filterOx = input
		var filterCo = input

		var oneArrOx []string
		var zerArrOx []string

		var oneArrCo []string
		var zerArrCo []string

		for i := 0; i < 12; i++ {
				for _, s := range filterOx {
						if s[i] == '0' {
								zerArrOx = append(zerArrOx, s)
						} else {
								oneArrOx = append(oneArrOx, s)
						}

				}

				for _, s := range filterCo {
						if s[i] == '0' {
								zerArrCo = append(zerArrCo, s)
						} else {
								oneArrCo = append(oneArrCo, s)
						}
				}

				if len(filterOx) > 1 {
						if len(zerArrOx) > len(oneArrOx) {
								filterOx = zerArrOx
						} else {
								filterOx = oneArrOx
						}
				}

				if len(filterCo) > 1 {
						if len(oneArrCo) < len(zerArrCo) {
								filterCo = oneArrCo
						} else {
								filterCo = zerArrCo
						}
				}

				zerArrOx = []string{}
				oneArrOx = []string{}

				oneArrCo = []string{}
				zerArrCo = []string{}

		}

		ox, err := strconv.ParseInt(filterOx[0], 2, 64)
		if err != nil {
				log.Fatal(err)
		}

		co, err := strconv.ParseInt(filterCo[0], 2, 64)
		if err != nil {
				log.Fatal(err)
		}

		fmt.Println("Part2 Answer is: ", ox*co)
}

func main() {
		part1()
		part2()
}
