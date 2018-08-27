package main

import (
  "fmt"
  "log"
)

// Food Struct (model)
type Food struct {
  ID        int    `json:"id"`
  Name      string  `json:"name"`
  Calories  int     `json:"calories"`
}

func getFoodsFromDB() []Food {
  queryString := "SELECT * FROM foods"
  rows, err := db().Query(queryString)
  if err != nil {
    log.Fatal(err)
  }
  var (
    id int
    name string
    calories int
    foods []Food
  )
  defer rows.Close()
  for rows.Next() {
    err := rows.Scan(&id, &name, &calories)
    if err != nil {
      log.Fatal(err)
    }
    food := Food{ID: id, Name: name, Calories: calories}
    foods = append(foods, food)
  }
  return foods
}

func addFoodToDB(food Food) int {
  db := db()
  queryString := "INSERT INTO foods (name, calories) VALUES ($1, $2) RETURNING id"
  fmt.Println(queryString)
  id := 0
  err := db.QueryRow(queryString, food.Name, food.Calories).Scan(&id)
  if err != nil {
    log.Fatal(err)
  }
  return id
}
