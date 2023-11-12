## Recover
- Recover adalah function yang bisa kita gunakan untuk menangkap data panic
- Dengan recover proses panic akan terhenti, sehingga program akan tetap

```go
func endApp() {
  fmt.Println("App End")
  message := recover()
  fmt.Println("terjadi panic", message)
}

func runApp(error bool) {
  defer endApp()

  if error {
    panic("ERROR")
  }
}
```
```go
func main() {
  runApp(true)
  fmt.Println("panic recover")
}
```
