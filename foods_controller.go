package main

import (
  "fmt"
  "net/http"
  "encoding/json"
  "github.com/gorilla/mux"
  "strconv"
)

// Food Struct (model)
type Food struct {
  ID        int    `json:"id"`
  Name      string  `json:"name"`
  Calories  int     `json:"calories"`
}

func getFoods(w http.ResponseWriter, r *http.Request) {
  fmt.Println("Get foods func run")
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
