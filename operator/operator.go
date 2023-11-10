package main

import "fmt"

func main() {
	var a = 10
	var b = 5

	var jumlah = a * b
	var value = (((2+6)%3)*4 - 2) / 3

	fmt.Println(jumlah)
	fmt.Println(value)

	var i = 10
	i += 10 //i = i + 10
	fmt.Println(i)

	i += 5 // i = i + 5
	fmt.Println(i)
	var j = 1
	fmt.Println(j)

	j++ //j = j +	1
	fmt.Println(j)

	j-- //j = j -	1
	fmt.Println(j)

	var name1 = "John"
	var name2 = "John"

	var result bool = name1 == name2

	fmt.Println(result)

	var nilaiAkhir = 90
	var nilaiAbsen = 80

	var lulusNilaiAKhir bool = nilaiAkhir > 80
	var lulusNilaiAbsen bool = nilaiAbsen > 80

	var lulus bool = lulusNilaiAKhir && lulusNilaiAbsen

	fmt.Println(lulus) // false
}
