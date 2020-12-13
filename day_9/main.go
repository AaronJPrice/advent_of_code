package main

import (
	"fmt"

	"github.com/aaronjprice/advent_of_code/lib"
)

func main() {
	answer1 := part1(25, lib.ReadInts("input.txt"))
	fmt.Println(answer1)
	fmt.Println(part2(answer1, lib.ReadInts("input.txt")))
}

func part1(size int, data []int) int {
	for i := size; i < len(data); i++ {
		if !isValid(data[i], data[i-size:i]) {
			return data[i]
		}
	}
	panic("Could not find number")
}

func isValid(number int, data []int) bool {
	for i, v1 := range data[:len(data)-1] {
		for _, v2 := range data[i+1:] {
			if v1+v2 == number {
				return true
			}
		}
	}
	return false
}

func part2(target int, data []int) int {
	return solve(findSet(target, data))
}

func findSet(target int, data []int) []int {
	for i := 0; i < len(data)-1; i++ {
		cumulative := data[i]
		for j := i + 1; j < len(data); j++ {
			cumulative += data[j]
			if cumulative == target {
				return data[i : j+1]
			} else if cumulative > target {
				break
			}
		}
	}
	panic("Could not find set")
}

func solve(data []int) int {
	min, max := data[0], data[0]
	for _, x := range data {
		if x < min {
			min = x
		}
		if max < x {
			max = x
		}
	}
	return min + max
}
