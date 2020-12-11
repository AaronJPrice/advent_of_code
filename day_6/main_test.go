package main

import (
	"testing"

	"github.com/aaronjprice/advent_of_code/lib"
)

func TestGetGroups(t *testing.T) {
	lib.AssertDeepEqual(t,
		getGroups("test_data.txt"),
		[]group{
			{"abc"},
			{"a", "b", "c"},
			{"ab", "ac"},
			{"a", "a", "a", "a"},
			{"b"},
		},
	)
}

func TestCountAny(t *testing.T) {
	lib.AssertEqual(t, countAny(group{"abc"}), 3)
	lib.AssertEqual(t, countAny(group{"a", "b", "c"}), 3)
	lib.AssertEqual(t, countAny(group{"ab", "ac"}), 3)
	lib.AssertEqual(t, countAny(group{"a", "a", "a", "a"}), 1)
	lib.AssertEqual(t, countAny(group{"b"}), 1)
}

func TestCountAll(t *testing.T) {
	lib.AssertEqual(t, countAll(group{"abc"}), 3)
	lib.AssertEqual(t, countAll(group{"a", "b", "c"}), 0)
	lib.AssertEqual(t, countAll(group{"ab", "ac"}), 1)
	lib.AssertEqual(t, countAll(group{"a", "a", "a", "a"}), 1)
	lib.AssertEqual(t, countAll(group{"b"}), 1)
}

func TestSumGroups(t *testing.T) {
	groups := []group{{"abc"}, {"a", "b", "c"}, {"ab", "ac"}, {"a", "a", "a", "a"}, {"b"}}
	lib.AssertEqual(t, sumGroups(countAny, groups), 11)
	lib.AssertEqual(t, sumGroups(countAll, groups), 6)
}
