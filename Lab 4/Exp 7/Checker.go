package main

import (
	"fmt"
	"math"
	"strings"
)

func performDivision(top, bottom float64) (float64, error) {
	if bottom == 0 {
		return 0, fmt.Errorf("math error: division by zero attempted with numerator %v", top)
	}
	return top / bottom, nil
}

func processSquareRoot(num float64) (float64, error) {
	if num < 0 {
		return 0, fmt.Errorf("invalid value: cannot calculate square root of negative number (%v)", num)
	}
	return math.Sqrt(num), nil
}

func verifyUserData(rawText string) (string, error) {
	cleaned := strings.TrimSpace(rawText)
	if cleaned == "" {
		return "", fmt.Errorf("input failure: received an empty string or just blank spaces")
	}
	return strings.ToLower(cleaned), nil
}

func main() {
	fmt.Println("Validation system ready. Use 'go test -v' to execute verification suite.")
}
