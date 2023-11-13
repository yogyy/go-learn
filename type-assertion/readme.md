# Type Assertions
- Type Assertions merupakan kemampuan merubah tipe data menjadi tipe data yang diinginkan
- Fitur ini sering sekali digunakan ketika kita bertemu dengan data interface kosong
```go
func random() any {
  return "OK"
}

func main() {
  var res any = random()
  var resString string = res.(string)
  fmt.Println(resString)

  resInt := res.(int) //panic
  fmt.Println(resInt)
}
```

## Type Assertions Menggunakan Switch
- Saat salah menggunakan type assertions, maka bisa berakibat terjadi panic di aplikasi kita
- Jika panic dan tidak ter-recover, maka otomatis program kita akan mati
- Agar lebih aman, sebaiknya kita menggunakan switch expression untuk melakukan type assertions
```go
func main(){
  var res any = random()

  switch value := res.(type) {
  case string:
    fmt.Println("String", value)
  case int:
    fmt.Println("Int", value)
  default:
    fmt.Println("Unknown", value)
  }
}
```
