package main

import (
  "fmt"
  "net/http"
  "os"
  "log"
  "github.com/gorilla/mux"
  "github.com/gorilla/handlers"
)

func main() {
  r := mux.NewRouter()
  setRoutes(r)

  port := getPort()
  fmt.Println("Preparing to listening on port", port)

  log.Fatal(http.ListenAndServe(port, handlers.CORS(
    handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
    handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "DELETE", "OPTIONS"}),
    handlers.AllowedOrigins([]string{"*"}))(r)))
}

func setRoutes(r *mux.Router) {
  r.HandleFunc("/api/v1/foods/", getFoods).Methods("GET")
  r.HandleFunc("/api/v1/foods/{id}/", getFood).Methods("GET")
  r.HandleFunc("/api/v1/foods/", createFood).Methods("POST")
  r.HandleFunc("/api/v1/foods/{id}/", updateFood).Methods("PATCH")
  r.HandleFunc("/api/v1/foods/{id}/", deleteFood).Methods("DELETE")
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
