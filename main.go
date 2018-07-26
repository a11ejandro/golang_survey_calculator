package main
import (
  "fmt"
  _ "github.com/joho/godotenv/autoload"
  "github.com/benmanns/goworker"
)

func init() {
  registerWorker()
}

func main() {
  if err := goworker.Work(); err != nil {
    fmt.Println("Error:", err)
  }
}
