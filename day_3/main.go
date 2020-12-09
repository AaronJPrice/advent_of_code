package main

import (
	"fmt"

	"github.com/aaronjprice/advent_of_code/lib"
)

func main() {
	input := getInput("input.txt")
	fmt.Println(countTrees(3, 1, input))
}

func getInput(filename string) [][]int {
	lines := lib.ReadLines(filename)
	input := [][]int{}
	for _, l := range lines {
		input = append(input, parseLine(l))
	}
	return input
}

func parseLine(line string) []int {
	data := []int{}
	for _, characterCode := range line {
		character := string(characterCode)
		if character == "." {
			data = append(data, 0)
		} else if character == "#" {
			data = append(data, 1)
		} else {
			panic("Invalid data: " + character)
		}
	}
	return data
}

func countTrees(right, down int, input [][]int) int {
	x, y := 0, 0
	numRows := len(input)
	numColumns := len(input[0])
	count := 0
	for y < numRows {
		count += input[y][x]
		x += right
		y += down
		if x >= numColumns {
			x -= numColumns
		}
	}
	return count
}
