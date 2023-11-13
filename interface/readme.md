# Interface
- Interface adalah tipe data Abstract, dia tidak memiliki implementasi langsung
- Sebuah interface berisikan definisi-definisi method
- Biasanya interface digunakan sebagai kontrak
```go
type HasName interface {
  GetName() string
}

func SayHello(value HasName) {
  fmt.Println("Hello", value.GetName())
}
```

## Implementasi Interface
- Setiap tipe data yang sesuai dengan kontrak interface, secara otomatis dianggap sebagai interface tersebut
- Sehingga kita tidak perlu mengimplementasikan interface secara manual
- Hal ini agak berbeda dengan bahasa pemrograman lain yang ketika membuat interface, kita harus menyebutkan secara eksplisit akan menggunakan interface mana
```go
type Person struct {
  Name string
}

func (person Person) GetName() string {
  return person.Name
}

func main() {
  person := Person{Name: "Constantine"}
  SayHello(person)
}
```

## Interface Kosong
- Go-Lang bukanlah bahasa pemrograman yang berorientasi objek
- Biasanya dalam pemrograman berorientasi objek, ada satu data parent di puncak yang bisa dianggap sebagai semua implementasi data yang ada di bahasa pemrograman tersebut
- Contoh di Java ada java.lang.Object
- Untuk menangani kasus seperti ini, di Go-Lang kita bisa menggunakan interface kosong
- Interface kosong adalah interface yang tidak memiliki deklarasi method satupun, hal ini membuat secara otomatis semua tipe data akan menjadi implementasi nya
- Interface kosong, juga memiliki type alias bernama any


### Penggunaan Interface Kosong di Go-Lang
Ada banyak contoh penggunaan interface kosong di Go-Lang, seperti :
  - fmt.Println(a ...interface{})
  - panic(v interface{})
  - recover() interface{}
  - dan lain-lain
