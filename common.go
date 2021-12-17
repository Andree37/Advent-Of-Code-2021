package advent

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type NumberStringInput struct {
	Key   string
	Value int
}

func ReadNumberStringInput(input string) []NumberStringInput {
	var arr []NumberStringInput
	f, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(f)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		keyvalue := strings.Split(scanner.Text(), " ")
		key := keyvalue[0]
		value, err := strconv.Atoi(keyvalue[1])
		if err != nil {
			log.Fatal(err)
		}

		line := NumberStringInput{
			Key: key, Value: value,
		}
		arr = append(arr, line)
	}
	return arr
}

func ReadNumberInput(input string) []int {
	var arr []int
	f, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(f)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		intVar, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		arr = append(arr, intVar)
	}
	return arr
}
