### 🎯 Professional `go` Library: Thai Baht to Text Converter (bahttext)

**Goal:** Develop a production-quality, open-source library in `go` for converting numerical values into their Thai word representation, specifically for Thai Baht (THB) currency. The solution should be ready for production use, embodying the idiomatic style and best practices of the chosen language.

-----

### 💡 Core Functionality

The function must accept a numerical value (integer or float) and return a string representing the amount in Thai. The core logic should adhere to the following rules:

1.  **Place Value Conversion:** Convert each number group (up to six digits) into Thai words (e.g., 123 -\> `หนึ่งร้อยยี่สิบสาม`).
2.  **Number Grouping:** Process numbers in groups of six digits, using the word `ล้าน` (million) as the separator for each group. For example, 1,000,000,000,000 should be read as "หนึ่งล้านล้าน" (one million million).
3.  **Special Cases:**
      * The number `1` as a leading digit of a number group (e.g., 10,000) is simply "หนึ่ง."
      * The number `1` as the last digit of a number group (e.g., 21, 101) is "เอ็ด."
      * The number `20` is pronounced "ยี่สิบ."
4.  **Currency & Fraction Handling:**
      * Append `บาท` (baht) after the whole number part.
      * If there are no decimal digits, append `ถ้วน` (thuan).
      * If there are decimal digits, append `สตางค์` (satang) after the word representation of the decimal part.

-----

### 🛠️ Non-Functional Requirements

The quality of the code is as important as its functionality. The solution must adhere to the following principles:

  * **Idiomatic Code**: The code must feel native to `go`. Use the idiomatic expressions, data structures, and standard library features that an expert in `go` would. For example, use standard `go` error handling (e.g., exceptions in Python, `error` returns in Go) and file/package structures.
  * **Simplicity and Readability**: Prioritize clarity. The code should be easy to read and understand for a developer with an intermediate skill level. Avoid overly clever or convoluted logic.
  * **Testability**: The core conversion logic must be separated from any output or formatting. The function should be pure (given the same input, it always returns the same output) and easily testable with unit tests.
  * **Documentation**: Include clear, concise code comments where necessary. The function signature should be intuitive, and a brief usage example should be provided.

-----

### 📚 Input & Output

The function must accept a numerical value (integer or float) and return a string representing the amount in Thai. The core logic should adhere to the following rules:

| Input | type | Function | Output |
| :--- | :--- | :--- | :--- |
| 10 | `int` | `bath.Text(10)` | `สิบบาทถ้วน` |
| "10" | `string` | `bath.Text("10")` | `สิบบาทถ้วน` |
| 10.0 | `float` | `bath.Text(10.0)` | `สิบบาทถ้วน` |
| 10.00 | `float` | `bath.Text(10.00)` | `สิบบาทถ้วน` |
| 10.01 | `float` | `bath.Text(10.01)` | `สิบบาทหนึ่งสตางค์` |
| 10.1 | `float` | `bath.Text(10.1)` | `สิบบาทสิบสตางค์` |
| 10.11 | `float` | `bath.Text(10.11)` | `สิบบาทสิบเอ็ดสตางค์` |
| "10.0" | `string` | `bath.Text("10.0")` | `สิบบาทถ้วน` |
| "1,010.10" | `string` | `bath.Text("1,010.10")` | `หนึ่งพันหนึ่งร้อยบาทสิบสตางค์` |
| "1,010.1" | `string` | `bath.Text("1,010.1")` | `หนึ่งพันหนึ่งร้อยบาทสิบสตางค์` |
| "1,010.11" | `string` | `bath.Text("1,010.11")` | `หนึ่งพันหนึ่งร้อยบาทสิบเอ็ดสตางค์` |

### 🧪 Provided Test Data

The function must pass all the following test cases. The output format should be a clean string without any additional punctuation.

| Category | Input | Expected Output |
| :--- | :--- | :--- |
| **Zero and Single Digits** | 0 | `ศูนย์บาทถ้วน` |
| | 1 | `หนึ่งบาทถ้วน` |
| | 2 | `สองบาทถ้วน` |
| | 3 | `สามบาทถ้วน` |
| | 4 | `สี่บาทถ้วน` |
| | 5 | `ห้าบาทถ้วน` |
| | 6 | `หกบาทถ้วน` |
| | 7 | `เจ็ดบาทถ้วน` |
| | 8 | `แปดบาทถ้วน` |
| | 9 | `เก้าบาทถ้วน` |
| **Tens** | 10 | `สิบบาทถ้วน` |
| | 20 | `ยี่สิบบาทถ้วน` |
| | 21 | `ยี่สิบเอ็ดบาทถ้วน` |
| | 22 | `ยี่สิบสองบาทถ้วน` |
| | 23 | `ยี่สิบสามบาทถ้วน` |
| | 24 | `ยี่สิบสี่บาทถ้วน` |
| | 25 | `ยี่สิบห้าบาทถ้วน` |
| | 26 | `ยี่สิบหกบาทถ้วน` |
| | 27 | `ยี่สิบเจ็ดบาทถ้วน` |
| | 28 | `ยี่สิบแปดบาทถ้วน` |
| | 29 | `ยี่สิบเก้าบาทถ้วน` |
| | 30 | `สามสิบบาทถ้วน` |
| | 31 | `สามสิบเอ็ดบาทถ้วน` |
| | 40 | `สี่สิบบาทถ้วน` |
| | 41 | `สี่สิบเอ็ดบาทถ้วน` |
| | 50 | `ห้าสิบบาทถ้วน` |
| | 51 | `ห้าสิบเอ็ดบาทถ้วน` |
| | 55 | `ห้าสิบห้าบาทถ้วน` |
| | 60 | `หกสิบบาทถ้วน` |
| | 61 | `หกสิบเอ็ดบาทถ้วน` |
| | 70 | `เจ็ดสิบบาทถ้วน` |
| | 71 | `เจ็ดสิบเอ็ดบาทถ้วน` |
| | 80 | `แปดสิบบาทถ้วน` |
| | 81 | `แปดสิบเอ็ดบาทถ้วน` |
| | 90 | `เก้าสิบบาทถ้วน` |
| | 91 | `เก้าสิบเอ็ดบาทถ้วน` |
| | 99 | `เก้าสิบเก้าบาทถ้วน` |
| **Hundreds** | 100 | `หนึ่งร้อยบาทถ้วน` |
| | 100 | `หนึ่งร้อยบาทถ้วน` |
| | 101 | `หนึ่งร้อยเอ็ดบาทถ้วน` |
| | 102 | `หนึ่งร้อยสองบาทถ้วน` |
| | 103 | `หนึ่งร้อยสามบาทถ้วน` |
| | 104 | `หนึ่งร้อยสี่บาทถ้วน` |
| | 105 | `หนึ่งร้อยห้าบาทถ้วน` |
| | 106 | `หนึ่งร้อยหกบาทถ้วน` |
| | 107 | `หนึ่งร้อยเจ็ดบาทถ้วน` |
| | 108 | `หนึ่งร้อยแปดบาทถ้วน` |
| | 109 | `หนึ่งร้อยเก้าบาทถ้วน` |
| | 110 | `หนึ่งร้อยสิบบาทถ้วน` |
| | 111 | `หนึ่งร้อยสิบเอ็ดบาทถ้วน` |
| | 115 | `หนึ่งร้อยสิบห้าบาทถ้วน` |
| | 120 | `หนึ่งร้อยยี่สิบบาทถ้วน` |
| | 123 | `หนึ่งร้อยยี่สิบสามบาทถ้วน` |
| | 125 | `หนึ่งร้อยยี่สิบห้าบาทถ้วน` |
| | 150 | `หนึ่งร้อยห้าสิบบาทถ้วน` |
| | 199 | `หนึ่งร้อยเก้าสิบเก้าบาทถ้วน` |
| | 200 | `สองร้อยบาทถ้วน` |
| | 221 | `สองร้อยยี่สิบเอ็ดบาทถ้วน` |
| | 225 | `สองร้อยยี่สิบห้าบาทถ้วน` |
| | 250 | `สองร้อยห้าสิบบาทถ้วน` |
| | 289 | `สองร้อยเอ็ดสิบแปดบาทถ้วน` |
| | 505 | `ห้าร้อยห้าบาทถ้วน` |
| | 789 | `เจ็ดร้อยเอ็ดสิบแปดบาทถ้วน` |
| | 999 | `เก้าร้อยเก้าสิบเก้าบาทถ้วน`
| **Thousands** | 1000 | `หนึ่งพันบาทถ้วน` |
| | 1001 | `หนึ่งพันหนึ่งบาทถ้วน` |
| | 1010 | `หนึ่งพันสิบบาทถ้วน` |
| | 1100 | `หนึ่งพันหนึ่งร้อยบาทถ้วน` |
| | 1111 | `หนึ่งพันหนึ่งร้อยสิบเอ็ดบาทถ้วน` |
| | 2500 | `สองพันห้าร้อยบาทถ้วน` |
| | 5005 | `ห้าพันห้าบาทถ้วน` |
| | 9999 | `เก้าพันเก้าร้อยเก้าสิบเก้าบาทถ้วน` |
| **Thousands, Millions, and Billions** | 10000 | `หนึ่งหมื่นบาทถ้วน` |
| | 100000 | `หนึ่งแสนบาทถ้วน` |
| | 123456 | `หนึ่งแสนสองหมื่นสามพันสี่ร้อยห้าสิบหกบาทถ้วน` |
| | 1000000 | `หนึ่งล้านบาทถ้วน` |
| | 1234567 | `หนึ่งล้านสองแสนสามหมื่นสี่พันห้าร้อยหกสิบเจ็ดบาทถ้วน` |
| | 10000000 | `สิบล้านบาทถ้วน` |
| | 100000000 | `หนึ่งร้อยล้านบาทถ้วน` |
| | 1000000000 | `หนึ่งพันล้านบาทถ้วน` |
| | 10000000000 | `หนึ่งหมื่นล้านบาทถ้วน` |
| | 1234567890 | `หนึ่งพันสองร้อยสามสิบสี่ล้านห้าแสนหกหมื่นเจ็ดพันแปดร้อยเก้าสิบบาทถ้วน` |
| **Trillions and Beyond** | 1000000000000 | `หนึ่งล้านล้านบาทถ้วน` |
| | 1234567890123 | `หนึ่งล้านสองแสนสามหมื่นสี่พันห้าร้อยหกสิบเจ็ดล้านแปดแสนเก้าหมื่นพันสองร้อยสามสิบบาทถ้วน` |
| **Floating-Point Numbers (with Satang)** | 5.00 | `ห้าบาทถ้วน` |
| | 5.25 | `ห้าบาทยี่สิบห้าสตางค์` |
| | 10.00 | `สิบบาทถ้วน` |
| | 10.50 | `สิบบาทห้าสิบสตางค์` |
| | 1.01 | `หนึ่งบาทหนึ่งสตางค์` |
| | 1.75 | `หนึ่งบาทเจ็ดสิบห้าสตางค์` |
| | 100.50 | `หนึ่งร้อยบาทห้าสิบสตางค์` |
| | 1000.05 | `หนึ่งพันบาทห้าสตางค์` |
| | 1000000.01 | `หนึ่งล้านบาทหนึ่งสตางค์` |
| | 1234.05 | `หนึ่งพันสองร้อยสามสิบสี่บาทห้าสตางค์` |
| | 123456789.25 | `หนึ่งร้อยยี่สิบสามล้านสี่แสนห้าหมื่นหกพันเจ็ดร้อยแปดสิบเก้าบาทยี่สิบห้าสตางค์` |
| **Edge Cases & Special Combinations** | -100 | `ติดลบหนึ่งร้อยบาทถ้วน` |
| | 10000001 | `สิบล้านเอ็ดบาทถ้วน` |
| | 200000001 | `สองร้อยล้านเอ็ดบาทถ้วน` |
| | 1000000000.01 | `หนึ่งพันล้านบาทหนึ่งสตางค์` |
| | 1.234 | `หนึ่งบาทยี่สิบสามสตางค์` |
| | 123456789012.34 | `หนึ่งแสนสองหมื่นสามพันสี่ร้อยห้าสิบหกล้านเจ็ดร้อยแปดสิบเก้าพันสิบสองบาทสามสิบสี่สตางค์` |

-----

### 🚀 Deliverables

Please provide the complete, well-commented source code for the `bahttext` library. The output should include:

  * The primary function for the conversion.
  * Any necessary helper functions or data structures.
  * A brief, runnable example demonstrating how to use the function with an input from the provided test cases.