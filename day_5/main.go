package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/aaronjprice/advent_of_code/lib"
)

func main() {
	ids := getIds(getPasses("input.txt"))
	fmt.Println(findHighest(ids))
	fmt.Println(findMissing(ids))
}

func getPasses(filename string) []boardingPass {
	lines := lib.ReadLines(filename)
	passes := []boardingPass{}
	for _, l := range lines {
		passes = append(passes, parseBoardingPass(l))
	}
	return passes
}

func getIds(passes []boardingPass) []int {
	ids := []int{}
	for _, p := range passes {
		ids = append(ids, p.toSeat().id())
	}
	return ids
}

func findHighest(integers []int) int {
	high := 0
	for _, integer := range integers {
		if integer > high {
			high = integer
		}
	}
	return high
}

func findMissing(integers []int) []int {
	sort.Ints(integers)
	missing := []int{}
	index := 0
	for number := integers[0]; number < integers[len(integers)-1]; number++ {
		if number == integers[index] {
			index++
		} else {
			missing = append(missing, number)
		}
	}
	return missing
}

func parseBoardingPass(s string) boardingPass {
	return boardingPass{row: string(s[:7]), column: string(s[7:10])}
}

type boardingPass struct {
	row    string
	column string
}

func (bp boardingPass) toSeat() seat {
	return seat{
		row:    parseBinary(bp.row, "F", "B"),
		column: parseBinary(bp.column, "L", "R"),
	}
}

type seat struct {
	row    int
	column int
}

func (s seat) id() int {
	return s.row*8 + s.column
}

func parseBinary(bits string, offChar, onChar string) int {
	bits = reverse(bits)
	value := 1
	count := 0
	for i, b := range bits {
		if string(b) == onChar {
			count += value
		} else if string(b) != offChar {
			panic(fmt.Sprintf("invalid char at [%v]", i))
		}
		value = value * 2
	}
	return count
}

func reverse(s string) string {
	b := strings.Builder{}
	for i := len(s) - 1; i >= 0; i-- {
		err := b.WriteByte(s[i])
		lib.ErrCheck(err)
	}
	return b.String()
}
