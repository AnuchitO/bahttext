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
