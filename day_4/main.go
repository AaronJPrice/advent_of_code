package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/aaronjprice/advent_of_code/lib"
)

func main() {
	input := getInput("input.txt")

	fmt.Println(countValid(hasRequiredFields, input))

	fmt.Println(countValid(multiChecks(allChecks), input))
}

func getInput(filename string) []map[string]string {
	unformattedPassports := getUnformattedPassports(filename)

	var input []map[string]string
	for _, unformatted := range unformattedPassports {
		input = append(input, parsePassport(unformatted))
	}

	return input
}

func getUnformattedPassports(filename string) []string {
	lines := lib.ReadLines(filename)

	passports := []string{}
	passport := ""
	for _, l := range lines {
		if l == "" {
			passports = append(passports, passport)
			passport = ""
			continue
		}
		if passport != "" {
			passport += " "
		}
		passport += l
	}
	passports = append(passports, passport)

	return passports
}

func parsePassport(s string) map[string]string {
	fields := strings.Split(s, " ")
	passport := make(map[string]string)
	for _, f := range fields {
		keyAndValue := strings.Split(f, ":")
		passport[keyAndValue[0]] = keyAndValue[1]
	}
	return passport
}

func countValid(isValidFunc func(map[string]string) bool, passports []map[string]string) int {
	c := 0
	for _, p := range passports {
		if isValidFunc(p) {
			c++
		}
	}
	return c
}

var allChecks = []func(map[string]string) bool{
	hasRequiredFields, checkByr, checkIyr, checkEyr, checkHgt, checkHcl, checkEcl, checkPid,
}

// Ignored "cid"
var requiredFields = []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

func hasRequiredFields(passport map[string]string) bool {
	for _, f := range requiredFields {
		if _, fieldIsPresent := passport[f]; !fieldIsPresent {
			return false
		}
	}
	return true
}

func multiChecks(checks []func(map[string]string) bool) func(map[string]string) bool {
	return func(passport map[string]string) bool {
		for _, c := range checks {
			if !c(passport) {
				return false
			}
		}
		return true
	}
}

// byr (Birth Year) - four digits; at least 1920 and at most 2002.
func checkByr(passport map[string]string) bool {
	return checkMinMaxNumField("byr", 1920, 2002, passport)
}

// iyr (Issue Year) - four digits; at least 2010 and at most 2020.
func checkIyr(passport map[string]string) bool {
	return checkMinMaxNumField("iyr", 2010, 2020, passport)
}

// eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
func checkEyr(passport map[string]string) bool {
	return checkMinMaxNumField("eyr", 2020, 2030, passport)
}

// hgt (Height) - a number followed by either cm or in:
// If cm, the number must be at least 150 and at most 193.
// If in, the number must be at least 59 and at most 76.
func checkHgt(passport map[string]string) bool {
	heightField, present := passport["hgt"]
	if !present {
		return false
	}

	length := len(heightField)
	if length <= 2 {
		return false
	}

	numberStr, unit := string(heightField[:length-2]), string(heightField[length-2:])

	switch unit {
	case "cm":
		return checkMinMaxNum(numberStr, 150, 193)
	case "in":
		return checkMinMaxNum(numberStr, 59, 76)
	default:
		return false
	}
}

// hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
func checkHcl(passport map[string]string) bool {
	hairColour, present := passport["hcl"]
	if !present {
		return false
	}

	if len(hairColour) != 7 {
		return false
	}

	if string(hairColour[0]) != "#" {
		return false
	}

	for i := 1; i < len(hairColour); i++ {
		if !strings.Contains("1234567890abcdef", string(hairColour[i])) {
			return false
		}
	}
	return true
}

// ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
func checkEcl(passport map[string]string) bool {
	eyeColour, present := passport["ecl"]
	if !present {
		return false
	}
	for _, validColour := range []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"} {
		if eyeColour == validColour {
			return true
		}
	}
	return false
}

// pid (Passport ID) - a nine-digit number, including leading zeroes.
func checkPid(passport map[string]string) bool {
	pid, present := passport["pid"]
	if !present {
		return false
	}

	if len(pid) != 9 {
		return false
	}

	_, err := strconv.Atoi(pid)
	if err != nil {
		return false
	}

	return true
}

func checkMinMaxNumField(field string, min, max int, passport map[string]string) bool {
	valStr, present := passport[field]
	if !present {
		return false
	}
	return checkMinMaxNum(valStr, min, max)
}

func checkMinMaxNum(valStr string, min, max int) bool {
	valInt, err := strconv.Atoi(valStr)
	if err != nil {
		return false
	}
	return min <= valInt && valInt <= max
}
