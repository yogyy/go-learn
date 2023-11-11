# Array

## Tipe Data Array
- Array adalah tipe data yang berisikan kumpulan data dengan tipe yang sama
- Saat membuat array, kita perlu menentukan jumlah data yang bisa ditampung oleh Array tersebut
- Daya tampung Array tidak bisa bertambah setelah Array dibuat

| Index 	| Data      	|
|-------	|-----------	|
| 0     	| John      	|
| 1     	| Constantine	|
| 2     	| Prime    	  |

```go
func main() {
  var names [3]string
  names[0] = "John"
  names[1] = "Constantine"
  names[2] = "Prime"
  fmt.Println(names[0], names[1])
}
```

### Membuat Array Langsung
Di Go-Lang kita juga bisa membuat Array secara langsung saat deklarasi variable

```go
var earths = [3]string{"Earth-Zero", "Earth-Prime", "Earth-One"}

fmt.Println(earths)
```
or
```go
var earths = [3]string{
  "Earth-Zero",
  "Earth-Prime",
  "Earth-One",
}

fmt.Println(earths)
```

## Function Array
| Operasi              	| Keterangan                      	|
|----------------------	|---------------------------------	|
| len(array)           	| Untuk mendapatkan panjang Array 	|
| array[index]         	| Mendapat data di posisi index   	|
| array[index] = value 	| Mengubah data di posisi index   	|

```go
var batmanVillains = [...]string{"Joker", "Ra's al Ghul", "Bane", "Riddler", "Penguin", "Two-Face", "Scarecrow"}

fmt.Println(len(batmanVillains))
fmt.Println(batmanVillains[2])
```



## Hati-Hati Saat Membuat Array
Saat membuat Array, kita harus berhati-hati, jika salah, maka yang kita buat bukanlah Array, melainkan Slice

Array vs Slice
```go
  iniArray := [...]int{1, 2, 4, 5} // array
  iniSlice := []int{1, 2, 4, 5} // slice

  fmt.Println(iniArray)
  fmt.Println(iniSlice)
```
