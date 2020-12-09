package main

import (
	"testing"

	"github.com/aaronjprice/advent_of_code/lib"
)

func TestParseLine(t *testing.T) {
	lib.AssertDeepEqual(t, parseLine("..##......."), []int{0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0})
	lib.AssertDeepEqual(t, parseLine("#...#...#.."), []int{1, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0})
	lib.AssertDeepEqual(t, parseLine(".#....#..#."), []int{0, 1, 0, 0, 0, 0, 1, 0, 0, 1, 0})
}

func TestCountTrees(t *testing.T) {
	input := getInput("test_data.txt")
	c := countTrees(3, 1, input)
	lib.AssertEqual(t, c, 7)
}
