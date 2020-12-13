package main

import (
	"testing"

	"github.com/aaronjprice/advent_of_code/lib"
)

func TestParseInput(t *testing.T) {
	result := parseInput("test_data.txt")
	expect := []instruction{
		nop(+0),
		acc(+1),
		jmp(+4),
		acc(+3),
		jmp(-3),
		acc(-99),
		acc(+1),
		jmp(-4),
		acc(+6),
	}
	lib.AssertDeepEqual(t, result, expect)
}

func TestRun(t *testing.T) {
	acc, terminates := run(parseInput("test_data.txt"))
	lib.AssertEqual(t, acc, 5)
	lib.AssertEqual(t, terminates, false)
}

func TestFindAccAfterTerm(t *testing.T) {
	lib.AssertEqual(t, findAccAfterTerm(parseInput("test_data.txt")), 8)
}
