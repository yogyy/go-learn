## Defer
- Defer function adalah function yang bisa kita jadwalkan untuk dieksekusi setelah sebuah function selesai di eksekusi
- Defer function akan selalu dieksekusi walaupun terjadi error di function yang dieksekusi

```go
func logging() {
  fmt.Println("selesai memanggil function")
}

func runApp() {
  defer logging()
  fmt.Println("Run application")
}

func main() {
  runApp()
}
```
result :
```bash
Run application
selesai memanggil function
```
