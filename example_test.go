package baht_test

import (
	"fmt"

	"github.com/anuchito/baht"
)

// ExampleText demonstrates basic usage of the Text function
func ExampleText() {
	text := baht.Text(1234)
	fmt.Println(text)
	// Output: หนึ่งพันสองร้อยสามสิบสี่บาทถ้วน
}

// ExampleText_decimals demonstrates converting amounts with satang (decimal places)
func ExampleText_decimals() {
	amounts := []float64{1.25, 10.50, 100.75}

	for _, amount := range amounts {
		text := baht.Text(amount)
		fmt.Printf("%.2f -> %s\n", amount, text)
	}
	// Output:
	// 1.25 -> หนึ่งบาทยี่สิบห้าสตางค์
	// 10.50 -> สิบบาทห้าสิบสตางค์
	// 100.75 -> หนึ่งร้อยบาทเจ็ดสิบห้าสตางค์
}

// ExampleText_largeNumbers demonstrates converting large amounts
func ExampleText_largeNumbers() {
	amounts := []float64{1_000_000, 1_234_567_890}

	for _, amount := range amounts {
		text := baht.Text(amount)
		fmt.Printf("%.0f -> %s\n", amount, text)
	}
	// Output:
	// 1000000 -> หนึ่งล้านบาทถ้วน
	// 1234567890 -> หนึ่งพันสองร้อยสามสิบสี่ล้านห้าแสนหกหมื่นเจ็ดพันแปดร้อยเก้าสิบบาทถ้วน
}

func ExampleText_billion() {
	text := baht.Text(1_000_000_000)
	fmt.Println(text)
	// Output: หนึ่งพันล้านบาทถ้วน
}

func ExampleText_multiBillion() {
	text := baht.Text(10_000_000_000)
	fmt.Println(text)
	// Output: หนึ่งหมื่นล้านบาทถ้วน
}

func ExampleText_trillion() {
	text := baht.Text(1_000_000_000_000)
	fmt.Println(text)
	// Output: หนึ่งล้านล้านบาทถ้วน
}

func ExampleText_multiTrillion() {
	text := baht.Text(10_000_000_000_000)
	fmt.Println(text)
	// Output: สิบล้านล้านบาทถ้วน
}

// ExampleText_negative demonstrates converting negative amounts
func ExampleText_negative() {
	text := baht.Text(-100.50)
	fmt.Println(text)
	// Output: ลบหนึ่งร้อยบาทห้าสิบสตางค์
}

// ExampleText_edgeCases demonstrates edge cases and special number patterns
func ExampleText_edgeCases() {
	amounts := []float64{0, 11, 21, 111, 1001}

	for _, amount := range amounts {
		text := baht.Text(amount)
		fmt.Printf("%.0f -> %s\n", amount, text)
	}
	// Output:
	// 0 -> ศูนย์บาทถ้วน
	// 11 -> สิบเอ็ดบาทถ้วน
	// 21 -> ยี่สิบเอ็ดบาทถ้วน
	// 111 -> หนึ่งร้อยสิบเอ็ดบาทถ้วน
	// 1001 -> หนึ่งพันเอ็ดบาทถ้วน
}

// ExampleTextFromString demonstrates basic usage of the TextFromString function
func ExampleTextFromString() {
	text, err := baht.TextFromString("1234.56")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Println(text)
	// Output: หนึ่งพันสองร้อยสามสิบสี่บาทห้าสิบหกสตางค์
}

// ExampleTextFromString_various demonstrates converting various string amounts
func ExampleTextFromString_various() {
	amounts := []string{"0", "1000", "1234.56", "-100.50"}

	for _, amount := range amounts {
		text, err := baht.TextFromString(amount)
		if err != nil {
			fmt.Printf("%s -> Error: %v\n", amount, err)
			continue
		}
		fmt.Printf("%s -> %s\n", amount, text)
	}
	// Output:
	// 0 -> ศูนย์บาทถ้วน
	// 1000 -> หนึ่งพันบาทถ้วน
	// 1234.56 -> หนึ่งพันสองร้อยสามสิบสี่บาทห้าสิบหกสตางค์
	// -100.50 -> ลบหนึ่งร้อยบาทห้าสิบสตางค์
}

// ExampleTextFromString_errorHandling demonstrates error handling
func ExampleTextFromString_errorHandling() {
	invalidAmounts := []string{"abc", "12.34.56", "", "12@34"}

	for _, amount := range invalidAmounts {
		text, err := baht.TextFromString(amount)
		if err != nil {
			fmt.Printf("%q -> Error: invalid input\n", amount)
		} else {
			fmt.Printf("%q -> %s\n", amount, text)
		}
	}
	// Output:
	// "abc" -> Error: invalid input
	// "12.34.56" -> Error: invalid input
	// "" -> Error: invalid input
	// "12@34" -> Error: invalid input
}

// ExampleMustTextFromString demonstrates basic usage of the MustTextFromString function
func ExampleMustTextFromString() {
	money := fmt.Sprintf("%.2f", 1234.56)
	text := baht.MustTextFromString(money)
	fmt.Println(text)
	// Output: หนึ่งพันสองร้อยสามสิบสี่บาทห้าสิบหกสตางค์
}
