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

  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Request made to /")
    fmt.Fprintf(w, "Hello from go")
  })

  fmt.Println("Listening on port", port)
  err := http.ListenAndServe(port, nil)
  if err != nil {
    panic(err)
  }
}
