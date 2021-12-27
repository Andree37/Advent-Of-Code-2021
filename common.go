package advent

import (
	"bufio"
	"log"
	"os"
	"sort"
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

type Point struct {
	X int
	Y int
}

type Line struct {
	Points []Point
}

func ReadStringInputOutput(input string) ([]string, []string) {
	var digits = ReadStringInput(input)

	var inputs []string
	var outputs []string

	for _, digit := range digits {
		var line = strings.Split(digit, "|")
		inputs = append(inputs, line[0])
		outputs = append(outputs, line[1])
	}

	return inputs, outputs
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

func ReadNumberInlineInput(input string) []int {
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
		line := scanner.Text()
		for _, s := range strings.Split(line, ",") {
			intVar, err := strconv.Atoi(s)
			if err != nil {
				log.Fatal(err)
			}
			arr = append(arr, intVar)
		}

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

func ReadLines(input string) []Line {
	lines := ReadStringInput(input)
	var result []Line

	for _, line := range lines {
		strPoints := strings.Split(line, " -> ")
		strStart := strings.Split(strPoints[0], ",")
		strEnd := strings.Split(strPoints[1], ",")

		xStart, _ := strconv.Atoi(strStart[0])
		yStart, _ := strconv.Atoi(strStart[1])

		xEnd, _ := strconv.Atoi(strEnd[0])
		yEnd, _ := strconv.Atoi(strEnd[1])

		startPoint := Point{X: xStart, Y: yStart}
		endPoint := Point{X: xEnd, Y: yEnd}

		points := []Point{startPoint, endPoint}

		sort.Slice(points, func(i, j int) bool {
			if points[i].X < points[j].X {
				return true
			} else if points[i].Y < points[j].Y && !(points[i].X > points[j].X) {
				return true
			} else {
				return false
			}
		})

		// find all intermediate points between start and end
		horiz := xStart != xEnd
		vert := yStart != yEnd

		var lastInsert int

		if horiz && vert {
			for i := points[0].X + 1; i < points[1].X; i++ {
				if points[0].X > points[1].X {
					points = append(points, Point{X: points[0].X - i, Y: points[1].X + i})
				} else if points[0].X < points[1].X {
					if points[0].Y > points[1].Y {
						var upper int
						var lwr int
						if points[0].Y > points[1].Y {
							upper = points[0].Y
							lwr = points[1].Y
						} else {
							upper = points[1].Y
							lwr = points[0].Y
						}
						if points[0].X == points[1].Y || points[0].Y == points[1].X {
							points = append(points, Point{X: i, Y: upper - i})
						} else {
							if points[0].Y < points[1].Y {
								if lastInsert == 0 {
									lastInsert = lwr + 1
								} else {
									lastInsert++
								}
							} else {
								if lastInsert == 0 {
									lastInsert = upper - 1
								} else {
									lastInsert--
								}
							}
							points = append(points, Point{X: i, Y: lastInsert})
						}
					} else {
						var lwr int
						if points[0].Y > points[1].Y {
							lwr = points[1].Y
						} else {
							lwr = points[0].Y
						}
						if lastInsert == 0 {
							lastInsert = lwr + 1
						} else {
							lastInsert++
						}
						points = append(points, Point{X: i, Y: lastInsert})
					}
				}
			}
		} else if horiz {
			for i := points[0].X + 1; i < points[1].X; i++ {
				points = append(points, Point{X: i, Y: points[0].Y})
			}
		} else if vert {
			for i := points[0].Y + 1; i < points[1].Y; i++ {
				points = append(points, Point{Y: i, X: points[0].X})
			}
		}

		result = append(result, Line{Points: points})
	}

	return result
}
