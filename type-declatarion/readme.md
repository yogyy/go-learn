## Type Declaration

- Type Declarations adalah kemampuan membuat ulang tipe data baru dari tipe data yang sudah ada
- Type Declarations biasanya digunakan untuk membuat alias terhadap tipe data yang sudah ada, dengan tujuan agar lebih mudah dimengerti


```go
package main

import "fmt"

func main() {
  type noKTP = string
  var ktpJohn = "200999"
  var contoh string = "200666"
  var contohKtp noKTP = noKTP(contoh)
  fmt.Println(ktpJohn)
  fmt.Println(contohKtp)
}

```
