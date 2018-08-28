package main

import (
  "qs_go/models"
  "fmt"
  "net/http"
  "encoding/json"
  "github.com/gorilla/mux"
  "strconv"
)

func getMeals(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  meals := models.GetMealsFromDB()
  json.NewEncoder(w).Encode(meals)
}

func getMeal(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  params := mux.Vars(r)
  fmt.Printf("%#v\n", params)
  id, _ := strconv.Atoi(params["id"])
  fmt.Println("Getting meal with id:", id)
  meal := models.GetMealFromDB(id)
  meal.Foods = models.GetMealFoods(id)
  json.NewEncoder(w).Encode(meal)
}

func postMealFood(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  params := mux.Vars(r)
  foodID, _ := strconv.Atoi(params["food_id"])
  mealID, _ := strconv.Atoi(params["meal_id"])
  message := models.PostMealFoodToDB(foodID, mealID)
  w.WriteHeader(http.StatusNotFound)
  json.NewEncoder(w).Encode(message)
}

func deleteMealFood(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  params := mux.Vars(r)
  foodID, _ := strconv.Atoi(params["food_id"])
  mealID, _ := strconv.Atoi(params["meal_id"])
  message := models.DeleteMealFoodFromDB(foodID, mealID)
  w.WriteHeader(http.StatusNotFound)
  json.NewEncoder(w).Encode(message)
}
