package advent

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

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
