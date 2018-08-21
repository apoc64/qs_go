package main

import (
  "fmt"
  "net/http"
  "os"
)

func getPort() string {
  p := os.Getenv("PORT")
  if p != "" {
    return ":" + p
  }
  return ":3000"
}

func main() {
  port := getPort()

  fmt.Println("hello world")

  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello from go")
  })

  fmt.Println(http.ListenAndServe(port, nil))
}
