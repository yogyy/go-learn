## Konversi Tipe Data

- Di Go-Lang kadang kita butuh melakukan konversi tipe data dari satu tipe ke tipe lain
- Misal kita ingin mengkonversi tipe data int32 ke int63, dan lain-lain

```go
package main

import "fmt"

func main() {
  ...
}
```

#### Kode Program Konversi Tipe Data (1)

```go
  var nilai32 int32 = 32768
  var nilai64 int64 = int64(nilai32)
  var nilai16 int16 = int16(nilai32)

  fmt.Println(nilai32)
  fmt.Println(nilai64)
  fmt.Println(nilai16)
```

#### Kode Program Konversi Tipe Data (2)

```go
  var name = "John Constantine"
  var e uint8 = name[0]
  var eString = string(e)

  fmt.Println(name)
  fmt.Println(e)
  fmt.Println(eString)
```
