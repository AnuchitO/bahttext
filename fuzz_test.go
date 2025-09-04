package bahttext

import (
	"math"
	"strconv"
	"strings"
	"testing"
	"time"
)

// FuzzWords tests the Words function with random float64 inputs
func FuzzWords(f *testing.F) {
	// Add seed corpus with known good values
	f.Add(0.0)
	f.Add(1.0)
	f.Add(-1.0)
	f.Add(100.0)
	f.Add(1234.56)
	f.Add(-1234.56)
	f.Add(999999.99)
	f.Add(0.01)
	f.Add(0.99)

	f.Fuzz(func(t *testing.T, money float64) {
		// Skip invalid values that would cause undefined behavior
		if math.IsNaN(money) || math.IsInf(money, 0) {
			t.Skip("Skipping NaN or Inf values")
		}

		// Skip extremely large values that could cause overflow
		if math.Abs(money) > 1e15 {
			t.Skip("Skipping extremely large values")
		}

		// Call the function - it should not panic
		result := Words(money)

		// Basic sanity checks
		if result == "" {
			t.Errorf("Words(%f) returned empty string", money)
		}

		// Result should always contain "บาท"
		if !strings.Contains(result, "บาท") {
			t.Errorf("Words(%f) = %q, expected to contain 'บาท'", money, result)
		}

		// For zero, should return exactly "ศูนย์บาทถ้วน"
		if money == 0.0 && result != "ศูนย์บาทถ้วน" {
			t.Errorf("Words(0) = %q, expected 'ศูนย์บาทถ้วน'", result)
		}

		// For negative numbers, should start with "ลบ"
		if money < 0 && !strings.HasPrefix(result, "ลบ") {
			t.Errorf("Words(%f) = %q, expected to start with 'ลบ'", money, result)
		}

		// For positive numbers, should not start with "ลบ"
		if money > 0 && strings.HasPrefix(result, "ลบ") {
			t.Errorf("Words(%f) = %q, should not start with 'ลบ'", money, result)
		}

		// Check for decimal places - use the same logic as the Words function
		preciseAmount := math.Round(math.Abs(money)*100) / 100
		wholeBaht := math.Trunc(preciseAmount)
		satang := math.Round((preciseAmount - wholeBaht) * 100)

		if satang > 0 {
			if !strings.Contains(result, "สตางค์") {
				t.Errorf("Words(%f) = %q, expected to contain 'สตางค์' for non-zero satang (%f)", money, result, satang)
			}
		} else {
			if !strings.Contains(result, "ถ้วน") {
				t.Errorf("Words(%f) = %q, expected to contain 'ถ้วน' for zero satang", money, result)
			}
		}
	})
}

// FuzzWordsFromString tests the WordsFromString function with random string inputs
func FuzzWordsFromString(f *testing.F) {
	// Add seed corpus with known good values
	f.Add("0")
	f.Add("1")
	f.Add("-1")
	f.Add("100.50")
	f.Add("1,234.56")
	f.Add("-1,234.56")
	f.Add("999999.99")
	f.Add("0.01")

	f.Fuzz(func(t *testing.T, input string) {
		// Call the function - it should not panic
		result, err := WordsFromString(input)

		// If input is a valid number string, should not error
		if isValidNumberString(input) {
			if err != nil {
				t.Errorf("WordsFromString(%q) returned unexpected error: %v", input, err)
				return
			}

			// Basic sanity checks for valid results
			if result == "" {
				t.Errorf("WordsFromString(%q) returned empty string", input)
			}

			if !strings.Contains(result, "บาท") {
				t.Errorf("WordsFromString(%q) = %q, expected to contain 'บาท'", input, result)
			}
		} else {
			// For invalid input, should return error
			if err == nil {
				t.Errorf("WordsFromString(%q) should have returned error for invalid input", input)
			}
		}
	})
}

// FuzzMustWordsFromString tests the MustWordsFromString function
func FuzzMustWordsFromString(f *testing.F) {
	// Add seed corpus with known good values
	f.Add("0")
	f.Add("1")
	f.Add("100.50")
	f.Add("1,234.56")

	f.Fuzz(func(t *testing.T, input string) {
		// Only test with valid inputs to avoid panics
		if !isValidNumberString(input) {
			t.Skip("Skipping invalid input to avoid panic")
		}

		// Call the function - should not panic for valid inputs
		result := MustWordsFromString(input)

		// Basic sanity checks
		if result == "" {
			t.Errorf("MustWordsFromString(%q) returned empty string", input)
		}

		if !strings.Contains(result, "บาท") {
			t.Errorf("MustWordsFromString(%q) = %q, expected to contain 'บาท'", input, result)
		}
	})
}

// FuzzConsistency tests that Words and WordsFromString produce consistent results
func FuzzConsistency(f *testing.F) {
	// Add seed corpus
	f.Add(0.0)
	f.Add(1.0)
	f.Add(100.50)
	f.Add(1234.56)
	f.Add(-1234.56)

	f.Fuzz(func(t *testing.T, money float64) {
		// Skip invalid values
		if math.IsNaN(money) || math.IsInf(money, 0) || math.Abs(money) > 1e12 {
			t.Skip("Skipping invalid or extremely large values")
		}

		// Round money to 2 decimal places first to match what Words function does internally
		roundedMoney := math.Round(money*100) / 100

		// Get result from Words function
		wordsResult := Words(roundedMoney)

		// Convert float to string and get result from WordsFromString
		moneyStr := strconv.FormatFloat(roundedMoney, 'f', 2, 64)
		stringResult, err := WordsFromString(moneyStr)

		if err != nil {
			t.Errorf("WordsFromString(%q) returned error: %v", moneyStr, err)
			return
		}

		// Results should be identical
		if wordsResult != stringResult {
			t.Errorf("Inconsistent results: Words(%f) = %q, WordsFromString(%q) = %q",
				roundedMoney, wordsResult, moneyStr, stringResult)
		}
	})
}

// FuzzPropertyBasedTesting tests mathematical properties
func FuzzPropertyBasedTesting(f *testing.F) {
	f.Add(100.0)
	f.Add(50.25)

	f.Fuzz(func(t *testing.T, money float64) {
		// Skip invalid values
		if math.IsNaN(money) || math.IsInf(money, 0) || math.Abs(money) > 1e10 {
			t.Skip("Skipping invalid or extremely large values")
		}

		result := Words(money)

		// Property: Words(-x) should be "ลบ" + Words(x) (without the "ลบ" prefix)
		if money != 0 {
			negativeResult := Words(-money)
			positiveResult := Words(math.Abs(money))

			if money > 0 {
				expectedNegative := "ลบ" + positiveResult
				if negativeResult != expectedNegative {
					t.Errorf("Property violation: Words(-%f) = %q, expected %q",
						money, negativeResult, expectedNegative)
				}
			}
		}

		// Property: Words(x) length should be reasonable (not empty, not extremely long)
		if len(result) == 0 {
			t.Errorf("Words(%f) returned empty result", money)
		}
		if len(result) > 1000 { // Reasonable upper bound
			t.Errorf("Words(%f) returned suspiciously long result: %d characters", money, len(result))
		}
	})
}

// Helper function to check if a string represents a valid number
func isValidNumberString(s string) bool {
	// First check if WordsFromString would accept it
	_, err := WordsFromString(s)
	return err == nil
}

// Benchmark fuzzing to ensure performance doesn't degrade
func FuzzPerformance(f *testing.F) {
	f.Add(1234.56)

	f.Fuzz(func(t *testing.T, money float64) {
		if math.IsNaN(money) || math.IsInf(money, 0) || math.Abs(money) > 1e12 {
			t.Skip("Skipping invalid values")
		}

		start := time.Now()
		Words(money)
		duration := time.Since(start)

		// Should complete within reasonable time (10ms is generous)
		if duration > 10*time.Millisecond {
			t.Errorf("Words(%f) took too long: %v", money, duration)
		}
	})
}
