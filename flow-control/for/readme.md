## For
- Dalam bahasa pemrograman, biasanya ada fitur yang bernama perulangan
- Salah satu fitur perulangan adalah for loops

```go
func main() {
  counter := 1

  for counter <= 10 {
    fmt.Println("perulangan ke", counter)
    counter++
  }
}
```

### Kode Program For dengan Statement
```go
for counter := 1; counter <= 10; counter++ {
  fmt.Println("perulangan ke", counter)
}
```

### For Range
- For bisa digunakan untuk melakukan iterasi terhadap semua data collection
- Data collection contohnya Array, Slice dan Map
```go
names := []string{"Batman", "Robin", "Nightwing"}
for i := 0; i < len(names); i++ {
  fmt.Println(names[i])
}
```
or
```go
for index, name := range names {
  fmt.Println("index", index, "=", name)
}
 // mor simply
for _, v := range names {
  fmt.Println(v)
}
```
