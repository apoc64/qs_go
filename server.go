package main

import (
  "fmt"
  "net/http"
  "os"
  "encoding/json"
  "log" // errors
  "github.com/gorilla/mux"
  "strconv"
)

func getPort() string {
  p := os.Getenv("PORT")
  if p != "" {
    return ":" + p
  }
  return ":3000"
}

// Food Struct (model)
type Food struct {
  ID        int    `json:"id"`
  Name      string  `json:"name"`
  Calories  int     `json:"calories"`
}

// slice - mutable array data type
var foods []Food

func getFoods(w http.ResponseWriter, r *http.Request) {
  // w reporesents response writer
  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(foods)
}

func getFood(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  params := mux.Vars(r)
  // Loop through books - DB???
  for _, item := range foods {
    id, err := strconv.Atoi(params["id"])
    fmt.Println(err)
    if item.ID == id {
      json.NewEncoder(w).Encode(item)
      return
    }
  }
  json.NewEncoder(w).Encode(&Food{})
}

func createFood(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  var food Food
  _ = json.NewDecoder(r.Body).Decode(&food)
  food.ID = len(foods) + 1 // not for db
  foods = append(foods, food)
  json.NewEncoder(w).Encode(food)

}

func updateFood(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  params := mux.Vars(r)
  for index, item := range foods {
    id, err := strconv.Atoi(params["id"])
    fmt.Println(err)
    if item.ID == id {
      foods = append(foods[:index], foods[index+1:]...)
      var food Food
      _ = json.NewDecoder(r.Body).Decode(&food)
      food.ID = index
      foods = append(foods, food)
      json.NewEncoder(w).Encode(food)
      return
    }
  }
  json.NewEncoder(w).Encode(foods)
}

func deleteFood(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  params := mux.Vars(r)
  for index, item := range foods {
    id, err := strconv.Atoi(params["id"])
    fmt.Println(err)
    if item.ID == id {
      foods = append(foods[:index], foods[index+1:]...)
      break
    }
  }
  json.NewEncoder(w).Encode(foods)
}

func main() {
  // Init router
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
  //
  // http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
  //   fmt.Println("Request made to /")
  //   fmt.Fprintf(w, "Hello from go")
  // })
  //
  // fmt.Println("Hello")
  // err := http.ListenAndServe(port, nil)
  // if err != nil {
  //   panic(err)
  // }
} // end main
