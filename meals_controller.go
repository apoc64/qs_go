package main

import (
  "fmt"
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

func getMeal(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  params := mux.Vars(r)
  fmt.Printf("%#v\n", params)
  id, _ := strconv.Atoi(params["id"])
  fmt.Println("Getting meal with id:", id)
  meal := getMealFromDB(id)
  meal.Foods = getMealFoods(id)
  json.NewEncoder(w).Encode(meal)
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

func deleteMealFood(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  params := mux.Vars(r)
  foodID, _ := strconv.Atoi(params["food_id"])
  mealID, _ := strconv.Atoi(params["meal_id"])
  message := deleteMealFoodFromDB(foodID, mealID)
  w.WriteHeader(http.StatusNotFound)
  json.NewEncoder(w).Encode(message)
}
