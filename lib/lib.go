package lib

import (
	"bufio"
	"os"
	"reflect"
	"strconv"
	"testing"
)

func ErrCheck(err error) {
	if err != nil {
		panic(err)
	}
}

func ReadLines(filename string) []string {
	f, err := os.Open(filename)
	ErrCheck(err)
	defer f.Close()

	data := []string{}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return data
}

func ReadInts(filename string) []int {
	lines := ReadLines(filename)
	data := []int{}
	for _, l := range lines {
		data = append(data, ToInt(l))
	}
	return data
}

func ToInt(s string) int {
	i, err := strconv.Atoi(s)
	ErrCheck(err)
	return i
}

func AssertEqual(t *testing.T, have interface{}, want interface{}) {
	if have != want {
		t.Fatalf("have: \n%v \nwant: \n%v", have, want)
	}
}

func AssertDeepEqual(t *testing.T, have interface{}, want interface{}) {
	if !reflect.DeepEqual(have, want) {
		t.Fatalf("have: \n%+v \nwant: \n%+v \n", have, want)
	}
}
