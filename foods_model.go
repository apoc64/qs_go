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
  fmt.Println(db)
  queryString := "INSERT INTO foods (name, calories) VALUES ($1, $2)"
  fmt.Println(queryString)
  result, err := db.Exec(queryString, food.Name, food.Calories)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println(result)
}
