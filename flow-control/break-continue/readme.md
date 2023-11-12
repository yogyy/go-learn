## Break & Continue
- Break & continue adalah kata kunci yang bisa digunakan dalam perulangan
- Break digunakan untuk menghentikan seluruh perulangan
- Continue adalah digunakan untuk menghentikan perulangan yang berjalan, dan langsung melanjutkan ke perulangan selanjutnya

break :
```go
for i := 0; i < 10; i++ {
  if i == 5 {
    break
  }
  fmt.Println("perulangan ke", i)
}
```
continue:
```go
for j := 0; j < 10; j++ {
  if j%2 == 0 {
    continue
  }
  fmt.Println("perulangan ke", j)
}
```
