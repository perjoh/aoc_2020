package main

import "testing"

func assertPolicy(actual passwordPolicy, expected passwordPolicy, t *testing.T) {
	if actual.min != expected.min {
		t.Fatalf("min, expected: %d, actual: %d", expected.min, actual.min)
	}

	if actual.max != expected.max {
		t.Fatalf("max, expected: %d, actual: %d", expected.max, actual.max)
	}

	if actual.char != expected.char {
		t.Fatalf("char, expected: %c, actual: %c", expected.char, actual.char)
	}
}

func TestParsePassword(t *testing.T) {
	tests := []struct {
		source   string
		policy   passwordPolicy
		password string
		ok       bool
		ok2      bool
	}{
		{"1-3 a: abcde", passwordPolicy{min: 1, max: 3, char: 'a'}, "abcde", true, true},
		{"1-3 b: cdefg", passwordPolicy{min: 1, max: 3, char: 'b'}, "cdefg", false, false},
		{"2-9 c: ccccccccc", passwordPolicy{min: 2, max: 9, char: 'c'}, "ccccccccc", true, false},
	}

	for _, tt := range tests {
		t.Run(tt.source, func(t *testing.T) {
			policy, password, _ := parsePassword(tt.source)
			assertPolicy(policy, tt.policy, t)
			if password != tt.password {
				t.Fatalf("password, expected: %s, actual: %s", tt.password, password)
			}

			if pwcheck := checkPassword(password, policy); pwcheck != tt.ok {
				t.Fatalf("policy check failed for password: %s, expected: %t, actual: %t", password, tt.ok, pwcheck)
			}

			if pwcheck := checkPasswordRev(password, policy); pwcheck != tt.ok2 {
				t.Fatalf("revised policy check failed for password: %s, expected: %t, actual: %t", password, tt.ok2, pwcheck)
			}
		})
	}
}
