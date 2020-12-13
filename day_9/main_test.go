package main

import (
	"testing"

	"github.com/aaronjprice/advent_of_code/lib"
)

func TestPart1(t *testing.T) {
	lib.AssertEqual(t, part1(5, lib.ReadInts("test_data.txt")), 127)
}

func TestPart2(t *testing.T) {
	lib.AssertEqual(t, part2(127, lib.ReadInts("test_data.txt")), 62)
}
