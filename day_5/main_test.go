package main

import (
	"testing"

	"github.com/aaronjprice/advent_of_code/lib"
)

func TestReverse(t *testing.T) {
	lib.AssertEqual(t, reverse("abc"), "cba")
	lib.AssertEqual(t, reverse("ghj765"), "567jhg")
}

func TestParseBinary(t *testing.T) {
	lib.AssertEqual(t, parseBinary("abab", "b", "a"), 10)
	lib.AssertEqual(t, parseBinary("a   ", " ", "a"), 8)
	lib.AssertEqual(t, parseBinary("   a", " ", "a"), 1)
	lib.AssertEqual(t, parseBinary("    ", " ", "a"), 0)
}

func TestParseBoardingPass(t *testing.T) {
	lib.AssertEqual(t, parseBoardingPass("FBFBBFFRLR"), boardingPass{row: "FBFBBFF", column: "RLR"})
}

func TestPassToSeat(t *testing.T) {
	lib.AssertEqual(t, parseBoardingPass("BFFFBBFRRR").toSeat(), seat{row: 70, column: 7})
	lib.AssertEqual(t, parseBoardingPass("FFFBBBFRRR").toSeat(), seat{row: 14, column: 7})
	lib.AssertEqual(t, parseBoardingPass("BBFFBBFRLL").toSeat(), seat{row: 102, column: 4})
}

func TestSeatId(t *testing.T) {
	lib.AssertEqual(t, seat{row: 70, column: 7}.id(), 567)
	lib.AssertEqual(t, seat{row: 14, column: 7}.id(), 119)
	lib.AssertEqual(t, seat{row: 102, column: 4}.id(), 820)
}
