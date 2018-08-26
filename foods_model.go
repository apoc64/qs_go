package main

import (
  "fmt"
  "log"
  // "os"
  // "database/sql"
  // _ "github.com/lib/pq"
)

// Food Struct (model)
type Food struct {
  ID        int    `json:"id"`
  Name      string  `json:"name"`
  Calories  int     `json:"calories"`
}

func getFoodsFromDB() {
  db()
}

func addFoodToDB(food Food) {
  fmt.Println(food)
  fmt.Println(food.Name)
  fmt.Println(food.Calories)
  db := db()
  queryString := fmt.Sprintf("INSERT INTO foods (name, calories) VALUES (%s, %v);", food.Name, food.Calories)
  fmt.Println(queryString)
  if _, err := db.Exec(queryString); err != nil {
    log.Fatal(err)
  }
}
