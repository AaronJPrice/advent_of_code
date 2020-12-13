package main

import (
	"fmt"
	"strings"

	"github.com/aaronjprice/advent_of_code/lib"
)

func main() {
	fmt.Println(len(findContainers("shiny gold", readRuleset("input.txt"))))
	fmt.Println(findContentsNumber("shiny gold", readRuleset("input.txt")))
}

type ruleset map[string][]contentsRule

type contentsRule struct {
	count int
	bag   string
}

func readRuleset(filename string) ruleset {
	lines := lib.ReadLines(filename)
	rs := make(ruleset)
	for _, l := range lines {
		bag, contentsRules := parseLine(l)
		rs[bag] = contentsRules
	}
	return rs
}

func parseLine(s string) (string, []contentsRule) {
	bagAndContents := strings.Split(s, " bags contain ")

	bag := bagAndContents[0]
	contents := []contentsRule{}
	if bagAndContents[1] == "no other bags." {
		return bag, contents
	}

	unparsedContents := strings.Split(strings.TrimSuffix(bagAndContents[1], "."), ", ")

	for _, unparsedContentsrule := range unparsedContents {
		contents = append(contents, parseContentsrule(unparsedContentsrule))
	}

	return bag, contents
}

func parseContentsrule(s string) contentsRule {
	s = strings.TrimRight(strings.TrimRight(s, "bags"), " ")
	count := lib.ToInt(strings.Split(s, " ")[0])
	bag := strings.TrimLeft(s, "1234567890 ")
	return contentsRule{count: count, bag: bag}
}

func findContainers(target string, rs ruleset) []string {
	return findContainersRecursive([]string{target}, []string{}, rs)
}

func findContainersRecursive(targets []string, accumulator []string, rs ruleset) []string {
	if len(targets) == 0 {
		return accumulator
	}

	containers := []string{}
	for bag, contents := range rs {
		if contains(accumulator, bag) {
			continue
		}
		for _, c := range contents {
			if contains(targets, c.bag) {
				containers = append(containers, bag)
				accumulator = append(accumulator, bag)
				break
			}
		}
	}

	return findContainersRecursive(containers, accumulator, rs)
}

func findContentsNumber(target string, rs ruleset) int {
	number := 0
	for _, content := range rs[target] {
		number += content.count * (1 + findContentsNumber(content.bag, rs))
	}
	return number
}

func contains(slice []string, target string) bool {
	for _, s := range slice {
		if s == target {
			return true
		}
	}
	return false
}
