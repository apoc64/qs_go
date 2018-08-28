package main

import (
  // "fmt"
  "net/http"
  "encoding/json"
  "github.com/gorilla/mux"
  "strconv"
)

func getMeals(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  meals := getMealsFromDB()
  json.NewEncoder(w).Encode(meals)
}

func postMealFood(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  params := mux.Vars(r)
  foodID, _ := strconv.Atoi(params["food_id"])
  mealID, _ := strconv.Atoi(params["meal_id"])
  message := postMealFoodToDB(foodID, mealID)
  w.WriteHeader(http.StatusNotFound)
  json.NewEncoder(w).Encode(message)
}
