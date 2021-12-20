package main

import (
	"advent"
	"fmt"
)

type Fish struct {
	age int
}

func simulateDay(fishes *[]Fish) {
	var initLen = len(*fishes)
	for i := 0; i < initLen; i++ {
		switch (*fishes)[i].age {
		case 0:
			(*fishes)[i].age = 6
			*fishes = append(*fishes, Fish{age: 8})
		default:
			(*fishes)[i].age--
		}
	}
}

func part1() {
	input := advent.ReadNumberInlineInput("day06/day06.txt")
	var fishes []Fish
	var simulatedDays = 80

	for _, s := range input {
		fishes = append(fishes, Fish{age: s})
	}

	for i := 0; i < simulatedDays; i++ {
		simulateDay(&fishes)
	}

	fmt.Println("Part1 Answer is: ", len(fishes))
}

func simulateBucketDay(bucket *[]int) {
	var fathersN = (*bucket)[0]
	*bucket = (*bucket)[1:]
	*bucket = append(*bucket, fathersN)
	(*bucket)[6] += fathersN
}

func part2() {
	input := advent.ReadNumberInlineInput("day06/day06.txt")
	var buckets = []int{0, 0, 0, 0, 0, 0, 0, 0, 0}

	var simulatedDays = 256

	for _, s := range input {
		buckets[s]++
	}

	for i := 0; i < simulatedDays; i++ {
		simulateBucketDay(&buckets)
	}

	fmt.Println(buckets)

	var sum = 0
	for _, bucket := range buckets {
		sum += bucket
	}
	fmt.Println("Part2 Answer is: ", sum)
}

func main() {
	//part1()
	part2()
}
