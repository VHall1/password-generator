package main

import "testing"

func TestGenerate(t *testing.T) {
	result := generate(6, []string{"lowercase"})

	if len(result) != 6 {
		t.Errorf("Expected password length 6, got %d", len(result))
	}

	for _, c := range result {
		if c < 'a' || c > 'z' {
			t.Errorf("Password contains non-lowercase character: %c", c)
		}
	}
}

func TestGenerateUppercase(t *testing.T) {
	result := generate(6, []string{"uppercase"})

	if len(result) != 6 {
		t.Errorf("Expected password length 6, got %d", len(result))
	}

	for _, c := range result {
		if c < 'A' || c > 'Z' {
			t.Errorf("Password contains non-uppercase character: %c", c)
		}
	}
}

func TestGenerateDigits(t *testing.T) {
	result := generate(6, []string{"digits"})

	if len(result) != 6 {
		t.Errorf("Expected password length 6, got %d", len(result))
	}

	for _, c := range result {
		if c < '0' || c > '9' {
			t.Errorf("Password contains non-digit character: %c", c)
		}
	}
}

func TestGenerateSpecial(t *testing.T) {
	specials := "@%!?*^&"
	result := generate(6, []string{"special"})

	if len(result) != 6 {
		t.Errorf("Expected password length 6, got %d", len(result))
	}

	for _, c := range result {
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
