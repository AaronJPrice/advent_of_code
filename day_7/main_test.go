package main

import (
	"testing"

	"github.com/aaronjprice/advent_of_code/lib"
)

func TestFindContainers(t *testing.T) {
	result := len(findContainers("shiny gold", readRuleset("test_data.txt")))
	lib.AssertDeepEqual(t, result, 4)
}

func TestParseLine(t *testing.T) {
	bag, contents := parseLine("drab brown bags contain 2 dark green bags, 2 plaid fuchsia bags, 1 muted tomato bag, 5 light tan bags.")
	lib.AssertEqual(t, bag, "drab brown")
	lib.AssertDeepEqual(t, contents, []contentsRule{
		{count: 2, bag: "dark green"},
		{count: 2, bag: "plaid fuchsia"},
		{count: 1, bag: "muted tomato"},
		{count: 5, bag: "light tan"},
	})
}

func TestFindContentsNumber(t *testing.T) {
	lib.AssertEqual(t, findContentsNumber("shiny gold", readRuleset("test_data.txt")), 32)
	lib.AssertEqual(t, findContentsNumber("shiny gold", readRuleset("test_data2.txt")), 126)
}
