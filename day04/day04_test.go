package main

import (
	"fmt"
	"strconv"
	"testing"
)

func fatalError(t *testing.T, desc string, expected string, actual string) {
	t.Fatalf("%s, expected: %s, actual: %s", desc, expected, actual)
}

func fatalErrorInt(t *testing.T, desc string, expected int, actual int) {
	fatalError(t, desc, strconv.Itoa(expected), strconv.Itoa(actual))
}

func fatalErrorBool(t *testing.T, desc string, expected bool, actual bool) {
	fatalError(t, desc, fmt.Sprintf("%t", expected), fmt.Sprintf("%t", actual))
}

func TestExtractPassports(t *testing.T) {
	input := []string{
		"ecl:gry pid:860033327 eyr:2020 hcl:#fffffd",
		"byr:1937 iyr:2017 cid:147 hgt:183cm",
		"",
		"iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884",
		"hcl:#cfa07d byr:1929",
		"",
		"hcl:#ae17e1 iyr:2013",
		"eyr:2024",
		"ecl:brn pid:760753108 byr:1931",
		"hgt:179cm",
		"",
		"hcl:#cfa07d eyr:2025 pid:166559648",
		"iyr:2011 ecl:brn hgt:59in",
	}

	expectedLines := 31
	if numLines := len(tidyInput(input)); numLines != expectedLines {
		fatalErrorInt(t, "num lines", expectedLines, numLines)
	}

	passports := extractPassports(input)
	if actual := len(passports); actual != 4 {
		fatalErrorInt(t, "num passports", 4, actual)
	}

	validTests := []bool{true, false, true, false}
	for i, expected := range validTests {
		t.Run("valid passports", func(t *testing.T) {
			if actual := checkPassport(passports[i]); actual != expected {
				fatalErrorBool(t, "checkPassport", expected, actual)
			}
		})
	}
}

func TestBirthYear(t *testing.T) {
	tests := []struct {
		year  string
		valid bool
	}{
		{"2002", true},
		{"2003", false},
	}

	for _, test := range tests {
		t.Run(test.year, func(t *testing.T) {
			if ok := verifyYear(test.year, 1920, 2002); ok != test.valid {
				t.Fatal()
			}
		})
	}
}

func TestHeight(t *testing.T) {
	tests := []struct {
		height string
		valid  bool
	}{
		{"60in", true},
		{"190cm", true},
		{"190in", false},
		{"190", false},
	}

	for _, test := range tests {
		t.Run(test.height, func(t *testing.T) {
			if ok := verifyHeight(test.height); ok != test.valid {
				t.Fatal()
			}
		})
	}
}

func TestHairColor(t *testing.T) {
	tests := []struct {
		hcl   string
		valid bool
	}{
		{"#123abc", true},
		{"#123abz", false},
		{"123abc", false},
	}

	for _, test := range tests {
		t.Run(test.hcl, func(t *testing.T) {
			if ok := verifyHairColor(test.hcl); ok != test.valid {
				t.Fatal()
			}
		})
	}
}

func TestEyeColor(t *testing.T) {
	tests := []struct {
		ecl   string
		valid bool
	}{
		{"brn", true},
		{"wat", false},
	}

	for _, test := range tests {
		t.Run(test.ecl, func(t *testing.T) {
			if ok := verifyEyeColor(test.ecl); ok != test.valid {
				t.Fatal()
			}
		})
	}
}

func TestPassportID(t *testing.T) {
	tests := []struct {
		pid   string
		valid bool
	}{
		{"000000001", true},
		{"0123456789", false},
	}

	for _, test := range tests {
		t.Run(test.pid, func(t *testing.T) {
			if ok := verifyPassportID(test.pid); ok != test.valid {
				t.Fatal()
			}
		})
	}
}

func TestInvalidPassports(t *testing.T) {
	input := []string{
		"eyr:1972 cid:100",
		"hcl:#18171d ecl:amb hgt:170 pid:186cm iyr:2018 byr:1926",
		"",
		"iyr:2019",
		"hcl:#602927 eyr:1967 hgt:170cm",
		"ecl:grn pid:012533040 byr:1946",
		"",
		"hcl:dab227 iyr:2012",
		"ecl:brn hgt:182cm pid:021572410 eyr:2020 byr:1992 cid:277",
		"",
		"hgt:59cm ecl:zzz",
		"eyr:2038 hcl:74454a iyr:2023",
		"pid:3556412378 byr:2007",
	}

	passports := extractPassports(input)

	for i, passport := range passports {
		if checkPassportRev(passport) {
			t.Fatalf("passport %d was valid, expected invalid", i+1)
		}
	}

}

func TestValidPassports(t *testing.T) {
	input := []string{
		"pid:087499704 hgt:74in ecl:grn iyr:2012 eyr:2030 byr:1980",
		"hcl:#623a2f",
		"",
		"eyr:2029 ecl:blu cid:129 byr:1989",
		"iyr:2014 pid:896056539 hcl:#a97842 hgt:165cm",
		"",
		"hcl:#888785",
		"hgt:164cm byr:2001 iyr:2015 cid:88",
		"pid:545766238 ecl:hzl",
		"eyr:2022",
		"",
		"iyr:2010 hgt:158cm hcl:#b6652a ecl:blu byr:1944 eyr:2021 pid:093154719"}

	passports := extractPassports(input)

	for i, passport := range passports {
		if !checkPassportRev(passport) {
			t.Fatalf("passport %d was invalid, expected valid", i)
		}
	}
}
