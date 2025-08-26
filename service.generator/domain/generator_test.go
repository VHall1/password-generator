package domain

import "testing"

const passwordLength = 6

func TestGenerateLength(t *testing.T) {
	pw := RandPassword(passwordLength, []string{"lowercase"})
	if len(pw) != passwordLength {
		t.Errorf("Expected password length %d, got %d", passwordLength, len(pw))
	}
}

func TestGenerate(t *testing.T) {
	pw := RandPassword(passwordLength, []string{"lowercase"})
	for _, c := range pw {
		if c < 'a' || c > 'z' {
			t.Errorf("Password contains non-lowercase character: %c", c)
		}
	}
}

func TestGenerateUppercase(t *testing.T) {
	pw := RandPassword(passwordLength, []string{"uppercase"})
	for _, c := range pw {
		if c < 'A' || c > 'Z' {
			t.Errorf("Password contains non-uppercase character: %c", c)
		}
	}
}

func TestGenerateDigits(t *testing.T) {
	pw := RandPassword(passwordLength, []string{"digits"})
	for _, c := range pw {
		if c < '0' || c > '9' {
			t.Errorf("Password contains non-digit character: %c", c)
		}
	}
}

func TestGenerateSpecial(t *testing.T) {
	specials := "@%!?*^&"
	pw := RandPassword(passwordLength, []string{"special"})
	for _, c := range pw {
		if !containsRune(specials, c) {
			t.Errorf("Password contains non-special character: %c", c)
		}
	}
}

func containsRune(s string, r rune) bool {
	for _, c := range s {
		if c == r {
			return true
		}
	}
	return false
}
