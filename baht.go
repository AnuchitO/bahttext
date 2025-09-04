package baht

import (
	"fmt"
	"math"
	"strings"
)

var (
	unitWords  = []string{"", "หนึ่ง", "สอง", "สาม", "สี่", "ห้า", "หก", "เจ็ด", "แปด", "เก้า"}
	unitPlaces = []string{"", "สิบ", "ร้อย", "พัน", "หมื่น", "แสน", "ล้าน"}
)

// Text converts a float64 amount into its Thai word representation.
// It returns the converted string
// Don't be foolish and pass nonsensical values to this function such as NaN, Inf, or extremely large numbers.
//
// Example usage:
//
//	text := baht.Text(1234.56)
//	fmt.Println(text) // Output: หนึ่งพันสองร้อยสามสิบสี่บาทห้าสิบหกสตางค์
//
// The function supports:
//   - Whole numbers: Text(1000) -> "หนึ่งพันบาทถ้วน"
//   - Decimals (satang): Text(10.50) -> "สิบบาทห้าสิบสตางค์"
//   - Large numbers: Text(1000000) -> "หนึ่งล้านบาทถ้วน"
//   - Negative numbers: Text(-100) -> "ลบหนึ่งร้อยบาทถ้วน"
func Text(money float64) string {
	minus := ""
	if money < 0 {
		minus = "ลบ"
	}

	preciseAmount := math.Round(math.Abs(money)*100) / 100
	wholeBaht := math.Trunc(preciseAmount)
	satang := math.Round((preciseAmount - wholeBaht) * 100)

	bahtText := moneyToThaiWords(uint64(wholeBaht))

	if satang == 0 {
		return minus + bahtText + "บาทถ้วน"
	}

	satangText := moneyToThaiWords(uint64(satang))
	return fmt.Sprintf("%s%s%s%s%s", minus, bahtText, "บาท", satangText, "สตางค์")
}

// moneyToThaiWords converts an integer to its Thai word representation.
// This is a helper function to be used internally.
func moneyToThaiWords(m uint64) string {
	if m == 0 {
		return "ศูนย์"
	}

	var result strings.Builder

	// Handle millions (ล้าน)
	if m >= 1000000 {
		millionPart := m / 1000000
		m %= 1000000
		result.WriteString(moneyToThaiWords(millionPart) + unitPlaces[6])
	}

	s := fmt.Sprintf("%d", m)
	lenS := len(s)

	for i, char := range s {
		digit := int(char - '0')
		if digit == 0 {
			continue
		}

		place := lenS - i - 1
		isLast := place == 0

		// Special case for "ยี่สิบ"
		if place == 1 && digit == 2 {
			result.WriteString("ยี่สิบ")
			continue
		}

		// Special case for "เอ็ด"
		if isLast && digit == 1 && result.Len() > 0 {
			result.WriteString("เอ็ด")
			continue
		}

		// Special case for "สิบ" when digit is 1 and place is 1
		if place == 1 && digit == 1 && result.Len() == 0 {
			result.WriteString("สิบ")
			continue
		}

		// Skip "หนึ่ง" for tens place when there are higher places
		if place == 1 && digit == 1 && result.Len() > 0 {
			if place < len(unitPlaces) {
				result.WriteString(unitPlaces[place])
			}
			continue
		}

		if digit >= 0 && digit < len(unitWords) {
			result.WriteString(unitWords[digit])
		}

		// Add unit places
		if place > 0 && place < len(unitPlaces) {
			result.WriteString(unitPlaces[place])
		}
	}

	return result.String()
}
