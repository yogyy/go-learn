## Tipe Data Slice
- Tipe data Slice adalah potongan dari data Array
- Slice mirip dengan Array, yang membedakan adalah ukuran Slice bisa berubah
- Slide dan Array selalu terkoneksi, dimana Slice adalah data yang mengakses sebagian atau seluruh data di Array

### Detail Tipe Data Slice
- Tipe Data Slice memiliki 3 data, yaitu pointer, length dan capacity
- Pointer adalah penunjuk data pertama di array para slice
- Length adalah panjang dari slice, dan
- Capacity adalah kapasitas dari slice, dimana length tidak boleh lebih dari capacity


### Membuat Slice Dari Array
| Membuat Slice   	| Keterangan                                                             	|
|-----------------	|------------------------------------------------------------------------	|
| array[low:high] 	| Membuat slice dari array dimulai index low sampai index sebelum high   	|
| array[low:]     	| Membuat slide dari array dimulai index low sampai index akhir di array 	|
| array[:high]    	| Membuat slice dari array dimulai index 0 sampai index sebelum high     	|
| array[:]        	| Membuat slice dari array dimulai index 0 sampai index akhir di array   	|

### Slice dan Array
![diagram](https://www.plantuml.com/plantuml/dpng/XP71IiD048Rl-nH3xm6JnhGDA2s81q45FOg7JPoagzlTC9a8HNntTzQjY1Pw_Fxz-9dPMO-iYQUTLRvvnPP14-azV2Y0CxY0sOrsaoOp2vmBNsD3Xw2Gu5OJij1SQ3EGiK9bVCFedSUYoKMeli56M0Xi-cdaQHHU2Z_YmBXha2HPINB_nZvn7gUwy-y_HfvBSeLZIZxLOPNHcXRpk0l0u8rZfE2MaIswdhS1vn5RFUZMy3u4oxyLnR-AwaH59HNECmBx0NoNHb7nSv_AUl6iqBs064Uy8dbLlVhjfNKoClsMiyjqgHeUi4D2rw9AqsOMk-F2EhI5Mty1)

```go
batmanVillains := [...]string{"Joker", "Ra's al Ghul", "Bane", "Riddler", "Penguin", "Two-Face", "Scarecrow"}

slice1 := batmanVillains[4:6]
fmt.Println(slice1) // [Penguin Two-Face]

slice2 := batmanVillains[:2]
fmt.Println(slice2) // [Joker Ra's al Ghul]

slice3 := batmanVillains[3:]
fmt.Println(slice3) // [Riddler Penguin Two-Face Scarecrow]

slice4 := batmanVillains[:]
fmt.Println(slice4) // [Joker Ra's al Ghul Bane Riddler Penguin Two-Face Scarecrow]
```

### Function Slice
| Operasi                            	| Keterangan                                                                                                                 	|
|------------------------------------	|----------------------------------------------------------------------------------------------------------------------------	|
| len(slice)                         	| Untuk mendapatkan panjang                                                                                                  	|
| cap(slice)                         	| Untuk mendapat kapasitas                                                                                                   	|
| append(slice, data)                	| Membuat slice baru dengan menambah data ke posisi terakhir slice, jika kapasitas sudah penuh, maka akan membuat array baru 	|
| make([]TypeData, length, capacity) 	| Membuat slice baru                                                                                                         	|
| copy(destination, source)          	| Menyalin slice dari source ke destination                                                                                  	|

#### Append Slice
```go
  days := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jum'at", "Sabtu", "Minggu"}

  daySlice1 := days[5:] // Sabtu,Minggu
  fmt.Println(daySlice1)

  daySlice1[0] = "Sabtu Baru"
  daySlice1[1] = "Minggu Baru"

  fmt.Println(daySlice1)
  fmt.Println(days)

  daySlice2 := append(daySlice1, "Libur Baru")
  fmt.Println(daySlice1)
  fmt.Println(daySlice2)
  fmt.Println(days)

```

#### Make Slice
```go
  var newSlice []string = make([]string, 2, 5)
  newSlice[0] = "Bruce"
  newSlice[1] = "Wayne"

  fmt.Println(newSlice)
  fmt.Println(len(newSlice))
  fmt.Println(cap(newSlice))

  newSlice2 := append(newSlice, "and Alfred")
  fmt.Println(newSlice2)
  fmt.Println(len(newSlice2))
  fmt.Println(cap(newSlice2))

  newSlice2[0] = "Thomas"
  fmt.Println(newSlice2)
  fmt.Println(newSlice)
```

#### Copy Slice
```go
  fromSlice := batmanVillains[:]
  toSlice := make([]string, len(fromSlice), cap(fromSlice))

  copy(toSlice, fromSlice)

  fmt.Println(fromSlice)
  fmt.Println(toSlice)
```
