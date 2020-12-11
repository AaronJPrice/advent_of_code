package main

import (
	"fmt"
	"strings"

	"github.com/aaronjprice/advent_of_code/lib"
)

func main() {
	groups := getGroups("input.txt")
	fmt.Println(sumGroups(countAny, groups))
	fmt.Println(sumGroups(countAll, groups))
}

func getGroups(filename string) []group {
	lines := lib.ReadLines(filename)

	groups := []group{}
	g := group{}
	for _, l := range lines {
		if l == "" {
			groups = append(groups, g)
			g = group{}
		} else {
			g.add(l)
		}
	}
	groups = append(groups, g)

	return groups
}

type group []string

func (g *group) add(s string) {
	*g = append(*g, s)
}

func sumGroups(countFunc func(group) int, groups []group) int {
	sum := 0
	for _, g := range groups {
		sum += countFunc(g)
	}
	return sum
}

func countAny(g group) int {
	unique := ""
	all := strings.Join(g, "")
	for _, answer := range all {
		if !strings.Contains(unique, string(answer)) {
			unique = unique + string(answer)
		}
	}
	return len(unique)
}

func countAll(g group) int {
	possible := g[0]
	for _, member := range g {
		possibleCopy := possible
		for _, p := range possibleCopy {
			if !strings.Contains(member, string(p)) {
				possible = strings.ReplaceAll(possible, string(p), "")
			}
		}
	}
	return len(possible)
}
