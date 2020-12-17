package util

import (
	"bufio"
	"os"
	"strconv"
)

// ReadStdin return a slice of strings.
func ReadStdin() []string {
	result := []string{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}
	return result
}

// ReadFile returns a slice of strings.
func ReadFile(filename string) []string {
	result := []string{}
	f, e := os.Open(filename)
	if e == nil {
		defer f.Close()
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			result = append(result, scanner.Text())
		}
	}
	return result
}

// ReadStdinOrFile returns a slice of strings.
func ReadStdinOrFile(filename string) []string {
	input := ReadStdin()
	if len(input) == 0 {
		input = ReadFile(filename)
	}
	return input
}

//
func ReadInput() []string {
	return ReadStdinOrFile("input.txt")
}

// ToInts returns a slice of ints.
func ToInts(strings []string) []int {
	result := []int{}
	for _, s := range strings {
		value, _ := strconv.Atoi(s)
		result = append(result, value)
	}
	return result
}
