# go-learn
```go
package main

import "fmt"
```
## Tipe Data Number

- Ada dua jenis tipe data Number, yaitu :
  - Integer
  - Floating Point

#### Tipe Data Integer (1)

| Tipe Data | Nilai Minimum        | Nilai Maksimum      |
| --------- | -------------------- | ------------------- |
| int8      | -128                 | 127                 |
| int16     | -32768               | 32767               |
| int32     | -2147483648          | 2147483647          |
| int64     | -9223372036854775808 | 9223372036854775807 |

#### Tipe Data Integer (2)

| Tipe Data | Nilai Minimum | Nilai Maksimum       |
| --------- | ------------- | -------------------- |
| uint8     | 0             | 255                  |
| uint16    | 0             | 65535                |
| uint32    | 0             | 4294967295           |
| uint64    | 0             | 18446744073709551615 |

#### Tipe Data Floating Point

| Tipe Data  | Nilai Minimum                                          | Nilai Maksimum |
| ---------- | ------------------------------------------------------ | -------------- |
| float32    | 1.18×10−38                                             | 3.4×1038       |
| float64    | 2.23×10−308                                            | 1.80×10308     |
| complex64  | complex numbers with float32 real and imaginary parts. |                |
| complex128 | complex numbers with float64 real and imaginary parts. |                |

Alias

| Tipe Data | Alias untuk    | Nilai Maksimum |
| --------- | -------------- | -------------- |
| byte      | uint8          | 3.4×1038       |
| rune      | int32          | 1.80×10308     |
| int       | Minimal int32  |                |
| uint      | Minimal uint32 |                |


```go
func main() {
  fmt.Println("satu", 1)
  fmt.Println("dua", 2)
  fmt.Println("tiga koma lima", 3.5)
}
```

## Tipe Data Boolean

- Tipe data boolean adalah tipe data yang memiliki nilai dua nilai, yaitu benar atau salah
- Di Go-Lang, tipe data boolean direpresentasikan menggunakan kata kunci bool

| Nilai Boolean | Keterangan |
| ------------- | ---------- |
| true          | Benar      |
| false         | Salah      |


```go
func main() {
  fmt.Println("benar =", true)
  fmt.Println("salah =", false)
}
```

## Tipe Data String

- String ada tipe data kumpulan karakter
- Jumlah karakter di dalam String bisa nol sampai tidak terhingga
- Tipe data String di Go-Lang direpresentasikan dengan kata kunci string
- Nilai data String di Go-Lang selalu diawali dengan karakter “ (petik dua) dan diakhiri dengan karakter “ (petik dua)

```go
func main() {
  fmt.Println("John Constantine ")
}
```

### Function untuk String

| Function         | Keterangan                                     |
| ---------------- | ---------------------------------------------- |
| len(“string”)    | Menghitung jumlah karakter di String           |
| “string”[number] | Mengambil karakter pada posisi yang ditentukan |

```go
func main() {
  fmt.Println(len("John"))
  fmt.Println("John Constantine"[0])
  fmt.Println("John Constantine"[1])
}
```

## Variable

- Variable adalah tempat untuk menyimpan data
- Variable digunakan agar kita bisa mengakses data yang sama dimanapun kita mau
- Di Go-Lang Variable hanya bisa menyimpan tipe data yang sama, jika kita ingin menyimpan data yang berbeda-beda jenis, kita harus membuat beberapa variable
- Untuk membuat variable, kita bisa menggunakan kata kunci var, lalu diikuti dengan nama variable dan tipe datanya

```go
func main() {
  var name = "John Constantine"
  fmt.Println(name)
  name = "Bruce Wayne"
  fmt.Println(name)
}
```

### Tipe Data Variable

- Saat kita membuat variable, maka kita wajib menyebutkan tipe data variable tersebut
- Namun jika kita langsung menginisialisasikan data pada variable nya, maka kita tidak wajib menyebutkan tipe data variable nya

### Kata Kunci Var

- Di Go-Lang, kata kunci var saat membuat variable tidak lah wajib.
- Asalkan saat membuat variable kita langsung menginisialisasi datanya
- Agar tidak perlu menggunakan kata kunci var, kita perlu menggunakan kata kunci := saat menginisialisasikan data pada variable tersebut


```go
func main() {
  name := "John Constantine"
  fmt.Println(name)
  name = "Lex Luthor"
  fmt.Println(name)
}
```

#### Deklarasi Multiple Variable

- Di Go-Lang kita bisa membuat variable secara sekaligus banyak
- Code yang dibuat akan lebih bagus dan mudah dibaca

```go
func main() {
  var (
    firtName   = "Lucifer"
    middleName = "Morning"
    lastName   = "Star"
  )

  fmt.Println(firtName)
  fmt.Println(middleName)
  fmt.Println(lastName)
}
```

## Constant

- Constant adalah variable yang nilainya tidak bisa diubah lagi setelah pertama kali diberi nilai
- Cara pembuatan constant mirip dengan variable, yang membedakan hanya kata kunci yang digunakan adalah const, bukan var
- Saat pembuatan constant, kita wajib langsung menginisialisasikan datanya

```go
func main() {
  const fullName = "Lucifer MorningStar"
  fmt.Println(fullName)
}
```

### Deklarasi Multiple Constant

Sama seperti variable, di Go-Lang juga kita bisa membuat constant secara sekaligus banyak

```go
func main() {
  const (
    firstName = "Lucifer"
    lastName = "MorningStar"
  )

  // error
  firstName = "Bruce"
  lastName = "Wayne"
}
```

© [full materi](https://docs.google.com/presentation/d/1J0DbqyuLQVnGnkbL7bX3jL6iQc6RdXy8zQkfH8rbE0Q/edit#slide=id.g74a60ad919_0_1187)
