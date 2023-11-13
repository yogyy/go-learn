# Struct

- Struct adalah sebuah template data yang digunakan untuk menggabungkan nol atau lebih tipe data lainnya dalam satu kesatuan
- Struct biasanya representasi data dalam program aplikasi yang kita buat
- Data di struct disimpan dalam field
Sederhananya struct adalah kumpulan dari field

```go
type Customer struct {
  Name, Address string
  Age           int
}

```
## Membuat Data Struct
- Struct adalah template data atau prototype data
- Struct tidak bisa langsung digunakan
- Namun kita bisa membuat data/object dari struct yang telah kita buat
```go
func main() {
  var john Customer
  john.Address = "Earth Prime"
  john.Name = "Constantine"
  john.Age = 20

  fmt.Println(john)
}
```

## Struct Literals
Sebelumnya kita telah membuat data dari struct, namun sebenarnya ada banyak cara yang bisa kita gunakan untuk membuat data dari struct
```go
func main() {
  bruce := Customer{
    Name:    "Batman",
    Address: "Gotham",
    Age:     48,
  }

  fmt.Println(bruce)
}
```

## Struct Method
- Struct adalah tipe data seperti tipe data lainnya, dia bisa digunakan sebagai parameter untuk function
- Namun jika kita ingin menambahkan method ke dalam structs, sehingga seakan-akan sebuah struct memiliki function
- Method adalah function
```go
func (customer Customer) sayHello(name string) {
  fmt.Println("Hello", name, "my name is", customer.Name)
}

func main(){
  bruce := Customer{
    Name:    "Batman",
    Address: "Gotham",
    Age:     48,
  }

  bruce.sayHello("john")
}
```
