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

func addFoodToDB(food Food) int {
  db := db()
  fmt.Println(db)
  queryString := "INSERT INTO foods (name, calories) VALUES ($1, $2) RETURNING id"
  fmt.Println(queryString)
  id := 0
  err := db.QueryRow(queryString, food.Name, food.Calories).Scan(&id)
  if err != nil {
    log.Fatal(err)
  }
  return id
}
