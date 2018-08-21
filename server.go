package main

import (
  "fmt"
  "net/http"
  "os"
  "encoding/json"
  "log" // errors
  "github.com/gorilla/mux"
)

// func getPort() string {
//   p := os.Getenv("PORT")
//   if p != "" {
//     return ":" + p
//   }
//   return ":3000"
// }

func main() {
  // Init router
  r := mux.NewRouter()

  // Route handlers:
  r.HandleFunc("/api/v1/foods", getFoods).Methods("GET")
  r.HandleFunc("/api/v1/foods/{id}", getFood).Methods("GET")
  r.HandleFunc("/api/v1/foods", createFood).Methods("POST")
  r.HandleFunc("/api/v1/foods/{id}", updateFood).Methods("PATCH")
  r.HandleFunc("/api/v1/foods/{id}", deleteFood).Methods("DELETE")

  log.Fatal(http.ListenAndServe(":3000", r)
  // port := getPort()
  //
  // http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
  //   fmt.Println("Request made to /")
  //   fmt.Fprintf(w, "Hello from go")
  // })
  //
  // fmt.Println("Listening on port", port)
  // err := http.ListenAndServe(port, nil)
  // if err != nil {
  //   panic(err)
  // }
} // end main
