package bahttext_test

import (
	"fmt"

	"github.com/anuchito/bahttext"
)

// ExampleWords demonstrates basic usage of the Words function
func ExampleWords() {
	Words := bahttext.Words(1234)
	fmt.Println(Words)
	// Output: หนึ่งพันสองร้อยสามสิบสี่บาทถ้วน
}

// ExampleWords_decimals demonstrates converting amounts with satang (decimal places)
func ExampleWords_decimals() {
	amounts := []float64{1.25, 10.50, 100.75}

	for _, amount := range amounts {
		Words := bahttext.Words(amount)
		fmt.Printf("%.2f -> %s\n", amount, Words)
	}
	// Output:
	// 1.25 -> หนึ่งบาทยี่สิบห้าสตางค์
	// 10.50 -> สิบบาทห้าสิบสตางค์
	// 100.75 -> หนึ่งร้อยบาทเจ็ดสิบห้าสตางค์
}

// ExampleWords_largeNumbers demonstrates converting large amounts
func ExampleWords_largeNumbers() {
	amounts := []float64{1_000_000, 1_234_567_890}

	for _, amount := range amounts {
		Words := bahttext.Words(amount)
		fmt.Printf("%.0f -> %s\n", amount, Words)
	}
	// Output:
	// 1000000 -> หนึ่งล้านบาทถ้วน
	// 1234567890 -> หนึ่งพันสองร้อยสามสิบสี่ล้านห้าแสนหกหมื่นเจ็ดพันแปดร้อยเก้าสิบบาทถ้วน
}

func ExampleWords_billion() {
	Words := bahttext.Words(1_000_000_000)
	fmt.Println(Words)
	// Output: หนึ่งพันล้านบาทถ้วน
}

func ExampleWords_multiBillion() {
	Words := bahttext.Words(10_000_000_000)
	fmt.Println(Words)
	// Output: หนึ่งหมื่นล้านบาทถ้วน
}

func ExampleWords_trillion() {
	Words := bahttext.Words(1_000_000_000_000)
	fmt.Println(Words)
	// Output: หนึ่งล้านล้านบาทถ้วน
}

func ExampleWords_multiTrillion() {
	Words := bahttext.Words(10_000_000_000_000)
	fmt.Println(Words)
	// Output: สิบล้านล้านบาทถ้วน
}

// ExampleWords_negative demonstrates converting negative amounts
func ExampleWords_negative() {
	Words := bahttext.Words(-100.50)
	fmt.Println(Words)
	// Output: ลบหนึ่งร้อยบาทห้าสิบสตางค์
}

// ExampleWords_edgeCases demonstrates edge cases and special number patterns
func ExampleWords_edgeCases() {
	amounts := []float64{0, 11, 21, 111, 1001}

	for _, amount := range amounts {
		Words := bahttext.Words(amount)
		fmt.Printf("%.0f -> %s\n", amount, Words)
	}
	// Output:
	// 0 -> ศูนย์บาทถ้วน
	// 11 -> สิบเอ็ดบาทถ้วน
	// 21 -> ยี่สิบเอ็ดบาทถ้วน
	// 111 -> หนึ่งร้อยสิบเอ็ดบาทถ้วน
	// 1001 -> หนึ่งพันเอ็ดบาทถ้วน
}

// ExampleWordsFromString demonstrates basic usage of the WordsFromString function
func ExampleWordsFromString() {
	Words, err := bahttext.WordsFromString("1234.56")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Println(Words)
	// Output: หนึ่งพันสองร้อยสามสิบสี่บาทห้าสิบหกสตางค์
}

// ExampleWordsFromString_various demonstrates converting various string amounts
func ExampleWordsFromString_various() {
	amounts := []string{"0", "1000", "1234.56", "-100.50"}

	for _, amount := range amounts {
		Words, err := bahttext.WordsFromString(amount)
		if err != nil {
			fmt.Printf("%s -> Error: %v\n", amount, err)
			continue
		}
		fmt.Printf("%s -> %s\n", amount, Words)
	}
	// Output:
	// 0 -> ศูนย์บาทถ้วน
	// 1000 -> หนึ่งพันบาทถ้วน
	// 1234.56 -> หนึ่งพันสองร้อยสามสิบสี่บาทห้าสิบหกสตางค์
	// -100.50 -> ลบหนึ่งร้อยบาทห้าสิบสตางค์
}

// ExampleWordsFromString_errorHandling demonstrates error handling
func ExampleWordsFromString_errorHandling() {
	invalidAmounts := []string{"abc", "12.34.56", "", "12@34"}

	for _, amount := range invalidAmounts {
		Words, err := bahttext.WordsFromString(amount)
		if err != nil {
			fmt.Printf("%q -> Error: invalid input\n", amount)
		} else {
			fmt.Printf("%q -> %s\n", amount, Words)
		}
	}
	// Output:
	// "abc" -> Error: invalid input
	// "12.34.56" -> Error: invalid input
	// "" -> Error: invalid input
	// "12@34" -> Error: invalid input
}

// ExampleMustWordsFromString demonstrates basic usage of the MustWordsFromString function
func ExampleMustWordsFromString() {
	money := fmt.Sprintf("%.2f", 1234.56)
	Words := bahttext.MustWordsFromString(money)
	fmt.Println(Words)
	// Output: หนึ่งพันสองร้อยสามสิบสี่บาทห้าสิบหกสตางค์
}
