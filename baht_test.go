package bahttext

import "testing"

func TestBahtWords(t *testing.T) {
	tests := []struct {
		name  string
		input float64
		want  string
	}{
		// Zero and Single Digits
		{"zero", 0, "ศูนย์บาทถ้วน"},
		{"one", 1, "หนึ่งบาทถ้วน"},
		{"two", 2, "สองบาทถ้วน"},
		{"three", 3, "สามบาทถ้วน"},
		{"four", 4, "สี่บาทถ้วน"},
		{"five", 5, "ห้าบาทถ้วน"},
		{"six", 6, "หกบาทถ้วน"},
		{"seven", 7, "เจ็ดบาทถ้วน"},
		{"eight", 8, "แปดบาทถ้วน"},
		{"nine", 9, "เก้าบาทถ้วน"},

		// Tens
		{"ten", 10, "สิบบาทถ้วน"},
		{"twenty", 20, "ยี่สิบบาทถ้วน"},
		{"twenty-one", 21, "ยี่สิบเอ็ดบาทถ้วน"},
		{"twenty-two", 22, "ยี่สิบสองบาทถ้วน"},
		{"twenty-three", 23, "ยี่สิบสามบาทถ้วน"},
		{"twenty-four", 24, "ยี่สิบสี่บาทถ้วน"},
		{"twenty-five", 25, "ยี่สิบห้าบาทถ้วน"},
		{"twenty-six", 26, "ยี่สิบหกบาทถ้วน"},
		{"twenty-seven", 27, "ยี่สิบเจ็ดบาทถ้วน"},
		{"twenty-eight", 28, "ยี่สิบแปดบาทถ้วน"},
		{"twenty-nine", 29, "ยี่สิบเก้าบาทถ้วน"},
		{"thirty", 30, "สามสิบบาทถ้วน"},
		{"thirty-one", 31, "สามสิบเอ็ดบาทถ้วน"},
		{"forty", 40, "สี่สิบบาทถ้วน"},
		{"forty-one", 41, "สี่สิบเอ็ดบาทถ้วน"},
		{"fifty", 50, "ห้าสิบบาทถ้วน"},
		{"fifty-one", 51, "ห้าสิบเอ็ดบาทถ้วน"},
		{"fifty-five", 55, "ห้าสิบห้าบาทถ้วน"},
		{"sixty", 60, "หกสิบบาทถ้วน"},
		{"sixty-one", 61, "หกสิบเอ็ดบาทถ้วน"},
		{"seventy", 70, "เจ็ดสิบบาทถ้วน"},
		{"seventy-one", 71, "เจ็ดสิบเอ็ดบาทถ้วน"},
		{"eighty", 80, "แปดสิบบาทถ้วน"},
		{"eighty-one", 81, "แปดสิบเอ็ดบาทถ้วน"},
		{"ninety", 90, "เก้าสิบบาทถ้วน"},
		{"ninety-one", 91, "เก้าสิบเอ็ดบาทถ้วน"},
		{"ninety-nine", 99, "เก้าสิบเก้าบาทถ้วน"},

		// Hundreds
		{"one-hundred", 100, "หนึ่งร้อยบาทถ้วน"},
		{"one-hundred-one", 101, "หนึ่งร้อยเอ็ดบาทถ้วน"},
		{"one-hundred-two", 102, "หนึ่งร้อยสองบาทถ้วน"},
		{"one-hundred-three", 103, "หนึ่งร้อยสามบาทถ้วน"},
		{"one-hundred-four", 104, "หนึ่งร้อยสี่บาทถ้วน"},
		{"one-hundred-five", 105, "หนึ่งร้อยห้าบาทถ้วน"},
		{"one-hundred-six", 106, "หนึ่งร้อยหกบาทถ้วน"},
		{"one-hundred-seven", 107, "หนึ่งร้อยเจ็ดบาทถ้วน"},
		{"one-hundred-eight", 108, "หนึ่งร้อยแปดบาทถ้วน"},
		{"one-hundred-nine", 109, "หนึ่งร้อยเก้าบาทถ้วน"},
		{"one-hundred-ten", 110, "หนึ่งร้อยสิบบาทถ้วน"},
		{"one-hundred-eleven", 111, "หนึ่งร้อยสิบเอ็ดบาทถ้วน"},
		{"one-hundred-fifteen", 115, "หนึ่งร้อยสิบห้าบาทถ้วน"},
		{"one-hundred-twenty", 120, "หนึ่งร้อยยี่สิบบาทถ้วน"},
		{"one-hundred-twenty-three", 123, "หนึ่งร้อยยี่สิบสามบาทถ้วน"},
		{"one-hundred-twenty-five", 125, "หนึ่งร้อยยี่สิบห้าบาทถ้วน"},
		{"one-hundred-fifty", 150, "หนึ่งร้อยห้าสิบบาทถ้วน"},
		{"one-hundred-ninety-nine", 199, "หนึ่งร้อยเก้าสิบเก้าบาทถ้วน"},
		{"two-hundred", 200, "สองร้อยบาทถ้วน"},
		{"two-hundred-twenty-one", 221, "สองร้อยยี่สิบเอ็ดบาทถ้วน"},
		{"two-hundred-twenty-five", 225, "สองร้อยยี่สิบห้าบาทถ้วน"},
		{"two-hundred-fifty", 250, "สองร้อยห้าสิบบาทถ้วน"},
		{"two-hundred-eighty-nine", 289, "สองร้อยแปดสิบเก้าบาทถ้วน"},
		{"five-hundred-five", 505, "ห้าร้อยห้าบาทถ้วน"},
		{"seven-hundred-eighty-nine", 789, "เจ็ดร้อยแปดสิบเก้าบาทถ้วน"},
		{"nine-hundred-ninety-nine", 999, "เก้าร้อยเก้าสิบเก้าบาทถ้วน"},

		// Thousands
		{"one-thousand", 1000, "หนึ่งพันบาทถ้วน"},
		{"one-thousand-one", 1001, "หนึ่งพันเอ็ดบาทถ้วน"},
		{"one-thousand-ten", 1010, "หนึ่งพันสิบบาทถ้วน"},
		{"one-thousand-one-hundred", 1100, "หนึ่งพันหนึ่งร้อยบาทถ้วน"},
		{"one-thousand-one-hundred-eleven", 1111, "หนึ่งพันหนึ่งร้อยสิบเอ็ดบาทถ้วน"},
		{"two-thousand-five-hundred", 2500, "สองพันห้าร้อยบาทถ้วน"},
		{"five-thousand-five", 5005, "ห้าพันห้าบาทถ้วน"},
		{"nine-thousand-nine-hundred-ninety-nine", 9999, "เก้าพันเก้าร้อยเก้าสิบเก้าบาทถ้วน"},
		{"nine-thousand-twelve-and-thirty-four", 9012.34, "เก้าพันสิบสองบาทสามสิบสี่สตางค์"},

		// Thousands, Millions, and Billions
		{"ten-thousand", 10_000, "หนึ่งหมื่นบาทถ้วน"},
		{"one-hundred-thousand", 100_000, "หนึ่งแสนบาทถ้วน"},
		{"one-hundred-twenty-three-thousand-four-hundred-fifty-six", 123_456, "หนึ่งแสนสองหมื่นสามพันสี่ร้อยห้าสิบหกบาทถ้วน"},
		{"one-million", 1_000_000, "หนึ่งล้านบาทถ้วน"},
		{"one-million-two-hundred-thirty-four-thousand-five-hundred-sixty-seven", 1234567, "หนึ่งล้านสองแสนสามหมื่นสี่พันห้าร้อยหกสิบเจ็ดบาทถ้วน"},
		{"ten-million", 10_000_000, "สิบล้านบาทถ้วน"},
		{"one-hundred-million", 100_000_000, "หนึ่งร้อยล้านบาทถ้วน"},
		{"one-billion", 1_000_000_000, "หนึ่งพันล้านบาทถ้วน"},
		{"ten-billion", 10_000_000_000, "หนึ่งหมื่นล้านบาทถ้วน"},
		{"fifty-billion", 50_000_000_000, "ห้าหมื่นล้านบาทถ้วน"},
		{"one-hundred-billion", 100_000_000_000, "หนึ่งแสนล้านบาทถ้วน"},
		{"large-number", 1_234_567_890, "หนึ่งพันสองร้อยสามสิบสี่ล้านห้าแสนหกหมื่นเจ็ดพันแปดร้อยเก้าสิบบาทถ้วน"},

		// Some Trillions and Beyond
		{"one-trillion", 1_000_000_000_000, "หนึ่งล้านล้านบาทถ้วน"},
		{"very-large-number", 1_234_567_890_123, "หนึ่งล้านสองแสนสามหมื่นสี่พันห้าร้อยหกสิบเจ็ดล้านแปดแสนเก้าหมื่นหนึ่งร้อยยี่สิบสามบาทถ้วน"},

		// Floating-Point Numbers (with Satang)
		{"zero-baht-exact", 0.00, "ศูนย์บาทถ้วน"},
		{"zero-baht-twenty-five-satang", 0.25, "ศูนย์บาทยี่สิบห้าสตางค์"},
		{"zero-baht-fifty-satang", 0.50, "ศูนย์บาทห้าสิบสตางค์"},
		{"zero-baht-seventy-five-satang", 0.75, "ศูนย์บาทเจ็ดสิบห้าสตางค์"},
		{"five-baht-exact", 5.00, "ห้าบาทถ้วน"},
		{"five-baht-twenty-five-satang", 5.25, "ห้าบาทยี่สิบห้าสตางค์"},
		{"five-baht-fifty-satang", 5.50, "ห้าบาทห้าสิบสตางค์"},
		{"five-baht-seventy-five-satang", 5.75, "ห้าบาทเจ็ดสิบห้าสตางค์"},
		{"fifty-one-baht-ninety-nine-satang", -51.994, "ลบห้าสิบเอ็ดบาทเก้าสิบเก้าสตางค์"},
		{"fifty-one-baht-ninety-nine-satang", -51.995, "ลบห้าสิบสองบาทถ้วน"},
		{"fifty-one-baht-ninety-nine-satang", -51.99, "ลบห้าสิบเอ็ดบาทเก้าสิบเก้าสตางค์"},
		{"ten-baht-exact", 10.00, "สิบบาทถ้วน"},
		{"ten-baht-fifty-satang", 10.50, "สิบบาทห้าสิบสตางค์"},
		{"one-baht-one-satang", 1.01, "หนึ่งบาทหนึ่งสตางค์"},
		{"one-baht-seventy-five-satang", 1.75, "หนึ่งบาทเจ็ดสิบห้าสตางค์"},
		{"one-hundred-baht-fifty-satang", 100.50, "หนึ่งร้อยบาทห้าสิบสตางค์"},
		{"one-thousand-baht-five-satang", 1000.05, "หนึ่งพันบาทห้าสตางค์"},
		{"one-million-baht-one-satang", 1000000.01, "หนึ่งล้านบาทหนึ่งสตางค์"},
		{"one-thousand-two-hundred-thirty-four-baht-five-satang", 1234.05, "หนึ่งพันสองร้อยสามสิบสี่บาทห้าสตางค์"},
		{"large-float", 123456789.25, "หนึ่งร้อยยี่สิบสามล้านสี่แสนห้าหมื่นหกพันเจ็ดร้อยแปดสิบเก้าบาทยี่สิบห้าสตางค์"},

		// Edge Cases & Special Combinations
		{"negative-one-hundred", -100, "ลบหนึ่งร้อยบาทถ้วน"},
		{"ten-million-one", 10_000_001, "สิบล้านเอ็ดบาทถ้วน"},
		{"two-hundred-million-one", 200_000_001, "สองร้อยล้านเอ็ดบาทถ้วน"},
		{"one-billion-one-satang", 1_000_000_000.01, "หนึ่งพันล้านบาทหนึ่งสตางค์"},
		{"one-baht-rounded-satang", 1.234, "หนึ่งบาทยี่สิบสามสตางค์"},
		{"very-large-float", 123_456_789_012.34, "หนึ่งแสนสองหมื่นสามพันสี่ร้อยห้าสิบหกล้านเจ็ดแสนแปดหมื่นเก้าพันสิบสองบาทสามสิบสี่สตางค์"},
		{"very-large-total", 870886734867267.000000, "แปดร้อยเจ็ดสิบล้านแปดแสนแปดหมื่นหกพันเจ็ดร้อยสามสิบสี่ล้านแปดแสนหกหมื่นเจ็ดพันสองร้อยหกสิบเจ็ดบาทถ้วน"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Words(tt.input)
			if result != tt.want {
				t.Errorf("Text(%f) = %s, want %s", tt.input, result, tt.want)
			}
		})
	}
}

func TestWordsFromString(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    string
		wantErr bool
	}{
		// Valid string inputs
		{"zero-string", "0", "ศูนย์บาทถ้วน", false},
		{"one-string", "1", "หนึ่งบาทถ้วน", false},
		{"decimal-string", "1234.56", "หนึ่งพันสองร้อยสามสิบสี่บาทห้าสิบหกสตางค์", false},
		{"negative-string", "-100", "ลบหนึ่งร้อยบาทถ้วน", false},
		{"large-number-string", "1000000", "หนึ่งล้านบาทถ้วน", false},
		{"with-spaces", " 123.45 ", "หนึ่งร้อยยี่สิบสามบาทสี่สิบห้าสตางค์", false},
		{"scientific-notation", "1e3", "หนึ่งพันบาทถ้วน", false},

		// Comma-separated inputs
		{"comma-thousands", "1,000", "หนึ่งพันบาทถ้วน", false},
		{"comma-with-decimal", "1,234.56", "หนึ่งพันสองร้อยสามสิบสี่บาทห้าสิบหกสตางค์", false},
		{"comma-large-number", "1,234,567", "หนึ่งล้านสองแสนสามหมื่นสี่พันห้าร้อยหกสิบเจ็ดบาทถ้วน", false},
		{"comma-very-large", "1,234,567,890", "หนึ่งพันสองร้อยสามสิบสี่ล้านห้าแสนหกหมื่นเจ็ดพันแปดร้อยเก้าสิบบาทถ้วน", false},
		{"comma-negative", "-1,000", "ลบหนึ่งพันบาทถ้วน", false},
		{"comma-with-spaces", " 1,234.56 ", "หนึ่งพันสองร้อยสามสิบสี่บาทห้าสิบหกสตางค์", false},

		// Invalid string inputs
		{"empty-string", "", "", true},
		{"invalid-text", "abc", "", true},
		{"mixed-text", "123abc", "", true},
		{"multiple-dots", "12.34.56", "", true},
		{"special-chars", "12@34", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := WordsFromString(tt.input)

			if tt.wantErr {
				if err == nil {
					t.Errorf("TextFromString(%q) expected error, got nil", tt.input)
				}
				return
			}

			if err != nil {
				t.Errorf("TextFromString(%q) unexpected error: %v", tt.input, err)
				return
			}

			if result != tt.want {
				t.Errorf("TextFromString(%q) = %s, want %s", tt.input, result, tt.want)
			}
		})
	}
}

func TestMustWordsFromString(t *testing.T) {
	// Test valid inputs
	validTests := []struct {
		name  string
		input string
		want  string
	}{
		{"valid-integer", "1234", "หนึ่งพันสองร้อยสามสิบสี่บาทถ้วน"},
		{"valid-decimal", "1234.56", "หนึ่งพันสองร้อยสามสิบสี่บาทห้าสิบหกสตางค์"},
		{"valid-negative", "-100", "ลบหนึ่งร้อยบาทถ้วน"},
	}

	for _, tt := range validTests {
		t.Run(tt.name, func(t *testing.T) {
			result := MustWordsFromString(tt.input)
			if result != tt.want {
				t.Errorf("MustTextFromString(%q) = %s, want %s", tt.input, result, tt.want)
			}
		})
	}

	// Test panic behavior
	t.Run("panic-on-invalid", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("MustTextFromString should panic on invalid input")
			}
		}()
		MustWordsFromString("invalid")
	})
}
