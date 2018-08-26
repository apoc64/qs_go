package main

import (
  "fmt"
  "net/http"
  "os"
  "log" // error logging library
  "github.com/gorilla/mux"
  "github.com/rs/cors"
)

var foods []Food // Mock data Slice - Remove Later

func main() {
  r := mux.NewRouter()
  setRoutes(r)

  port := getPort()
  c := cors.New(cors.Options{
    AllowedMethods: []string{"GET", "POST", "PATCH", "DELETE"},
    Debug: true,
  })
  handler := c.Handler(r)
  fmt.Println("Preparing to listening on port", port)
  log.Fatal(http.ListenAndServe(port, handler))
}

func setRoutes(r *mux.Router) {
  // Mock Data - Remove Later:
  foods = append(foods, Food{ID: 1, Name: "Pizza", Calories: 400})
  foods = append(foods, Food{ID: 2, Name: "Cat", Calories: 800})

  r.HandleFunc("/api/v1/foods", getFoods).Methods("GET")
  r.HandleFunc("/api/v1/foods/{id}", getFood).Methods("GET")
  r.HandleFunc("/api/v1/foods", createFood).Methods("POST")
  r.HandleFunc("/api/v1/foods/{id}", updateFood).Methods("PATCH")
  r.HandleFunc("/api/v1/foods/{id}", deleteFood).Methods("DELETE")
}

func getPort() string {
  p := os.Getenv("PORT")
  if p != "" {
    return ":" + p
  }
  return ":3000"
}

/*
To run from the command line:
go run $(ls -1 *.go | grep -v _test.go)
*/
