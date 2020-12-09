package main

import (
	"fmt"
	"strings"

	"github.com/aaronjprice/advent_of_code/lib"
)

func main() {
	input := parseLines(lib.ReadLines("input.txt"))
	fmt.Println(countValidLines(isValid1, input))
	fmt.Println(countValidLines(isValid2, input))
}

func parseLines(unformattedLines []string) []Line {
	lines := []Line{}
	for _, l := range unformattedLines {
		lines = append(lines, parseLine(l))
	}
	return lines
}

func parseLine(unformattedLine string) Line {
	elements := strings.Split(unformattedLine, " ")
	numbers := strings.Split(elements[0], "-")
	return Line{
		num1:     lib.ToInt(numbers[0]),
		num2:     lib.ToInt(numbers[1]),
		letter:   strings.TrimSuffix(elements[1], ":"),
		password: elements[2],
	}
}

func countValidLines(isValidFunc func(Line) bool, input []Line) int {
	validLines := 0
	for _, l := range input {
		if isValidFunc(l) {
			validLines++
		}
	}
	return validLines
}

func isValid1(l Line) bool {
	c := strings.Count(l.password, l.letter)
	return l.num1 <= c && c <= l.num2
}

func isValid2(l Line) bool {
	letter1 := string(l.password[l.num1-1])
	letter2 := string(l.password[l.num2-1])
	return (letter1 == l.letter) != (letter2 == l.letter)
}

type Line struct {
	num1     int
	num2     int
	letter   string
	password string
}
