package main

import (
	"advent"
	"fmt"
)

func findBoard(inputs []int, boards []advent.Board) (advent.Board, []advent.Cell, int) {
	var foundBoard advent.Board
	var foundCell []advent.Cell
	var foundInput int

	for _, input := range inputs {
		for w, board := range boards {

			for _, cells := range board.Rows {
				for i := 0; i < len(cells); i++ {
					if input == cells[i].Value {
						cells[i].Found = true
					}
				}

				var count = 0
				// check for row find
				for i := 0; i < len(cells); i++ {
					if cells[i].Found {
						count++
					}
					if count == 5 {
						boards[w].Found = true
						foundBoard = boards[w]
						copy(foundCell, cells)
						foundInput = input
						var boardSum = 0
						for _, bb := range boards {
							if bb.Found {
								boardSum += 1
							}
						}
						if boardSum == len(boards) {
							return foundBoard, foundCell, foundInput
						}
					}
				}

				// check for column find
				for w, b := range boards {
					for i := 0; i < 5; i++ {
						var foundCells []advent.Cell
						var count = 0
						for j := 0; j < 5; j++ {
							// check for row i
							if b.Rows[j][i].Found {
								foundCells = append(foundCells, b.Rows[j][i])
								count++
							}
							if count == 5 {
								boards[w].Found = true
								var newRows [][]advent.Cell
								copy(newRows, b.Rows)
								foundBoard = advent.Board{Rows: newRows}
								copy(foundCell, foundCells)
								foundInput = input
								var boardSum = 0
								for _, bb := range boards {
									if bb.Found {
										boardSum += 1
									}
								}
								if boardSum == len(boards) {
									print("potato")
									return foundBoard, foundCell, foundInput
								}
							}
						}
					}
				}
			}
		}

	}
	return foundBoard, foundCell, foundInput
}

func parts() {
	inputs, boards := advent.ReadNumbersAndBoard("day04/day04.txt")

	foundBoard, _, lastNum := findBoard(inputs, boards)

	var sum = 0
	for _, cells := range foundBoard.Rows {
		for _, cell := range cells {
			if !cell.Found {
				sum += cell.Value
			}
		}
	}

	fmt.Println(foundBoard)
	fmt.Println("Part1 Answer is: ", sum*lastNum)
}

func main() {
	parts()
}
