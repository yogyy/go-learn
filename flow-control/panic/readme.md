## Panic
- Panic function adalah function yang bisa kita gunakan untuk menghentikan program
- Panic function biasanya dipanggil ketika terjadi panic pada saat program kita berjalan
- Saat panic function dipanggil, program akan terhenti, namun defer function tetap akan dieksekusi

```go
func endApp() {
  fmt.Println("App End")
}

func runApp(error bool) {
  defer endApp()

  if error {
    panic("ERROR")
  }

  fmt.Println("Hello")
}
```
```go
func main() {
  runApp(true)
}
```
