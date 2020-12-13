package main

import (
	"testing"

	"github.com/aaronjprice/advent_of_code/lib"
)

func TestGetDistribution(t *testing.T) {
	lib.AssertDeepEqual(t, getDistribution(lib.ReadInts("test_data1.txt")), [3]int{7, 0, 5})
	lib.AssertDeepEqual(t, getDistribution(lib.ReadInts("test_data2.txt")), [3]int{22, 0, 10})
}

func TestGetSolutions(t *testing.T) {
	lib.AssertEqual(t, getSolutions(lib.ReadInts("test_data1.txt")), 8)
	lib.AssertEqual(t, getSolutions(lib.ReadInts("test_data2.txt")), 19208)
}
