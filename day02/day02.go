package main

import (
	"aoc2020/util"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type passwordPolicy struct {
	min  int
	max  int
	char byte
}

func parsePassword(input string) (passwordPolicy, string, error) {
	r, e := regexp.Compile("([0-9]+)-([0-9]+) ([a-z]): ([a-z]+)")
	if e == nil {
		m := r.FindStringSubmatch(input)
		if len(m) == 5 {
			min, _ := strconv.Atoi(m[1])
			max, _ := strconv.Atoi(m[2])
			char := m[3][0]
			return passwordPolicy{min: min, max: max, char: char}, m[4], nil
		}
	}

	return passwordPolicy{}, "", errors.New("invalid password format")
}

func checkPassword(password string, policy passwordPolicy) bool {
	count := strings.Count(password, string(policy.char))
	return count <= policy.max && policy.min <= count
}

type checkPw func(pw string, p passwordPolicy) bool

func part(input []string, checkPw checkPw) {
	var okPasswordCount int
	for _, line := range input {
		policy, password, e := parsePassword(line)
		if e == nil {
			if checkPw(password, policy) {
				okPasswordCount++
			}
		}
	}
	fmt.Println(okPasswordCount)
}

func checkPasswordRev(password string, policy passwordPolicy) bool {
	pwlen := len(password)
	if policy.min < pwlen {
		if policy.max <= pwlen {

			cmin := password[policy.min-1]
			cmax := password[policy.max-1]

			if cmin == policy.char {
				return cmax != policy.char
			}

			return cmax == policy.char
		}
	}

	return false
}

func main() {
	input := util.ReadStdin()
	part(input, checkPassword)
	part(input, checkPasswordRev)
}
