# THB-to-Text: Thai Baht to Word Conversion Library

## 🇹🇭 THB-to-Text (Production Ready)

**THB-to-Text** คือไลบรารีที่สร้างขึ้นเพื่อแก้ปัญหาการแปลงตัวเลขจำนวนเงินบาทไทยให้เป็นข้อความภาษาไทยอย่างถูกต้องและแม่นยำ โดยเลียนแบบฟังก์ชัน `BAHTTEXT()` ของ [Microsoft Excel](https://support.microsoft.com/en-us/office/bahttext-function-5ba4d0b4-abd3-4325-8d22-7a92d59aab9c) ซึ่งเป็นที่คุ้นเคยในวงการบัญชีและการเงิน

โค้ดที่มีคุณภาพสูง ทำงานได้อย่างรวดเร็ว ปลอดภัย และใช้งานง่ายที่สุดเท่าที่จะเป็นไปได้ สามารถนำไปใช้ได้จริงในระดับ**โปรดักชัน**ได้จริง


### ตารางตัวอย่างการใช้งาน / Examples Table

#### Basic Numbers (จำนวนพื้นฐาน)

| Input | Code | Output (Thai Text) |
|-------|------|-------------------|
| `0` | `baht.Text(0)` | ศูนย์บาทถ้วน |
| `1` | `baht.Text(1)` | หนึ่งบาทถ้วน |
| `2` | `baht.Text(2)` | สองบาทถ้วน |
| `5` | `baht.Text(5)` | ห้าบาทถ้วน |
| `9` | `baht.Text(9)` | เก้าบาทถ้วน |
| `10` | `baht.Text(10)` | สิบบาทถ้วน |
| `11` | `baht.Text(11)` | สิบเอ็ดบาทถ้วน |
| `20` | `baht.Text(20)` | ยี่สิบบาทถ้วน |
| `21` | `baht.Text(21)` | ยี่สิบเอ็ดบาทถ้วน |
| `31` | `baht.Text(31)` | สามสิบเอ็ดบาทถ้วน |
| `55` | `baht.Text(55)` | ห้าสิบห้าบาทถ้วน |
| `99` | `baht.Text(99)` | เก้าสิบเก้าบาทถ้วน |

#### Hundreds (ร้อย)

| Input | Code | Output (Thai Text) |
|-------|------|-------------------|
| `100` | `baht.Text(100)` | หนึ่งร้อยบาทถ้วน |
| `101` | `baht.Text(101)` | หนึ่งร้อยเอ็ดบาทถ้วน |
| `111` | `baht.Text(111)` | หนึ่งร้อยสิบเอ็ดบาทถ้วน |
| `120` | `baht.Text(120)` | หนึ่งร้อยยี่สิบบาทถ้วน |
| `221` | `baht.Text(221)` | สองร้อยยี่สิบเอ็ดบาทถ้วน |
| `505` | `baht.Text(505)` | ห้าร้อยห้าบาทถ้วน |
| `999` | `baht.Text(999)` | เก้าร้อยเก้าสิบเก้าบาทถ้วน |

#### Thousands (พัน)

| Input | Code | Output (Thai Text) |
|-------|------|-------------------|
| `1000` | `baht.Text(1000)` | หนึ่งพันบาทถ้วน |
| `1001` | `baht.Text(1001)` | หนึ่งพันเอ็ดบาทถ้วน |
| `1111` | `baht.Text(1111)` | หนึ่งพันหนึ่งร้อยสิบเอ็ดบาทถ้วน |
| `2500` | `baht.Text(2500)` | สองพันห้าร้อยบาทถ้วน |
| `9999` | `baht.Text(9999)` | เก้าพันเก้าร้อยเก้าสิบเก้าบาทถ้วน |

#### Large Numbers (จำนวนใหญ่)

| Input | Code | Output (Thai Text) |
|-------|------|-------------------|
| `10000` | `baht.Text(10000)` | หนึ่งหมื่นบาทถ้วน |
| `100000` | `baht.Text(100000)` | หนึ่งแสนบาทถ้วน |
| `123456` | `baht.Text(123456)` | หนึ่งแสนสองหมื่นสามพันสี่ร้อยห้าสิบหกบาทถ้วน |
| `1000000` | `baht.Text(1000000)` | หนึ่งล้านบาทถ้วน |
| `10000000` | `baht.Text(10000000)` | สิบล้านบาทถ้วน |
| `100000000` | `baht.Text(100000000)` | หนึ่งร้อยล้านบาทถ้วน |
| `1000000000` | `baht.Text(1000000000)` | หนึ่งพันล้านบาทถ้วน |
| `1000000000000` | `baht.Text(1000000000000)` | หนึ่งล้านล้านบาทถ้วน |

#### Decimals (ทศนิยม - สตางค์)

| Input | Code | Output (Thai Text) |
|-------|------|-------------------|
| `1.01` | `baht.Text(1.01)` | หนึ่งบาทหนึ่งสตางค์ |
| `1.25` | `baht.Text(1.25)` | หนึ่งบาทยี่สิบห้าสตางค์ |
| `5.50` | `baht.Text(5.50)` | ห้าบาทห้าสิบสตางค์ |
| `10.50` | `baht.Text(10.50)` | สิบบาทห้าสิบสตางค์ |
| `100.75` | `baht.Text(100.75)` | หนึ่งร้อยบาทเจ็ดสิบห้าสตางค์ |
| `1234.56` | `baht.Text(1234.56)` | หนึ่งพันสองร้อยสามสิบสี่บาทห้าสิบหกสตางค์ |
| `9012.34` | `baht.Text(9012.34)` | เก้าพันสิบสองบาทสามสิบสี่สตางค์ |

#### Special Cases (กรณีพิเศษ)

| Input | Code | Output (Thai Text) |
|-------|------|-------------------|
| `-100` | `baht.Text(-100)` | ลบหนึ่งร้อยบาทถ้วน |
| `10000001` | `baht.Text(10000001)` | สิบล้านเอ็ดบาทถ้วน |
| `1000000.01` | `baht.Text(1000000.01)` | หนึ่งล้านบาทหนึ่งสตางค์ |
| `1.234` | `baht.Text(1.234)` | หนึ่งบาทยี่สิบสามสตางค์ |

#### String Input Examples (ตัวอย่างการใส่ข้อมูลแบบ String)

| Input | Code | Output (Thai Text) |
|-------|------|-------------------|
| `"1,000"` | `baht.TextFromString("1,000")` | หนึ่งพันบาทถ้วน |
| `"1,234.56"` | `baht.TextFromString("1,234.56")` | หนึ่งพันสองร้อยสามสิบสี่บาทห้าสิบหกสตางค์ |
| `"1,234,567,890"` | `baht.TextFromString("1,234,567,890")` | หนึ่งพันสองร้อยสามสิบสี่ล้านห้าแสนหกหมื่นเจ็ดพันแปดร้อยเก้าสิบบาทถ้วน |
| `" 123.45 "` | `baht.TextFromString(" 123.45 ")` | หนึ่งร้อยยี่สิบสามบาทสี่สิบห้าสตางค์ |



### จุดประสงค์

การจัดการข้อมูลทางการเงินที่ละเอียดอ่อนอย่างถูกต้องและแม่นยำคือสิ่งสำคัญสูงสุด ไลบรารีนี้ถูกสร้างขึ้นเพื่อเป็นเครื่องมือที่คุณวางใจได้ ไม่ใช่แค่การแปลงตัวเลขเป็นข้อความ แต่เป็นการส่งมอบความมั่นใจว่าทุกธุรกรรมและทุกรายงานจะเนียนสวย

เร็ว, แม่นยำ, และ สามารถนำไปใช้ได้จริงในระดับ**โปรดักชัน**  ไลบรารีนี้ถูกสร้างอย่างพิถีพิถันเพื่อตอบโจทย์เหล่านั้น ทำให้คุณมีเครื่องมือที่เชื่อถือได้จริงในมือ และสามารถมุ่งเน้นไปที่การสร้างสรรค์นวัตกรรมใหม่ๆ ให้กับแอปพลิเคชันของคุณได้อย่างเต็มที่

  * **ความถูกต้องและแม่นยำสูง:** ไลบรารีนี้ได้รับการพัฒนาโดยคำนึงถึงทุกกฎเกณฑ์ทางภาษาและตัวเลขของไทย รวมถึงกรณีพิเศษ เช่น `ยี่สิบ`, `เอ็ด`, และการออกเสียงจำนวนเงินที่ถูกต้อง
  * **โค้ดที่เชื่อถือได้:** เหมาะสำหรับการใช้งานในระบบที่ต้องการความมั่นใจและเชื่อถือได้ ปลอดภัย และง่ายต่อการบำรุงรักษา
  * **ผ่านการทดสอบ รีวิวจากผู้เชี่ยวชาญ:** ไลบรารีนี้ได้รับการทดสอบและรีวิวจากผู้เชี่ยวชาญและมีประสบการณ์ในการใช้งานไลบรารี
  * **รองรับจำนวน ล้าน ล้าน ล้าน:** ไลบรารีนี้รองรับการแปลงตัวเลขจำนวนเงินที่มี ร้อยล้าน พันล้าน หมื่นล้าน แสนล้าน ล้านล้าน
  * **ใช้งานง่าย:** ด้วย API ที่เรียบง่ายและเอกสารประกอบที่ชัดเจน คุณสามารถนำไลบรารีนี้ไปใช้ในโปรเจกต์ของคุณได้ภายในไม่กี่นาที
  * **Open-Source และได้รับการสนับสนุนจากชุมชน:** โปรเจกต์นี้เปิดให้ทุกคนมีส่วนร่วมในการพัฒนาและปรับปรุงอย่างต่อเนื่อง ทำให้มั่นใจได้ว่าจะเป็นเครื่องมือที่ทันสมัยและเชื่อถือได้ในระยะยาว

### ติดตั้ง

```bash
# ตัวอย่างสำหรับภาษา Go
go get github.com/anuchito/baht
```

### ตัวอย่างการใช้งาน

```go
money := 12_345.50
fmt.Println(baht.Text(money))
// Output: หนึ่งหมื่นสองพันสามร้อยสี่สิบห้าบาทห้าสิบสตางค์

money := 63_233.33
fmt.Println(baht.Text(money))
// Output: หกหมื่นสามพันสองร้อยสามสิบสามบาทสามสิบสามสตางค์

money := 1234567890
fmt.Println(baht.Text(money))
// Output: หนึ่งพันสองร้อยสามสิบสี่ล้านห้าแสนหกหมื่นเจ็ดพันแปดร้อยเก้าสิบบาทถ้วน

money := 1234567890123456789
fmt.Println(baht.Text(money))
// Output: หนึ่งพันสองร้อยสามสิบสี่ล้านห้าแสนหกหมื่นเจ็ดพันแปดร้อยเก้าสิบบาทถ้วน


// From String
money := "1234567890"
fmt.Println(baht.TextFromString(money))
// Output: หนึ่งพันสองร้อยสามสิบสี่ล้านห้าแสนหกหมื่นเจ็ดพันแปดร้อยเก้าสิบบาทถ้วน

money := "1234567890.12"
fmt.Println(baht.TextFromString(money))
// Output: หนึ่งพันสองร้อยสามสิบสี่ล้านห้าแสนหกหมื่นเจ็ดพันแปดร้อยเก้าสิบบาทสิบสองสตางค์

money := "1234567890.123"
fmt.Println(baht.TextFromString(money))
// Output: หนึ่งพันสองร้อยสามสิบสี่ล้านห้าแสนหกหมื่นเจ็ดพันแปดร้อยเก้าสิบบาทสิบสองสตางค์

// From String Support comma
money := "1,000"
fmt.Println(baht.TextFromString(money))
// Output: หนึ่งพันบาทถ้วน

money := "1,234.56"
fmt.Println(baht.TextFromString(money))
// Output: หนึ่งพันสองร้อยสามสิบสี่บาทห้าสิบหกสตางค์

money := "1,234,567,890"
fmt.Println(baht.TextFromString(money))
// Output: หนึ่งพันสองร้อยสามสิบสี่ล้านห้าแสนหกหมื่นเจ็ดพันแปดร้อยเก้าสิบบาทถ้วน

```
-----

## 🇺🇸 THB-to-Text (Production Ready)

**THB-to-Text** is a library created to solve the problem of converting Thai Baht currency numbers into Thai text accurately and precisely. It mimics the `BAHTTEXT()` function from [Microsoft Excel](https://support.microsoft.com/en-us/office/bahttext-function-5ba4d0b4-abd3-4325-8d22-7a92d59aab9c), which is familiar in accounting and finance.

We are committed to providing high-quality code that works fast, securely, and is as easy to use as possible. It can be used in **production** environments.

### Purpose

Handling sensitive financial data accurately and precisely is of utmost importance. This library was created to be a tool you can trust. It's not just about converting numbers to text, but delivering confidence that every transaction and every report will be error-free.

Fast, accurate, and ready for **production** use - this library was meticulously built to meet those needs. It gives you a truly reliable tool in hand, allowing you to focus fully on creating new innovations for your applications.

* **High Accuracy and Precision:** This library is developed with consideration for all Thai language and numerical rules, including special cases such as `ยี่สิบ`, `เอ็ด`, and correct pronunciation of currency amounts.
* **Reliable Code:** Suitable for use in systems that require confidence and reliability, security, and ease of maintenance.
* **Tested and Reviewed by Experts:** This library has been tested and reviewed by experts with experience in library usage.
* **Supports Millions Upon Millions:** This library supports conversion of currency numbers with hundreds of millions, billions, ten billions, hundred billions, trillions.
* **Easy to Use:** With a simple API and clear documentation, you can integrate this library into your project within minutes.
* **Open-Source and Community-Supported:** This project is open for everyone to participate in continuous development and improvement, ensuring it will be a modern and reliable tool in the long term.

### Installation

```bash
# Example for Go language
go get github.com/anuchito/baht
```

### Usage Examples

```go
money := 12_345.50
fmt.Println(baht.Text(money))
// Output: หนึ่งหมื่นสองพันสามร้อยสี่สิบห้าบาทห้าสิบสตางค์

money := 63_233.33
fmt.Println(baht.Text(money))
// Output: หกหมื่นสามพันสองร้อยสามสิบสามบาทสามสิบสามสตางค์

money := 1234567890
fmt.Println(baht.Text(money))
// Output: หนึ่งพันสองร้อยสามสิบสี่ล้านห้าแสนหกหมื่นเจ็ดพันแปดร้อยเก้าสิบบาทถ้วน

money := 1234567890123456789
fmt.Println(baht.Text(money))
// Output: หนึ่งพันสองร้อยสามสิบสี่ล้านห้าแสนหกหมื่นเจ็ดพันแปดร้อยเก้าสิบบาทถ้วน


// From String
money := "1234567890"
fmt.Println(baht.TextFromString(money))
// Output: หนึ่งพันสองร้อยสามสิบสี่ล้านห้าแสนหกหมื่นเจ็ดพันแปดร้อยเก้าสิบบาทถ้วน

money := "1234567890.12"
fmt.Println(baht.TextFromString(money))
// Output: หนึ่งพันสองร้อยสามสิบสี่ล้านห้าแสนหกหมื่นเจ็ดพันแปดร้อยเก้าสิบบาทสิบสองสตางค์

money := "1234567890.123"
fmt.Println(baht.TextFromString(money))
// Output: หนึ่งพันสองร้อยสามสิบสี่ล้านห้าแสนหกหมื่นเจ็ดพันแปดร้อยเก้าสิบบาทสิบสองสตางค์

// From String Support comma
money := "1,000"
fmt.Println(baht.TextFromString(money))
// Output: หนึ่งพันบาทถ้วน

money := "1,234.56"
fmt.Println(baht.TextFromString(money))
// Output: หนึ่งพันสองร้อยสามสิบสี่บาทห้าสิบหกสตางค์

money := "1,234,567,890"
fmt.Println(baht.TextFromString(money))
// Output: หนึ่งพันสองร้อยสามสิบสี่ล้านห้าแสนหกหมื่นเจ็ดพันแปดร้อยเก้าสิบบาทถ้วน
```
