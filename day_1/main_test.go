package main

import (
	"testing"

	"github.com/aaronjprice/advent_of_code/lib"
)

func TestSolve(t *testing.T) {
	a1, a2 := solve()
	lib.AssertEqual(t, a1, 1020099)
	lib.AssertEqual(t, a2, 49214880)
}
