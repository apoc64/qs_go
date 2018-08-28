package controllers

import (
  "qs_go/models"
  "fmt"
  "net/http"
  "encoding/json"
  "github.com/gorilla/mux"
  "strconv"
)

func GetFoods(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  foods := models.GetFoodsFromDB()
  json.NewEncoder(w).Encode(foods)
}

func GetFood(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  params := mux.Vars(r)
  id, _ := strconv.Atoi(params["id"])
  food := models.GetFoodFromDB(id)
  json.NewEncoder(w).Encode(food)
}

type FoodHolder struct {
  TempFood        TempFood    `json:"food"`
}
type TempFood struct {
  Name string `json:"name"`
  Calories string `json:"calories"`
}

func CreateFood(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  var foodHolder FoodHolder
  _ = json.NewDecoder(r.Body).Decode(&foodHolder)
  fmt.Printf("Params: %#v\n", foodHolder)
  calories, _ := strconv.Atoi(foodHolder.TempFood.Calories)
  food := models.Food{Name: foodHolder.TempFood.Name, Calories: calories}
  id := models.AddFoodToDB(food)
  food.ID = id
  json.NewEncoder(w).Encode(food)
}

func UpdateFood(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  params := mux.Vars(r)
  id, _ := strconv.Atoi(params["id"])
  fmt.Println("Update food method called")
  var foodHolder FoodHolder
  _ = json.NewDecoder(r.Body).Decode(&foodHolder)
  calories, _ := strconv.Atoi(foodHolder.TempFood.Calories)
  food := models.Food{ID: id, Name: foodHolder.TempFood.Name, Calories: calories}
  fmt.Println("Preparing to update food:", food)
  if models.UpdateFoodInDB(food) {
    json.NewEncoder(w).Encode(food)
  } else {
    w.WriteHeader(http.StatusNotFound)
  }
}

func DeleteFood(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  params := mux.Vars(r)
  id, _ := strconv.Atoi(params["id"])
  if models.DeleteFoodFromDB(id) {
    w.WriteHeader(http.StatusNoContent)
  } else {
    w.WriteHeader(http.StatusNotFound)
  }
}
