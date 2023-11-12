## Switch Expression
- Selain if expression, untuk melakukan percabangan, kita juga bisa menggunakan Switch Expression
- Switch expression sangat sederhana dibandingkan if
- Biasanya switch  expression digunakan untuk melakukan pengecekan ke kondisi dalam satu variable

```go
func main() {
  name := "Flash"

  switch name {
  case "Batman":
    fmt.Println("hello Batman")
  case "Superman":
    fmt.Println("hello Superman")
  case "Flash":
    fmt.Println("hello Flash")
  default:
    fmt.Println("who are you?")
  }
}
```

### Switch dengan Short Statement
```go
switch length := len(name); length > 8 {
case true:
  fmt.Println("nama terlalu panjang")
case false:
  fmt.Println("nama sudah benar")
}
```

### Switch Tanpa Kondisi
- Kondisi di switch expression tidak wajib
- Jika kita tidak menggunakan kondisi di switch expression, kita bisa menambahkan kondisi tersebut di setiap case nya

```go
name := "The Batman Who Laughs"

length := len(name)
switch {
case length > 10:
  fmt.Println("nama terlalu panjang")
case length > 5:
  fmt.Println("nama lumayan panjang")
default:
  fmt.Println("nama sudah benar")
}
```
