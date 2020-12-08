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

// ToInts returns a slice of ints.
func ToInts(strings []string) []int {
	result := []int{}
	for _, s := range strings {
		value, _ := strconv.Atoi(s)
		result = append(result, value)
	}
	return result
}
