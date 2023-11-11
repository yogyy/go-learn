# Function

- Sebelumnya kita sudah mengenal sebuah function yang wajib dibuat agar program kita bisa berjalan, yaitu function main
  Function adalah sebuah blok kode yang sengaja dibuat dalam program agar bisa digunakan berulang-ulang
- Cara membuat function sangat sederhana, hanya dengan menggunakan kata kunci func lalu diikuti dengan nama function nya dan blok kode isi function nya
- Setelah membuat function, kita bisa mengeksekusi function tersebut dengan memanggilnya menggunakan kata kunci nama function nya diikuti tanda kurung buka, kurung tutup

```go
func sayHello() {
	fmt.Println("hello")
}

func main() {
	sayHello()
}
```

## Function Parameter

- Saat membuat function, kadang-kadang kita membutuhkan data dari luar, atau kita sebut parameter.
- Kita bisa menambahkan parameter di function, bisa lebih dari satu
- Parameter tidaklah wajib, jadi kita bisa membuat function tanpa parameter seperti sebelumnya yang sudah kita buat
- Namun jika kita menambahkan parameter di function, maka ketika memanggil function tersebut, kita wajib memasukkan data ke parameternya

```go
func sayHello(name string) {
	fmt.Println("hello", name)
}

func main() {
	sayHello("yogyy")
	sayHello("constantine")
}
```

## Function Return Value

- Function bisa mengembalikan data
- Untuk memberitahu bahwa function mengembalikan data, kita harus menuliskan tipe data kembalian dari function tersebut
- Jika function tersebut kita deklarasikan dengan tipe data pengembalian, maka wajib di dalam function nya kita harus mengembalikan data
- Untuk mengembalikan data dari function, kita bisa menggunakan kata kunci return, diikuti dengan datanya

```go
func getHello(name string) string {
	return "hello " + name
}

func main() {
	res := getHello("lex")
	fmt.Println(res)
}
```

## Returning Multiple Values

- Function tidak hanya dapat mengembalikan satu value, tapi juga bisa multiple value
- Untuk memberitahu jika function mengembalikan multiple value, kita harus menulis semua tipe data return value nya di function

```go
func getFullName() (string, string) {
	return "John", "Constantine"
}

func main() {
	firstName, lastName := getFullName()

	fmt.Println(firstName, lastName)
}
```

## Menghiraukan Return Value

- Multiple return value wajib ditangkap semua value nya
- Jika kita ingin menghiraukan return value tersebut, kita bisa menggunakan tanda \_ (garis bawah)

```go
func getFullName() (string, string) {
	return "John", "Constantine"
}

func main() {
	_, lastName := getFullName()
	fmt.Println(lastName)
}
```

## Named Return Values

- Biasanya saat kita memberi tahu bahwa sebuah function mengembalikan value, maka kita hanya mendeklarasikan tipe data return value di function
- Namun kita juga bisa membuat variable secara langsung di tipe data return function nya

```go
func getCompleteName() (firstName, middleName, lastName string) {
	firstName = "Bruce"
	middleName = "Gotham"
	lastName = "Wayne"

	return firstName, middleName, lastName
}

func main() {
	a, b, c := getCompleteName()
	fmt.Println(a, b, c)
}
```

## Variadic Function

- Parameter yang berada di posisi terakhir, memiliki kemampuan dijadikan sebuah varargs
- Varargs artinya datanya bisa menerima lebih dari satu input, atau anggap saja semacam Array.
- Apa bedanya dengan parameter biasa dengan tipe data Array?
  - Jika parameter tipe Array, kita wajib membuat array terlebih dahulu sebelum mengirimkan ke function
  - JIka parameter menggunakan varargs, kita bisa langsung mengirim data nya, jika lebih dari satu, cukup gunakan tanda koma

```go
func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	fmt.Println(sumAll(10, 10, 10))
}
```

### Slice Parameter

- Kadang ada kasus dimana kita menggunakan Variadic Function, namun memiliki variable berupa slice
- Kita bisa menjadikan slice sebagai vararg parameter

```go
numbers := []int{10, 10, 10, 10}
fmt.Println(sumAll(numbers...))
```

## Function Value

- Function adalah first class citizen
- Function juga merupakan tipe data, dan bisa disimpan di dalam variable

```go
func getGoodBye(name string) string {
	return "Good Bye " + name
}

func main() {
	goodbye := getGoodBye

	fmt.Println(goodbye("clark"))
}
```

## Function as Parameter

- Function tidak hanya bisa kita simpan di dalam variable sebagai value
- Namun juga bisa kita gunakan sebagai parameter untuk function

```go
func sayHiwithFilter(name string, filter func(string) string) {
	filtered := filter(name)

	fmt.Println(filtered)
}

func filterName(name string) string {
	if name == "Barry" {
		return "Flash"
	} else {
		return name
	}
}
```

```go
func main() {
	sayHiwithFilter("John", filterName)

	filter := filterName
	sayHiwithFilter("Barry", filter)

}
```

### Function Type Declarations

- Kadang jika function terlalu panjang, agak ribet untuk menuliskannya di dalam parameter
- Type Declarations juga bisa digunakan untuk membuat alias function, sehingga akan mempermudah kita menggunakan function sebagai parameter

```go
type Filter = func(string) string

func sayHiwithFilter(name string, filter Filter) {
	filtered := filter(name)

	fmt.Println(filtered)
}
```

## Anonymous Function

- Sebelumnya setiap membuat function, kita akan selalu memberikan sebuah nama pada function tersebut
- Namun kadang ada kalanya lebih mudah membuat function secara langsung di variable atau parameter tanpa harus membuat function terlebih dahulu
- Hal tersebut dinamakan anonymous function, atau function tanpa nama

```go
type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You are not Allowed", name)
	} else {
		fmt.Println("Welcome!", name)
	}
}
```

```go
func main() {
	blaclist := func(name string) bool {
		return name == "arthur"
	}

	registerUser("bruce", blaclist)
	registerUser("arthur", blaclist)
}
```

## Recursive Function

- Recursive function adalah function yang memanggil function dirinya sendiri
- Kadang dalam pekerjaan, kita sering menemui kasus dimana menggunakan recursive function lebih mudah dibandingkan tidak menggunakan recursive function
- Contoh kasus yang lebih mudah diselesaikan menggunakan recursive adalah Factorial

#### Factorial For Loop

```go
func factorialLoop(value int) int {
	res := 1
	for i := value; i > 0; i-- {
		res *= i
	}

	return res
}

func main() {
	res := 10 * 9 * 8 * 7 * 6 * 5 * 4 * 3 * 2 * 1
	fmt.Println(res)
	fmt.Println(factorialLoop(10))
}
```

#### Factorial Recursive

```go
func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialRecursive((value - 1))
	}
}

func main() {
	fmt.Println(factorialRecursive(10))
}
```

## Closure

- Closure adalah kemampuan sebuah function berinteraksi dengan data-data disekitarnya dalam scope yang sama
- Harap gunakan fitur closure ini dengan bijak saat kita membuat aplikasi

```go
func main() {
	counter := 0

	increment := func() {
		fmt.Println("Increment")
		counter++
	}

	increment()
	increment()
	fmt.Println(counter)
}
```
