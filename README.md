# THB-to-Text: Thai Baht to Word Conversion Library

## 🇹🇭 THB-to-Text คืออะไร?

**THB-to-Text** คือไลบรารีที่สร้างขึ้นเพื่อแก้ปัญหาการแปลงตัวเลขจำนวนเงินบาทไทยให้เป็นข้อความภาษาไทยอย่างถูกต้องและแม่นยำ โดยเลียนแบบฟังก์ชัน `BAHTTEXT()` ของ [Microsoft Excel](https://support.microsoft.com/en-us/office/bahttext-function-5ba4d0b4-abd3-4325-8d22-7a92d59aab9c) ซึ่งเป็นที่คุ้นเคยในวงการบัญชีและการเงิน

โค้ดที่มีคุณภาพสูง ทำงานได้อย่างรวดเร็ว ปลอดภัย และใช้งานง่ายที่สุดเท่าที่จะเป็นไปได้ สามารถนำไปใช้ได้จริงในระดับ**โปรดักชัน**ได้จริง

### จุดประสงค์

การจัดการข้อมูลทางการเงินที่ละเอียดอ่อนอย่างถูกต้องและแม่นยำคือสิ่งสำคัญสูงสุด ไลบรารีนี้ถูกสร้างขึ้นเพื่อเป็นเครื่องมือที่คุณวางใจได้ ไม่ใช่แค่การแปลงตัวเลขเป็นข้อความ แต่เป็นการส่งมอบความมั่นใจว่าทุกธุรกรรมและทุกรายงานของคุณจะไม่มีข้อผิดพลาด

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

## 🇺🇸 THB-to-Text?

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