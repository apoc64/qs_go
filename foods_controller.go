package main

import (
  "fmt"
  "net/http"
  "encoding/json"
  // "github.com/gorilla/mux"
  "strconv"
)


func getFoods(w http.ResponseWriter, r *http.Request) {
  fmt.Println("Get foods func run")
  foods := getFoodsFromDB()
  // w reporesents response writer
  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(foods)
}

func getFood(w http.ResponseWriter, r *http.Request) {
  // w.Header().Set("Content-Type", "application/json")
  // params := mux.Vars(r)
  // // Loop through books - DB???
  // for _, item := range foods {
  //   id, err := strconv.Atoi(params["id"])
  //   fmt.Println(err)
  //   if item.ID == id {
  //     json.NewEncoder(w).Encode(item)
  //     return
  //   }
  // }
  // json.NewEncoder(w).Encode(&Food{})
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
  // w.Header().Set("Content-Type", "application/json")
  // params := mux.Vars(r)
  // for index, item := range foods {
  //   id, err := strconv.Atoi(params["id"])
  //   fmt.Println(err)
  //   if item.ID == id {
  //     foods = append(foods[:index], foods[index+1:]...)
  //     var food Food
  //     _ = json.NewDecoder(r.Body).Decode(&food)
  //     food.ID = index
  //     foods = append(foods, food)
  //     json.NewEncoder(w).Encode(food)
  //     return
  //   }
  // }
  // json.NewEncoder(w).Encode(foods)
}

func deleteFood(w http.ResponseWriter, r *http.Request) {
  // w.Header().Set("Content-Type", "application/json")
  // params := mux.Vars(r)
  // for index, item := range foods {
  //   id, err := strconv.Atoi(params["id"])
  //   fmt.Println(err)
  //   if item.ID == id {
  //     foods = append(foods[:index], foods[index+1:]...)
  //     break
  //   }
  // }
  // json.NewEncoder(w).Encode(foods)
}
