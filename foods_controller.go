package main

import (
  "fmt"
  "net/http"
  "encoding/json"
  "github.com/gorilla/mux"
  "strconv"
)

func getFoods(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  foods := getFoodsFromDB()
  json.NewEncoder(w).Encode(foods)
}

func getFood(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  params := mux.Vars(r)
  id, _ := strconv.Atoi(params["id"])
  food := getFoodFromDB(id)
  json.NewEncoder(w).Encode(food)
}

type FoodHolder struct {
  TempFood        TempFood    `json:"food"`
}
type TempFood struct {
  Name string `json:"name"`
  Calories string `json:"calories"`
}

func createFood(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  var foodHolder FoodHolder
  _ = json.NewDecoder(r.Body).Decode(&foodHolder)
  fmt.Printf("Params: %#v\n", foodHolder)
  calories, _ := strconv.Atoi(foodHolder.TempFood.Calories)
  food := Food{Name: foodHolder.TempFood.Name, Calories: calories}
  id := addFoodToDB(food)
  food.ID = id
  json.NewEncoder(w).Encode(food)
}

func updateFood(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  params := mux.Vars(r)
  id, _ := strconv.Atoi(params["id"])
  var foodHolder FoodHolder
  _ = json.NewDecoder(r.Body).Decode(&foodHolder)
  calories, _ := strconv.Atoi(foodHolder.TempFood.Calories)
  food := Food{ID: id, Name: foodHolder.TempFood.Name, Calories: calories}
  if updateFoodInDB(food) {
    json.NewEncoder(w).Encode(food)
  } else {
    w.WriteHeader(http.StatusNotFound)
  }
}

func deleteFood(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  params := mux.Vars(r)
  id, _ := strconv.Atoi(params["id"])
  if deleteFoodFromDB(id) {
    w.WriteHeader(http.StatusNoContent)
  } else {
    w.WriteHeader(http.StatusNotFound)
  }
}
