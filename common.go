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

type Cell struct {
	Value int
	Found bool
}

type Board struct {
	Rows  [][]Cell
	Found bool
}

func ReadStringInput(input string) []string {
	var arr []string
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
		arr = append(arr, scanner.Text())
	}
	return arr
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

func ReadNumbersAndBoard(input string) ([]int, []Board) {
	var inputs []int
	var boards []Board

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

	board := Board{Found: false}
	var currentRow []Cell

	first := true
	for scanner.Scan() {
		line := scanner.Text()
		if first {
			var splitInput = strings.Split(line, ",")
			for _, s := range splitInput {
				intVar, err := strconv.Atoi(s)
				if err != nil {
					log.Fatal(err)
				}
				inputs = append(inputs, intVar)
			}
		} else {
			splitLine := strings.Split(line, " ")
			for _, l := range splitLine {
				if l != "" {
					intVar, err := strconv.Atoi(l)
					if err != nil {
						log.Fatal(err)
					}
					currentRow = append(currentRow, Cell{Found: false, Value: intVar})
				}
			}

			if len(currentRow) == 5 {
				board.Rows = append(board.Rows, currentRow)
				currentRow = []Cell{}
			}

			if len(board.Rows) == 5 {
				boards = append(boards, board)
				board = Board{}
			}
		}

		first = false
	}

	return inputs, boards
}
