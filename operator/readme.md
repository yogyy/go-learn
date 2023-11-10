# Operator

## Operator Aritmatika

| Operator | Keterangan     |
| -------- | -------------- |
| +        | Penjumlahan    |
| -        | Pengurangan    |
| \*       | Perkalian      |
| /        | Pembagian      |
| %        | Sisa Pembagian |

```go
	var a = 10
	var b = 5

	var jumlah = a * b
	var value = (((2+6)%3)*4 - 2) / 3

	fmt.Println(jumlah)
	fmt.Println(value)
```

## Augmented Assigment

| Operasi Matematika | Augmented Assignments |
| ------------------ | --------------------- |
| a = a + 10         | a += 10               |
| a = a - 10         | a -= 10               |
| a = a \* 10        | a \*= 10              |
| a = a / 10         | a /= 10               |
| a = a % 10         | a %= 10               |

```go
	var i = 10
	i += 10 //i = i + 10
	fmt.Println(i)

	i += 5 // i = i + 5
	fmt.Println(i)
```

## Unary Operator

| Operator | Keterangan        |
| -------- | ----------------- |
| ++       | a = a + 1         |
| --       | a = a - 1         |
| -        | Negative          |
| +        | Positive          |
| !        | Boolean kebalikan |

```go
	var j = 1
	fmt.Println(j)

	j++ //j = j +	1
	fmt.Println(j)

	j-- //j = j -	1
	fmt.Println(j)
```

## Operasi Perbandingan

- Operasi perbandingan adalah operasi untuk membandingkan dua buah data
- Operasi perbandingan adalah operasi yang menghasilkan nilai boolean (benar atau salah)
- Jika hasil operasinya adalah benar, maka nilainya adalah true
- Jika hasil operasinya adalah salah, maka nilainya adalah false

| Operator | Keterangan              |
| -------- | ----------------------- |
| >        | Lebih Dari              |
| <        | Kurang Dari             |
| >=       | Lebih Dari Sama Dengan  |
| <=       | Kurang Dari Sama Dengan |
| ==       | Sama Dengan             |
| !=       | Tidak Sama Dengan       |

```go
	var name1 = "John"
	var name2 = "John"

	var result bool = name1 == name2

	fmt.Println(result) //true
```

## Operasi Boolean

| Operator 	| Keterangan              	|
|----------	|-------------------------	|
| &&       	| Dan                     	|
| \|\|     	| Atau                    	|
| !        	| Kebalikan               	|
| <=       	| Kurang Dari Sama Dengan 	|
| ==       	| Sama Dengan             	|
| !=       	| Tidak Sama Dengan       	|

## Operasi &&

| Nilai 1 	| Operator 	| Nilai 2 	| Hasil 	|
|---------	|----------	|---------	|-------	|
| true    	| &&       	| true    	| true  	|
| true    	| &&       	| false   	| false 	|
| false   	| &&       	| true    	| false 	|
| false   	| &&       	| false   	| false 	|

## Operasi ||

| Nilai 1 	| Operator 	| Nilai 2 	| Hasil 	|
|---------	|----------	|---------	|-------	|
| true    	| \|\|     	| true    	| true  	|
| true    	| \|\|     	| false   	| true  	|
| false   	| \|\|     	| true    	| true  	|
| false   	| \|\|     	| false   	| false 	|

## Operasi !

| Operator 	| Nilai 2 	| Hasil 	|
|----------	|---------	|-------	|
| !        	| true    	| false 	|
| !        	| false   	| true  	|

```go
	var nilaiAkhir = 90
	var nilaiAbsen = 80

	var lulusNilaiAKhir bool = nilaiAkhir > 80
	var lulusNilaiAbsen bool = nilaiAbsen > 80

	var lulus bool = lulusNilaiAKhir && lulusNilaiAbsen

	fmt.Println(lulus) // false
```
