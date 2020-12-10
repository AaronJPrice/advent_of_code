package main

import (
	"testing"

	"github.com/aaronjprice/advent_of_code/lib"
)

func TestGetUnformattedPassports(t *testing.T) {
	result := getUnformattedPassports("test_data.txt")
	expect := []string{
		"ecl:gry pid:860033327 eyr:2020 hcl:#fffffd byr:1937 iyr:2017 cid:147 hgt:183cm",
		"iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884 hcl:#cfa07d byr:1929",
		"hcl:#ae17e1 iyr:2013 eyr:2024 ecl:brn pid:760753108 byr:1931 hgt:179cm",
		"hcl:#cfa07d eyr:2025 pid:166559648 iyr:2011 ecl:brn hgt:59in",
	}
	lib.AssertDeepEqual(t, result, expect)
}

func TestParsePassport(t *testing.T) {
	result := parsePassport("ecl:gry pid:860033327 eyr:2020 hcl:#fffffd byr:1937 iyr:2017 cid:147 hgt:183cm")
	expect := map[string]string{
		"ecl": "gry", "pid": "860033327", "eyr": "2020", "hcl": "#fffffd",
		"byr": "1937", "iyr": "2017", "cid": "147", "hgt": "183cm",
	}
	lib.AssertDeepEqual(t, result, expect)
}

func TestGetInput(t *testing.T) {
	result := getInput("test_data.txt")
	expect := []map[string]string{
		map[string]string{
			"ecl": "gry", "pid": "860033327", "eyr": "2020", "hcl": "#fffffd",
			"byr": "1937", "iyr": "2017", "cid": "147", "hgt": "183cm",
		},
		map[string]string{
			"iyr": "2013", "ecl": "amb", "cid": "350", "eyr": "2023", "pid": "028048884",
			"hcl": "#cfa07d", "byr": "1929",
		},
		map[string]string{
			"hcl": "#ae17e1", "iyr": "2013",
			"eyr": "2024",
			"ecl": "brn", "pid": "760753108", "byr": "1931",
			"hgt": "179cm",
		},
		map[string]string{
			"hcl": "#cfa07d", "eyr": "2025", "pid": "166559648",
			"iyr": "2011", "ecl": "brn", "hgt": "59in",
		},
	}
	lib.AssertDeepEqual(t, result, expect)
}

func TestIsValid(t *testing.T) {
	lib.AssertEqual(t,
		hasRequiredFields(parsePassport("ecl:gry pid:860033327 eyr:2020 hcl:#fffffd byr:1937 iyr:2017 cid:147 hgt:183cm")),
		true,
	)

	lib.AssertEqual(t,
		hasRequiredFields(parsePassport("iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884 hcl:#cfa07d byr:1929")),
		false,
	)

	lib.AssertEqual(t,
		hasRequiredFields(parsePassport("hcl:#ae17e1 iyr:2013 eyr:2024 ecl:brn pid:760753108 byr:1931 hgt:179cm")),
		true,
	)

	lib.AssertEqual(t,
		hasRequiredFields(parsePassport("hcl:#cfa07d eyr:2025 pid:166559648 iyr:2011 ecl:brn hgt:59in")),
		false,
	)
}

func TestCountValid(t *testing.T) {
	test_data := getInput("test_data.txt")
	count := countValid(hasRequiredFields, test_data)
	lib.AssertEqual(t, count, 2)
}

func TestCheckByr(t *testing.T) {
	lib.AssertEqual(t, checkByr(map[string]string{"byr": "2002"}), true)
	lib.AssertEqual(t, checkByr(map[string]string{"byr": "2003"}), false)
}

func TestCheckHgt(t *testing.T) {
	lib.AssertEqual(t, checkHgt(map[string]string{"hgt": "60in"}), true)
	lib.AssertEqual(t, checkHgt(map[string]string{"hgt": "190cm"}), true)
	lib.AssertEqual(t, checkHgt(map[string]string{"hgt": "190in"}), false)
	lib.AssertEqual(t, checkHgt(map[string]string{"hgt": "190"}), false)
}

func TestCheckHcl(t *testing.T) {
	lib.AssertEqual(t, checkHcl(map[string]string{"hcl": "#123abc"}), true)
	lib.AssertEqual(t, checkHcl(map[string]string{"hcl": "#123abz"}), false)
	lib.AssertEqual(t, checkHcl(map[string]string{"hcl": "123abc"}), false)
}

func TestCheckEcl(t *testing.T) {
	lib.AssertEqual(t, checkEcl(map[string]string{"ecl": "brn"}), true)
	lib.AssertEqual(t, checkEcl(map[string]string{"ecl": "wat"}), false)
}

func TestCheckPid(t *testing.T) {
	lib.AssertEqual(t, checkPid(map[string]string{"pid": "000000001"}), true)
	lib.AssertEqual(t, checkPid(map[string]string{"pid": "0123456789"}), false)
}

func TestCheckMultiple(t *testing.T) {
	lib.AssertEqual(t, countValid(multiChecks(allChecks), getInput("valid_passports.txt")), 4)
	lib.AssertEqual(t, countValid(multiChecks(allChecks), getInput("invalid_passports.txt")), 0)
}
