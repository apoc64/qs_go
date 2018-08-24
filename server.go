package main

import (
  "fmt"
  "net/http"
  "os"
  "log" // error loggin library
  "github.com/gorilla/mux"
)

func getPort() string {
  p := os.Getenv("PORT")
  if p != "" {
    return ":" + p
  }
  return ":3000"
}

var foods []Food

func main() {
  r := mux.NewRouter()

  // Mock Data
  foods = append(foods, Food{ID: 1, Name: "Pizza", Calories: 400})
  foods = append(foods, Food{ID: 2, Name: "Cat", Calories: 800})

  // Route handlers:
  r.HandleFunc("/api/v1/foods", getFoods).Methods("GET")
  r.HandleFunc("/api/v1/foods/{id}", getFood).Methods("GET")
  r.HandleFunc("/api/v1/foods", createFood).Methods("POST")
  r.HandleFunc("/api/v1/foods/{id}", updateFood).Methods("PATCH")
  r.HandleFunc("/api/v1/foods/{id}", deleteFood).Methods("DELETE")

  port := getPort()
  fmt.Println("About to listening on port", port)
  log.Fatal(http.ListenAndServe(port, r))

} // end main
