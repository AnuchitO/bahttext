# THB-to-Text: Thai Baht to Word Conversion Library

## 🇹🇭 THB-to-Text คืออะไร?

**THB-to-Text** คือไลบรารีที่สร้างขึ้นเพื่อแก้ปัญหาการแปลงตัวเลขจำนวนเงินบาทไทยให้เป็นข้อความภาษาไทยอย่างถูกต้องและแม่นยำ โดยเลียนแบบฟังก์ชัน `BAHTTEXT()` ของ [Microsoft Excel](https://support.microsoft.com/en-us/office/bahttext-function-5ba4d0b4-abd3-4325-8d22-7a92d59aab9c) ซึ่งเป็นที่คุ้นเคยในวงการบัญชีและการเงิน

เรามุ่งมั่นที่จะให้โค้ดที่มีคุณภาพสูง ทำงานได้อย่างรวดเร็ว ปลอดภัย และใช้งานง่ายที่สุดเท่าที่จะเป็นไปได้ สามารถนำไปใช้ได้จริงในระดับ**โปรดักชัน**ได้จริง

### จุดประสงค์

การจัดการข้อมูลทางการเงินที่ละเอียดอ่อนอย่างถูกต้องและแม่นยำคือสิ่งสำคัญสูงสุด ไลบรารีนี้ถูกสร้างขึ้นเพื่อเป็นเครื่องมือที่คุณวางใจได้ ไม่ใช่แค่การแปลงตัวเลขเป็นข้อความ แต่เป็นการส่งมอบความมั่นใจว่าทุกธุรกรรมและทุกรายงานของคุณจะไม่มีข้อผิดพลาด

เร็ว, แม่นยำ, และ สามารถนำไปใช้ได้จริงในระดับ**โปรดักชัน**  ไลบรารีนี้ถูกสร้างอย่างพิถีพิถันเพื่อตอบโจทย์เหล่านั้น ทำให้คุณมีเครื่องมือที่เชื่อถือได้จริงในมือ และสามารถมุ่งเน้นไปที่การสร้างสรรค์นวัตกรรมใหม่ๆ ให้กับแอปพลิเคชันของคุณได้อย่างเต็มที่

  * **ความถูกต้องและแม่นยำสูง:** ไลบรารีนี้ได้รับการพัฒนาโดยคำนึงถึงทุกกฎเกณฑ์ทางภาษาและตัวเลขของไทย รวมถึงกรณีพิเศษ เช่น `ยี่สิบ`, `เอ็ด`, และการออกเสียงจำนวนเงินที่ถูกต้อง
  * **โค้ดที่เชื่อถือได้:** เหมาะสำหรับการใช้งานในระบบที่ต้องการความมั่นใจและเชื่อถือได้ ปลอดภัย และง่ายต่อการบำรุงรักษา
  * **ผ่านการทดสอบ รีวิวจากผู้เชี่ยวชาญ:** ไลบรารีนี้ได้รับการทดสอบและรีวิวจากผู้เชี่ยวชาญและมีประสบการณ์ในการใช้งานไลบรารี
  * **รองรับจำนวน ล้าน ล้าน ล้าน:** ไลบรารีนี้รองรับการแปลงตัวเลขจำนวนเงินที่มี ร้อยล้าน พันล้าน หมื่นล้าน แสนล้าน ล้านล้าน ล้านล้านล้าน
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

**THB-to-Text** is a library created to solve the challenge of accurately and reliably converting Thai Baht currency figures into Thai text. It emulates the familiar `BAHTTEXT()` function from Microsoft Excel, which is a common standard in accounting and finance.

This project was initiated to build a single, authoritative tool for developers in Thailand who need to handle financial data in their systems. We are committed to providing high-quality code that is fast, reliable, and as simple to use as possible.

### Why You Should Trust This Library

  * **Uncompromising Accuracy:** The library is meticulously developed to cover all linguistic and numerical rules of the Thai language, including special cases like `yee-sip` (twenty), `ed` (one), and the correct pronunciation of large numbers.
  * **Professional-Grade Code:** We follow modern software development best practices to ensure the code is efficient, secure, and maintainable. This isn't just a quick fix; it's a long-term solution you can depend on.
  * **Dead Simple to Use:** With a straightforward API and clear documentation, you can integrate this library into your project in minutes.
  * **Open-Source and Community-Backed:** As a community-driven open-source project, it benefits from continuous contributions and improvements. This collaborative approach guarantees that the library remains up-to-date and trustworthy over time.

### Installation

```bash
# Example for Go
go get github.com/your-username/thb-to-text
```

### Usage Example

```go
amount := 12345.50
fmt.Println(baht.Text(amount))
// Output: หนึ่งหมื่นสองพันสามร้อยสี่สิบห้าบาทห้าสิบสตางค์
```