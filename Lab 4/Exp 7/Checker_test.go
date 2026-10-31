package main

import (
	"testing"
)

func TestDivisionLogic(t *testing.T) {
	output, err := performDivision(12, 3)
	if err != nil || output != 4 {
		t.Errorf("Standard math failed: expected 4, got %v (error: %v)", output, err)
	}

	_, badErr := performDivision(8, 0)
	if badErr == nil {
		t.Error("Safety check failed: zero denominator did not trigger an error")
	}
}

func TestSquareRootLogic(t *testing.T) {
	output, err := processSquareRoot(16)
	if err != nil || output != 4 {
		t.Errorf("Standard root failed: expected 4, got %v (error: %v)", output, err)
	}

	_, badErr := processSquareRoot(-9)
	if badErr == nil {
		t.Error("Safety check failed: negative value passed without error")
	}
}

func TestUserDataLogic(t *testing.T) {
	output, err := verifyUserData("  TESTING  ")
	if err != nil || output != "testing" {
		t.Errorf("String parsing failed: expected 'testing', got '%v' (error: %v)", output, err)
	}

	_, badErr := verifyUserData("    ")
	if badErr == nil {
		t.Error("Safety check failed: pure whitespace was incorrectly accepted")
	}
}
