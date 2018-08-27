package main

import (
  // "fmt"
  "net/http"
  "encoding/json"
  // "github.com/gorilla/mux"
  // "strconv"
)

func getMeals(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  meals := getMealsFromDB()
  json.NewEncoder(w).Encode(meals)
}
