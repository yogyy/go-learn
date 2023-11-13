## Pass by Value
- Secara default di Go-Lang semua variable itu di passing by value, bukan by reference
- Artinya, jika kita mengirim sebuah variable ke dalam function, method atau variable lain, sebenarnya yang dikirim adalah duplikasi value nya

```go
type Adress struct {
  Place, City, Earth string
}

func main() {
  address1 := Adress{"The Batcave", "outside Gotham City", "Earth Prime"}
  address2 := address1

  address2.Place = "S.T.A.R. Labs"
  address2.City = "Central City"

  fmt.Println(address1) // address1 tidak berubah
  fmt.Println(address2)
}
```

#### Penjelasan Detail Pass by Value
![diagram](https://www.plantuml.com/plantuml/dpng/ZOz1QiCm44NtSugFzp7Oou9Bt52wQI6qkS34ciGgaXMaSTjGUlUQ93OnP97PI7fVdoSjatBKtZjYvyAtjp-ssnFjWZCIaRs93_6OuPSmJDX-IgEeZOcIKd6cfraxRWGLDgtWYRNXBvcobTNZ48LUapM2bw0j-nEThZnpr1RLPLr7wy_rFxhXBMzuF_9DNkVlELvvcsRc2YlfDBAxrpWuTzO9ChvyCmCdZ8TvLbvd9T6umw8_B1wd1u-ycFE2web7ks_rge87DE5WnT1IEjDxzmy0)

# Pointer
- Pointer adalah kemampuan membuat reference ke lokasi data di memory yang sama, tanpa menduplikasi data yang sudah ada
- Sederhananya, dengan kemampuan pointer, kita bisa membuat pass by reference


#### Pass by Reference dengan Pointer
![diagram](https://www.plantuml.com/plantuml/dpng/NOwn3i8W48PtdkB2tg7TQwh6E9bqy0NESbeC50E7sfY-kwXje_uBy7rSxgy2e-TeZK1ZtCtgk-vEGXoTKUH1xiWwtqju24XnXpCr34kbF8MGZ_ILsc13KCA-9LWXDtYbsQqrzqcvo44hWgtZ5ksRpbysgKCBvNZMNjlF-226r4KR0edZKWaOO7hykkBR5K9CDyZY9LlCjIUU__EAEXMjUG80)

## Operator &
Untuk membuat sebuah variable dengan nilai pointer ke variable yang lain, kita bisa menggunakan operator & diikuti dengan nama variable nya
```go
func main(){
  address1 := Adress{"The Batcave", "outside Gotham City", "Earth Prime"}
  address2 := &address1

  address2.Earth = "Earth One"
  fmt.Println(address1)
  fmt.Println(address2)
}
```
## Operator *
- Saat kita mengubah variable pointer, maka yang berubah hanya variable tersebut.
- Semua variable yang mengacu ke data yang sama tidak akan berubah
- Jika kita ingin mengubah seluruh variable yang mengacu ke data tersebut, kita bisa menggunakan operator *

#### Variable Pointer Mengacu Ke Data yang Sama
![diagram](https://www.plantuml.com/plantuml/dpng/NOwn3i8W48PtdkB2tg7TQwh6E9bqy0NESbeC50E7sfY-kwXje_uBy7rSxgy2e-TeZK1ZtCtgk-vEGXoTKUH1xiWwtqju24XnXpCr34kbF8MGZ_ILsc13KCA-9LWXDtYbsQqrzqcvo44hWgtZ5ksRpbysgKCBvNZMNjlF-226r4KR0edZKWaOO7hykkBR5K9CDyZY9LlCjIUU__EAEXMjUG80)

#### Kode Program Operator * (1)
```go
func main(){
	address1 := Address{"The Batcave", "outside Gotham City", "Earth Prime"}
	address2 := &address1
	address2.Place = "Wayne Manor"

	address2 = &Address{"S.T.A.R. Labs", "Central City", "Earth Prime"}
	fmt.Println(address1) // address 1 tidak berubah
	fmt.Println(address2)
}
```
#### Tanpa Operator *
![diagram](https://www.plantuml.com/plantuml/dpng/ZO_1IiGm48RlynH3xXjexvLT5_6YiEY-m6mpQoEJHfA9keZlRjMsMEZ1_vBa_nA-p3Oh9Ir1W4JYM_3lldy2E4TndD3SSuZfOl03S4WyFNEhk7CkSSxLs7xd8DGoDdZi6Az9MthbcUszNWOKYsRl66-ZTXH-sldA3INji1dNG_BXz_adh7Zh5U-ZFTgTVR1uHwUyC3UuvruJoN-DW_FiXT7GEAP1oZWTbgp-osg0QOURyhAvcZyyzTMohm2stBiIv0i0)

#### Kode Program Operator * (2)
```go
func main() {
  address1 := Address{"The Batcave", "outside Gotham City", "Earth Prime"}
  address2 := &address1
  address2.Place = "Wayne Manor"

  *address2 = Address{"S.T.A.R. Labs", "Central City", "Earth Prime"}
  fmt.Println(address1)
  fmt.Println(address2)
}
```

#### Dengan Operator *
![diagram](https://www.plantuml.com/plantuml/dpng/ZO_1JiCm38RlVWehzqLQxqKR4-82qWHx0Q_nQ93JeCG19iIxKw2jAEfX_oN9_oNwuarMIbg208d4j-3VVVS4S8hYEA6vvn3JsU07O9zuVEP6SUjSufpNG_kTlL33MEEXPRmcRUYL9xRpUkvHB9gzOxoDsbBuQQShDvIqnNfOzyc7t-UViE2j5hoFzc2tzi7Y7Htpp5pZZZjD99SQU-V92wEXOKo3b74yp5drbrK0umvNvMLrDNruxDV_UzXmvqgGBm00)


## Operator new
- Sebelumnya untuk membuat pointer dengan menggunakan operator &
- Go-Lang juga memiliki function new yang bisa digunakan untuk membuat pointer
- Namun function new hanya mengembalikan pointer ke data kosong, artinya tidak ada data awal
```go
func main() {
  var alamat1 = new(Address)
  var alamat2 = alamat1

  alamat2.City = "Star City"

  fmt.Println(alamat1)
  fmt.Println(alamat2)
}
```

## Pointer di Function
- Saat kita membuat parameter di function, secara default adalah pass by value, artinya data akan di copy lalu dikirim ke function tersebut
- Oleh karena itu, jika kita mengubah data di dalam function, data yang aslinya tidak akan pernah berubah.
- Hal ini membuat variable menjadi aman, karena tidak akan bisa diubah
- Namun kadang kita ingin membuat function yang bisa mengubah data asli parameter tersebut
- Untuk melakukan ini, kita juga bisa menggunakan pointer di function
- Untuk menjadikan sebuah parameter sebagai pointer, kita bisa menggunakan operator * di parameternya

#### Kode Program Pointer di Function (1)
```go
func ChangeAddressToGotham(address Address) {
  address.City = "Gotham"
}

func main() {
  address := Address{}
  ChangeAddressToGotham(address)
  fmt.Println(address) // tidak berubah
}
```

#### Kode Program Pointer di Function (2)
```go
func ChangeAddressToGotham(address *Address) {
	address.City = "Gotham"
}

func main() {
	address := Address{}
	ChangeAddressToGotham(&address)
	fmt.Println(address) //berubah
}
```

## Pointer di Method
- Walaupun method akan menempel di struct, tapi sebenarnya data struct yang diakses di dalam method adalah pass by value
- Sangat direkomendasikan menggunakan pointer di method, sehingga tidak boros memory karena harus selalu diduplikasi ketika memanggil method

#### Kode Program Pointer di Method (1)
#### Kode Program Pointer di Method (2)
