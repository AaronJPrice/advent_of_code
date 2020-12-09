package main

import (
	"fmt"

	"github.com/aaronjprice/advent_of_code/lib"
)

func main() {
	a1, a2 := solve()
	fmt.Println(a1)
	fmt.Println(a2)
}

func solve() (int, int) {
	data := readData()

	a, b := findEntries2(data)
	answer1 := a * b

	a, b, c := findEntries3(data)
	answer2 := a * b * c

	return answer1, answer2
}

func readData() []int {
	lines := lib.ReadLines("input.txt")

	data := []int{}
	for _, l := range lines {
		data = append(data, lib.ToInt(l))
	}

	return data
}

func findEntries2(data []int) (int, int) {
	for i, v1 := range data[:len(data)-1] {
		for _, v2 := range data[i+1:] {
			if v1+v2 == 2020 {
				return v1, v2
			}
		}
	}
	panic("Could not find entries")
}

func findEntries3(data []int) (int, int, int) {
	for i, v1 := range data[:len(data)-2] {
		for j, v2 := range data[i+1 : len(data)-1] {
			for _, v3 := range data[j+1:] {
				if v1+v2+v3 == 2020 {
					return v1, v2, v3
				}
			}
		}
	}
	panic("Could not find entries")
}
