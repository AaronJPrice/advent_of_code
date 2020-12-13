package main

import (
	"fmt"
	"sort"

	"github.com/aaronjprice/advent_of_code/lib"
)

func main() {
	input := lib.ReadInts("input.txt")
	dist := getDistribution(input)
	fmt.Println(dist[0] * dist[2])
	fmt.Println(getSolutions(input))
}

func getDistribution(data []int) [3]int {
	sort.Ints(data)
	distributions := [3]int{0, 0, 1}

	previous := 0
	for _, jolt := range data {
		distributions[jolt-previous-1]++
		previous = jolt
	}

	return distributions
}

func getSolutions(data []int) int {
	sort.Ints(data)
	cumulative := 1
	chain := 1
	previous := 0
	for _, jolt := range data {
		if jolt-previous == 1 {
			chain++
		} else if jolt-previous == 3 {
			cumulative = cumulative * solutionsForChainLength(chain)
			chain = 1
		} else {
			panic("Unexpected difference")
		}
		previous = jolt
	}
	return cumulative * solutionsForChainLength(chain)
}

func solutionsForChainLength(length int) int {
	switch length {
	case 1:
		return 1
	case 2:
		return 1
	case 3:
		return 2
	case 4:
		return 4
	case 5:
		return 7
	case 6:
		return 13
	case 7:
		return 26
	default:
		panic("Unexpected chain length")
	}
}
