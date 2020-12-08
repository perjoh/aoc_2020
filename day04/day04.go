package main

import (
	"aoc2020/util"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func tidyInput(input []string) []string {
	var result []string
	for _, s := range input {
		result = append(result, strings.Split(strings.TrimSpace(s), " ")...)
	}
	return result
}

type stringMap map[string]string

func extractPassports(input []string) []stringMap {
	input = tidyInput(input)

	result := make([]stringMap, 1)
	result[0] = make(stringMap)

	for _, line := range input {
		if len(line) > 0 {
			// Assume no duplicate fields.
			keyVal := strings.Split(line, ":")
			key := keyVal[0]
			value := keyVal[1]
			result[len(result)-1][key] = value
		} else {
			result = append(result, make(stringMap))
		}
	}

	return result
}

func checkPassport(passport stringMap) bool {
	requiredFields := []string{
		"byr",
		"iyr",
		"eyr",
		"hgt",
		"hcl",
		"ecl",
		"pid",
		//"cid",
	}

	for _, field := range requiredFields {
		if _, exists := passport[field]; !exists {
			return false
		}
	}

	return true
}

type passportValidator func(passport stringMap) bool

func countValidPassports(passports []stringMap, validatePassport passportValidator) int {
	var numValidPassports int
	for _, passport := range passports {
		if validatePassport(passport) {
			numValidPassports++
		}
	}
	return numValidPassports
}

func partA(input []string) {
	passports := extractPassports(input)
	fmt.Println(countValidPassports(passports, checkPassport))
}

func verifyYear(year string, min int, max int) bool {
	ok, _ := regexp.MatchString("[0-9]{4}", year)
	if ok {
		value, _ := strconv.Atoi(year)
		return min <= value && value <= max
	}
	return false
}

func verifyHeight(heightData string) bool {
	heightValue, e := strconv.Atoi(heightData[:len(heightData)-2])
	if e == nil {
		if strings.HasSuffix(heightData, "cm") {
			return 150 <= heightValue && heightValue <= 193

		} else if strings.HasSuffix(heightData, "in") {
			return 59 <= heightValue && heightValue <= 76
		}
	}
	return false
}

func verifyHairColor(hcl string) bool {
	ok, _ := regexp.MatchString("#[0-9a-f]{6}", hcl)
	return ok
}

func verifyEyeColor(ecl string) bool {
	validColors := map[string]struct{}{
		"amb": {},
		"blu": {},
		"brn": {},
		"gry": {},
		"grn": {},
		"hzl": {},
		"oth": {},
	}

	_, contains := validColors[ecl]
	return contains
}

func verifyPassportID(pid string) bool {
	ok, _ := regexp.MatchString("[0-9]{9}", pid)
	return ok && len(pid) == 9
}

func checkPassportRev(passport stringMap) bool {
	if checkPassport(passport) {
		return verifyYear(passport["byr"], 1920, 2002) &&
			verifyYear(passport["iyr"], 2010, 2020) &&
			verifyYear(passport["eyr"], 2020, 2030) &&
			verifyHeight(passport["hgt"]) &&
			verifyHairColor(passport["hcl"]) &&
			verifyEyeColor(passport["ecl"]) &&
			verifyPassportID(passport["pid"])
	}

	return false
}

func partB(input []string) {
	passports := extractPassports(input)
	fmt.Println(countValidPassports(passports, checkPassportRev))
}

func main() {
	input := util.ReadStdin()
	partA(input)
	partB(input)
}
