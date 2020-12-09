package main

import (
	"testing"

	"github.com/aaronjprice/advent_of_code/lib"
)

func TestParseLine(t *testing.T) {
	l := parseLine("10-13 k: abc")
	lib.AssertEqual(t, l, Line{num1: 10, num2: 13, letter: "k", password: "abc"})
}

func TestIsValid1(t *testing.T) {
	lib.AssertEqual(t, isValid1(Line{num1: 1, num2: 3, letter: "a", password: "abcde"}), true)
	lib.AssertEqual(t, isValid1(Line{num1: 1, num2: 3, letter: "b", password: "cdefg"}), false)
	lib.AssertEqual(t, isValid1(Line{num1: 2, num2: 9, letter: "c", password: "ccccccccc"}), true)
}

func TestIsValid2(t *testing.T) {
	lib.AssertEqual(t, isValid2(Line{num1: 1, num2: 3, letter: "a", password: "abcde"}), true)
	lib.AssertEqual(t, isValid2(Line{num1: 1, num2: 3, letter: "b", password: "cdefg"}), false)
	lib.AssertEqual(t, isValid2(Line{num1: 2, num2: 9, letter: "c", password: "ccccccccc"}), false)
}
