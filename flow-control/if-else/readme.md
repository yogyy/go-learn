## If Expression
- If adalah salah satu kata kunci yang digunakan untuk percabangan
- Percabangan artinya kita bisa mengeksekusi kode program tertentu ketika suatu kondisi terpenuhi
- Hampir di semua bahasa pemrograman mendukung if expression

```go
name := "John"

if name == "John" {
  fmt.Println("hello John")
}
```

## Else Expression
- Blok if akan dieksekusi ketika kondisi if bernilai true
- Kadang kita ingin melakukan eksekusi program tertentu jika kondisi if bernilai false
- Hal ini bisa dilakukan menggunakan else expression

```go
name := "John"

if name == "John" {
  fmt.Println("hello John")
} else {
  fmt.Println("you not John")
}
```

## Else If Expression
- Kadang dalam If, kita butuh membuat beberapa kondisi
- Kasus seperti ini, kita bisa menggunakan Else If expression

```go
name := "John"

if name == "John" {
  fmt.Println("hello John")
} else if name == "Barry" {
  fmt.Println("hello Barry")
} else {
  fmt.Println("you not John")
}
```

## If dengan Short Statement
- If mendukung short statement sebelum kondisi
- Hal ini sangat cocok untuk membuat statement yang sederhana sebelum melakukan pengecekan terhadap kondisi

```go
if length := len(name); length >= 5 {
  fmt.Println("nama terlalu panjang")
} else {
  fmt.Println("Nama sudah benar")
}
```
